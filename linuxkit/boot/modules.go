package main

import (
	"io/ioutil"
	"os"
	"strings"
	"syscall"
)

func virtualizationFlag() string {
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(data), "vmx") {
		return "vmx"
	}
	if strings.Contains(string(data), "svm") {
		return "svm"
	}
	return ""
}

func loadKvm() {
	execCommand("modprobe", "kvm", "ignore_msrs=1")
	switch virtualizationFlag() {
	case "vmx":
		execCommand("modprobe", "kvm-intel",
			"nested=1",
			"ept=1",
			"emulate_invalid_guest_state=0",
			"enable_shadow_vmcs=1",
			"enable_apicv=1",
		)
	case "svm":
		execCommand("modprobe", "kvm-amd")
	default:
		panic("CPU doesn't support virtualization (vmx or svm)")
	}
}

func setupHugePages() {
	if err := ioutil.WriteFile("/proc/sys/vm/nr_hugepages", []byte("4100"), 0666); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("/dev/hugepages", 0777); err != nil {
		panic(err)
	}
	if err := syscall.Mount("hugetlbfs", "/dev/hugepages", "hugetlbfs", 0, "rw"); err != nil {
		panic(err)
	}
}

func loadModules() {
	loadKvm()
	execCommand("modprobe", "vhost_net")
}
