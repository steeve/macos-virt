package boot

import (
	"fmt"
	"runtime"
	"strings"
	"syscall"
)

const (
	hostMemory = 1 * GB
)

type VMConfig struct {
	Memory    int64
	HugePages int64
	CPUs      int
	CPUFlags  []string
}

func VM() VMConfig {
	si := &syscall.Sysinfo_t{}
	if err := syscall.Sysinfo(si); err != nil {
		panic(err)
	}
	totalRamMb := si.Totalram / MB

	flags := []string{
		"host",
		"kvm=on",
		"vmware-cpuid-freq=on",
		"check",
		"+invtsc", // needed on intel or else the clock runs too quick
		"+hypervisor",
	}

	return VMConfig{
		CPUs:   runtime.NumCPU() - 1,
		Memory: si.Totalram - hostMemory,
	}
}

func (vm VMConfig) CPUFlagsString() {

	if flag := virtualizationFlag(); flag != "" {
		flags = append(flags, "+"+flag)
	}

	return strings.Join(flags, ",")
}

func (vm VMConfig) SMP() {
	return fmt.Sprintf("cpus=%d,maxcpus=%d,sockets=1", vm.CPUs, vm.CPUs+(vm.CPUs%2))
}

func (vm VMConfig) MemoryString() string {
	return fmt.Sprintf("%d", vm.Memory/MB)
}

func memory() string {
	si := &syscall.Sysinfo_t{}
	if err := syscall.Sysinfo(si); err != nil {
		panic(err)
	}
	totalRamMb := si.Totalram / MB
	if totalRamMb < 2048 {
		panic("the VM should have at least 2GB of memory")
	}
	// Leave 1GB for linux
	return fmt.Sprintf("%d", totalRamMb-1024)
}

func smp() string {
	// real, used vcpus, leave one core for linux
	cpus := runtime.NumCPU() - 1
	// those are fake, disabled vcpus meant to make macOS happy with the CPU
	// topology (must be an even number)
	maxcpus := runtime.NumCPU()
	maxcpus += maxcpus % 2
	return fmt.Sprintf("cpus=%d,maxcpus=%d", cpus, maxcpus)
}
