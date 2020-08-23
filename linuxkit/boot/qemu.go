package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/steeve/macos-virt/linuxkit/boot/vm"
)

// const (
// 	KB = 1024
// 	MB = 1024 * KB
// 	GB = 1024 * MB
// )

// func cpuflags() string {
// 	flags := []string{
// 		"host",
// 		"kvm=on",
// 		"l3-cache=on",
// 		"vmware-cpuid-freq=on",
// 		"migratable=off", // to ensure invtsc and friends are enabled
// 		"enforce",        // do not start if all the in
// 		"+check",
// 		"+invtsc", // needed on intel or else the clock runs too quick
// 		"+hypervisor",
// 	}

// 	if flag := virtualizationFlag(); flag != "" {
// 		flags = append(flags, "+"+flag)
// 	}

// 	return strings.Join(flags, ",")
// }

// func memory() string {
// 	si := &syscall.Sysinfo_t{}
// 	if err := syscall.Sysinfo(si); err != nil {
// 		panic(err)
// 	}
// 	totalRamMb := si.Totalram / MB
// 	if totalRamMb < 2048 {
// 		panic("the VM should have at least 2GB of memory")
// 	}
// 	// Leave 1GB for linux
// 	return fmt.Sprintf("%d", totalRamMb-1024)
// }

// func smp() string {
// 	infos := cpuinfo.ReadInfos()

// 	threads := infos.Threads()
// 	cpus := infos.LogicalCores() - threads // leave one physical core for linux

// 	// those are fake, disabled vcpus meant to make macOS happy with the CPU
// 	// topology (must be an even number)
// 	maxcpus := infos.LogicalCores()
// 	maxcpus += maxcpus % 2
// 	return fmt.Sprintf("cpus=%d,maxcpus=%d,sockets=%d,cores=%d,threads=%d",
// 		cpus,
// 		maxcpus,
// 		infos.Sockets(),
// 		infos.PhysicalCores(),
// 		threads,
// 	)
// }

// func opt(options []string) string {
// 	return strings.Join(options, ",")
// }

func runQemu() {
	copyFile("/macos/bootloader.img", "/tmp/bootloader.img")
	// setupHugePages()

	args := []string{
		"/usr/bin/qemu-system-x86_64",
	}

	qemuVM := vm.MakeVM()
	args = append(args, qemuVM.QEMUArgs("/tmp/bootloader.img", "/dev/sdb")...)

	// args := append qemuVM.QEMUArgs()

	// args := []string{
	// 	"/usr/bin/qemu-system-x86_64",

	// 	"-serial", "/dev/ttyS0",
	// 	"-display", "vnc=:0",
	// 	"-vga", "std",
	// 	"-device", "VGA,vgamem_mb=32",
	// 	"-accel", "kvm",
	// 	"-machine", "q35",
	// 	"-cpu", cpuflags(),
	// 	"-m", memory(),
	// 	// "-mem-path", "/dev/hugepages",
	// 	"-smp", smp(),

	// 	"-usb",
	// 	"-device", "usb-kbd",
	// 	"-device", "usb-tablet",

	// 	"-smbios", "type=2",
	// 	"-drive", "if=pflash,format=raw,readonly,file=/usr/share/OVMF/OVMF_CODE.fd",
	// 	"-drive", "if=pflash,format=raw,file=/usr/share/OVMF/OVMF_VARS.fd",

	// 	"-device", "ich9-ahci,id=sata",

	// 	"-device", "ide-hd,bus=sata.1,drive=bootloader",
	// 	"-drive", "id=bootloader,if=none,format=raw,file=/tmp/bootloader.img",

	// 	"-device", "virtio-blk,drive=macos",
	// 	"-drive", "id=macos,if=none,format=raw,file=/dev/sdb",

	// 	// virtio is recognized, but packets are lost, so use vmxnet3 which is performant enough
	// 	"-nic", "tap,ifname=tap0,model=vmxnet3,script=no,downscript=no,vhost=on,mac=52:54:00:AB:B5:CC",
	// }

	fmt.Println("qemu command line", args)
	// return
	syscall.Exec(args[0], args, os.Environ())
}
