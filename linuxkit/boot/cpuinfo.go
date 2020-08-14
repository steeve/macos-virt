package boot

import (
	"bufio"
	"os"
	"strings"
)

type cpuInfo struct {
	Sockets int
	Cores   int
	Threads int
	Virt    string
}

func readCpuInfo() cpuInfo {
	f, err := os.Open("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	info := cpuInfo{}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		parts := strings.SplitN(fileScanner.Text(), ":", 2)
		switch strings.TrimSpace(parts[0]) {
		case "processor":

		}
	}
}
