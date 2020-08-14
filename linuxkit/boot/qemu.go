package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"syscall"
)

func cpuflags() string {
	flags := []string{
		"host",
		"kvm=on",
		"vendor=GenuineIntel",
		"vmware-cpuid-freq=on",
		"rdtscp",
		"check",

		// "+aes",
		// "+avx",
		// "+avx2",
		// "+bmi1",
		// "+bmi2",
		// "+invtsc",
		// "+fma",
		// "+hypervisor",
		// // "+kvm_pv_eoi",
		// // "+kvm_pv_unhalt",
		// "+pcid",
		// "+popcnt",
		// "+smep",
		// "+sse4.2",
		// "+sse4a",
		// "+ssse3",
		// "+xsave",
		// "+xsaveopt",
	}

	if flag := virtualizationFlag(); flag != "" {
		flags = append(flags, "+"+flag)
	}

	return strings.Join(flags, ",")
}

func memory() string {
	return fmt.Sprintf("%d", 4096)
}

func smp() string {
	return fmt.Sprintf("cpus=%d,maxcpus=%d", runtime.NumCPU()-1, runtime.NumCPU())
	return fmt.Sprintf("sockets=1,cores=%d,threads=%d", (runtime.NumCPU()/2)-1, 2)
}

func opt(options []string) string {
	return strings.Join(options, ",")
}

func runQemu() {
	copyFile("/macos/bootloader.img", "/tmp/bootloader.img")

	args := []string{
		"/usr/local/bin/qemu-system-x86_64",

		"-serial", "/dev/ttyS0",
		"-display", "vnc=:0",
		"-vga", "std",
		"-device", "VGA,vgamem_mb=32",
		"-accel", "kvm",
		"-machine", "q35",
		"-cpu", cpuflags(),
		"-m", memory(),
		"-smp", smp(),

		"-usb",
		"-device", "usb-kbd",
		"-device", "usb-tablet",

		"-smbios", "type=2",

		"-drive", "if=pflash,format=raw,readonly,file=/usr/share/OVMF/OVMF_CODE.fd",
		"-drive", "if=pflash,format=raw,file=/usr/share/OVMF/OVMF_VARS.fd",

		"-device", "ich9-ahci,id=sata",

		"-device", "ide-hd,bus=sata.1,drive=bootloader",
		"-drive", "id=bootloader,if=none,format=raw,file=/tmp/bootloader.img",

		"-device", "virtio-blk,drive=macos",
		"-drive", "id=macos,if=none,format=raw,file=/dev/sdb",

		"-nic", "tap,ifname=tap0,model=vmxnet3,script=no,downscript=no,vhost=on,mac=52:54:00:AB:B5:CC",
	}

	syscall.Exec(args[0], args, os.Environ())
}
