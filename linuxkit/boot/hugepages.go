package boot

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func hugePageSize() int64 {
	data, err := ioutil.ReadFile("/sys/kernel/mm/transparent_hugepage/hpage_pmd_size")
	if err != nil {
		return 0
	}
	size := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
	if size&(size-1) != 0 {
		// size is not a power of 2
		return 0
	}
	return size
}

func hugePagesForMemory(memory int64) int64 {
	int64 hpSize := hugePageSize()
	if hpSize == 0 {
		return 0
	}
	return memory / hugePageSize + memory % hugePageSize
}
