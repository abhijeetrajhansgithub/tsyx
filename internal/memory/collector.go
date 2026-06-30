package memory

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func parseKB (value string) (uint64, error) {
	fields := strings.Fields(value)

	if len(fields) != 2 {
		return 0, fmt.Errorf("invalid value: %s", value)
	}

	return strconv.ParseUint(fields[0], 10, 64)
}

func LinuxCollect() (MemoryInfo, error) {

	memInfo := MemoryInfo{}

	content, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return MemoryInfo{}, err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		v, err := parseKB(value)
		if err != nil {
			continue
		}

		switch key {

			case "MemTotal":
				memInfo.MemTotal = v
			case "MemFree":
				memInfo.MemFree = v
			case "MemAvailable":
				memInfo.MemAvailable = v

			case "Buffers":
				memInfo.Buffers = v
			case "Cached":
				memInfo.Cached = v
			case "SwapCached":
				memInfo.SwapCached = v

			case "Active":
				memInfo.Active = v
			case "Inactive":
				memInfo.Inactive = v
			case "Active(anon)":
				memInfo.ActiveAnon = v
			case "Inactive(anon)":
				memInfo.InactiveAnon = v
			case "Active(file)":
				memInfo.ActiveFile = v
			case "Inactive(file)":
				memInfo.InactiveFile = v

			case "Unevictable":
				memInfo.Unevictable = v
			case "Mlocked":
				memInfo.Mlocked = v

			case "SwapTotal":
				memInfo.SwapTotal = v
			case "SwapFree":
				memInfo.SwapFree = v

			case "Dirty":
				memInfo.Dirty = v
			case "Writeback":
				memInfo.Writeback = v
			case "WritebackTmp":
				memInfo.WritebackTmp = v

			case "AnonPages":
				memInfo.AnonPages = v
			case "Mapped":
				memInfo.Mapped = v
			case "Shmem":
				memInfo.Shmem = v

			case "KReclaimable":
				memInfo.KReclaimable = v
			case "Slab":
				memInfo.Slab = v
			case "SReclaimable":
				memInfo.SReclaimable = v
			case "SUnreclaim":
				memInfo.SUnreclaim = v

			case "KernelStack":
				memInfo.KernelStack = v
			case "PageTables":
				memInfo.PageTables = v
			case "SecPageTables":
				memInfo.SecPageTables = v

			case "CommitLimit":
				memInfo.CommitLimit = v
			case "Committed_AS":
				memInfo.CommittedAS = v

			case "VmallocTotal":
				memInfo.VmallocTotal = v
			case "VmallocUsed":
				memInfo.VmallocUsed = v
			case "VmallocChunk":
				memInfo.VmallocChunk = v

			case "Percpu":
				memInfo.Percpu = v

			case "HardwareCorrupted":
				memInfo.HardwareCorrupted = v

			case "AnonHugePages":
				memInfo.AnonHugePages = v
			case "ShmemHugePages":
				memInfo.ShmemHugePages = v
			case "ShmemPmdMapped":
				memInfo.ShmemPmdMapped = v
			case "FileHugePages":
				memInfo.FileHugePages = v
			case "FilePmdMapped":
				memInfo.FilePmdMapped = v

			case "Unaccepted":
				memInfo.Unaccepted = v
			case "Balloon":
				memInfo.Balloon = v

			case "HugePages_Total":
				memInfo.HugePagesTotal = v
			case "HugePages_Free":
				memInfo.HugePagesFree = v
			case "HugePages_Rsvd":
				memInfo.HugePagesRsvd = v
			case "HugePages_Surp":
				memInfo.HugePagesSurp = v

			case "Hugepagesize":
				memInfo.HugePageSize = v
			case "Hugetlb":
				memInfo.Hugetlb = v

			case "DirectMap4k":
				memInfo.DirectMap4K = v
			case "DirectMap2M":
				memInfo.DirectMap2M = v
			case "DirectMap1G":
				memInfo.DirectMap1G = v
			}
	}

	return memInfo, nil

}
