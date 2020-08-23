package vm

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/steeve/macos-virt/linuxkit/boot/cpuinfo"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

type VM struct {
}

func MakeVM() VM {
	return VM{}
}

func (vm VM) Memory() int64 {
	const hostRatio = 0.80 // 80% of the memory for the guest

	si := &syscall.Sysinfo_t{}
	if err := syscall.Sysinfo(si); err != nil {
		panic(err)
	}

	return int64(float64(si.Totalram) * 0.80)
}

func (vm VM) MemoryArgs() []string {
	return []string{
		"-m", fmt.Sprintf("%d", vm.Memory()/MB),
	}
}

func (vm VM) CPUArgs() []string {
	return []string{
		"-cpu", strings.Join([]string{
			"host",           // CPU passthrough
			"l3-cache=on",    // forward the L3 cache too
			"migratable=off", // ensures invtsc and friends are enabled
			"enforce",        // fail if not all flags are available
			"+check",
			"+invtsc", // invariant TSC is absolutely needed

			// ensures the VM sees it's running inside an hypervisor
			"kvm=on",
			"vmware-cpuid-freq=on",
			"+hypervisor",
		}, ","),
	}
}

func (vm VM) SMPArgs() []string {
	infos := cpuinfo.ReadInfos()

	threads := infos.Threads()
	lc := infos.LogicalCores()
	cpus := lc - threads // leave one physical core for linux

	// those are fake, disabled vcpus meant to make macOS happy with the CPU
	// topology (must be an even number)
	maxcpus := lc + (lc % 2)

	return []string{
		"-smp", fmt.Sprintf("cpus=%d,maxcpus=%d,sockets=%d,cores=%d,threads=%d",
			cpus,
			maxcpus,
			infos.Sockets(),
			infos.PhysicalCores(),
			threads,
		),
	}
}

func (vm VM) HuegPagesArgs() []string {
	return nil
}

func (vm VM) QEMUArgs(bootloaderImage, macosBlockDevice string) []string {
	args := []string{
		"-serial", "/dev/ttyS0", // serial forwarding
		"-display", "vnc=:0",
		"-vga", "std",
		// "-device", "VGA,vgamem_mb=32", // avoids glitches
		"-accel", "kvm", // obviously
		"-machine", "pc-q35-2.11", // known to be compatible with macOS

		"-usb",               // enable USB
		"-device", "usb-kbd", // add USB keyboard
		"-device", "usb-tablet", // add absolute pointing device

		// Use UEFI
		"-smbios", "type=2",
		"-drive", "if=pflash,format=raw,readonly,file=/usr/share/OVMF/OVMF_CODE.fd",
		"-drive", "if=pflash,format=raw,file=/usr/share/OVMF/OVMF_VARS.fd",

		// OpenCore bootloader
		"-device", "ide-hd,bus=ide.1,drive=bootloader",
		"-drive", "id=bootloader,if=none,format=raw,file=" + bootloaderImage,

		// macOS drive, uses VirtIO for faster perf
		"-device", "virtio-blk,drive=macos",
		"-drive", "id=macos,if=none,format=raw,file=" + macosBlockDevice,

		// Add NIC
		// virtio is recognized, but packets are lost, so use vmxnet3 which is performant enough
		"-nic", "tap,ifname=tap0,model=vmxnet3,script=no,downscript=no,vhost=on,mac=52:54:00:AB:B5:CC",
	}

	args = append(args, vm.CPUArgs()...)
	args = append(args, vm.SMPArgs()...)
	args = append(args, vm.MemoryArgs()...)

	return args
}
