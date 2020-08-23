package cpuinfo

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

const (
	VendorIntel = "GenuineIntel"
	VendorAMD   = "AuthenticAMD"
)

const (
	VirtualizationFlagNone = ""
	VirtualizationFlagVMX  = "vmx"
	VirtualizationFlagSVM  = "svm"
)

type Info map[string]string

type Infos []Info

func uniquesOf(infos Infos, keys ...string) int {
	ret := map[string]struct{}{}
	for _, info := range infos {
		v := ""
		for _, key := range keys {
			v += info[key] + "\x00"
		}
		ret[v] = struct{}{}
	}
	return len(ret)
}

func ReadInfos() Infos {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	infos := make(Infos, 0, runtime.NumCPU())
	fileScanner := bufio.NewScanner(f)
	info := Info{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			infos = append(infos, info)
			info = Info{}
			continue
		}
		parts := strings.SplitN(fileScanner.Text(), ":", 2)
		info[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}
	return infos
}

func (c Infos) Sockets() int {
	return uniquesOf(c, "physical id")
}

func (c Infos) LogicalCores() int {
	return uniquesOf(c, "processor")
}

func (c Infos) PhysicalCores() int {
	return uniquesOf(c, "core id")
}

func (c Infos) Threads() int {
	sockets := c.Sockets()
	lc := c.LogicalCores()
	pc := c.PhysicalCores()
	threads := lc / sockets / pc
	// If we are on a irregular topology, such as lc=5 and pc=3, ensure hyper
	// threading is still detected
	if lc%(sockets*pc) > 0 {
		threads += 1
	}
	return threads
}

func (c Infos) ModelName() string {
	return c[0]["model name"]
}

func (c Infos) Vendor() string {
	return c[0]["vendor_id"]
}

func (c Infos) VirtualizationFlag() string {
	flags := c[0]["flags"]
	if strings.Contains(flags, " "+VirtualizationFlagVMX+" ") {
		return VirtualizationFlagVMX
	}
	if strings.Contains(flags, " "+VirtualizationFlagSVM+" ") {
		return VirtualizationFlagSVM
	}
	return VirtualizationFlagNone
}
