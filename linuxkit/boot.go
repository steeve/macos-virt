package main

import (
	"io"
	"os"
	"strings"
	"syscall"
)

func cpuflags() string {
	flags := []string{
		"Penryn",
		"kvm=on",
		"vendor=GenuineIntel",
		"vmware-cpuid-freq=on",

		"+aes",
		"+avx",
		"+avx2",
		"+invtsc",
		"+pcid",
		"+popcnt",
		"+smep",
		"+sse4.2",
		"+ssse3",
		"+xsave",
		"+xsaveopt",

		"check",
	}

	return strings.Join(flags, ",")
}

func memory() string {
	return "2048"
}

func smp() string {
	return "sockets=1,cores=2,threads=2"
}

func copyFile(dst, src string) {
	from, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		panic(err)
	}
}

func main() {
	copyFile("/tmp/bootloader.img", "/macos/bootloader.img")

	args := []string{
		"/usr/bin/qemu-system-x86_64",

		"-serial", "/dev/ttyS0",
		"-display", "vnc=:0",
		"-accel", "kvm",
		"-machine", "pc-q35-2.11",
		"-cpu", cpuflags(),
		"-m", memory(),
		"-smp", smp(),

		"-usb",
		"-device", "usb-kbd",
		"-device", "usb-tablet",

		"-smbios", "type=2",
		"-drive", "if=pflash,format=raw,readonly,file=/macos/OVMF_CODE.fd",
		"-drive", "if=pflash,format=raw,readonly,file=/macos/OVMF_VARS.fd",

		"-device", "ide-drive,bus=ide.1,drive=bootloader",
		"-drive", "id=bootloader,if=none,format=raw,file=/tmp/bootloader.img",

		"-device", "ide-drive,bus=ide.2,drive=macos",
		"-drive", "id=macos,if=none,format=raw,cache=none,file=/dev/sdb",

		// "-device", "virtio-net-pci,id=net0,netdev=net0,vectors=0,mac=52:54:00:AB:B5:CC",
		// "-netdev", "tap,id=net0,ifname=tap0,script=no,downscript=no,vhost=on",
	}

	syscall.Exec(args[0], args, os.Environ())
}
