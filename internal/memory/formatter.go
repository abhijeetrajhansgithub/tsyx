package memory

import (
	"fmt"
	"strings"
)
func convertUnit(value uint64, unit string) string {
	unit = strings.ToLower(unit)

	valuef := float64(value)

	switch unit {
	case "b":
		return fmt.Sprintf("%.2f B", valuef*1024)

	case "kb", "k":
		return fmt.Sprintf("%.2f KB", valuef)

	case "mb", "m":
		return fmt.Sprintf("%.2f MB", valuef/1024)

	case "gb", "g":
		return fmt.Sprintf("%.2f GB", valuef/(1024*1024))

	case "tb", "t":
		return fmt.Sprintf("%.2f TB", valuef/(1024*1024*1024))

	case "auto":
		fallthrough
	default:
		switch {
		case value >= 1024*1024*1024:
			return fmt.Sprintf("%.2f TB", valuef/(1024*1024*1024))
		case value >= 1024*1024:
			return fmt.Sprintf("%.2f GB", valuef/(1024*1024))
		case value >= 1024:
			return fmt.Sprintf("%.2f MB", valuef/1024)
		default:
			return fmt.Sprintf("%.2f KB", valuef)
		}
	}
}

func FormatSummary(memInfo MemoryInfo, unit string) string {
	used := memInfo.MemTotal - memInfo.MemAvailable
	swapUsed := memInfo.SwapTotal - memInfo.SwapFree
	usage := (float64(used) / float64(memInfo.MemTotal)) * 100

	return fmt.Sprintf(`
================ MEMORY SUMMARY ================

Total          : %s
Used           : %s
Free           : %s
Available      : %s
Usage          : %.2f%%

Swap Total     : %s
Swap Used      : %s
Swap Free      : %s

===============================================
`,
		convertUnit(memInfo.MemTotal, unit),
		convertUnit(used, unit),
		convertUnit(memInfo.MemFree, unit),
		convertUnit(memInfo.MemAvailable, unit),
		usage,
		convertUnit(memInfo.SwapTotal, unit),
		convertUnit(swapUsed, unit),
		convertUnit(memInfo.SwapFree, unit),
	)
}

func FormatDetailed(memInfo MemoryInfo, unit string) string {
	return fmt.Sprintf(`
================ MEMORY DETAILED ================

%s

Buffers         : %s
Cached          : %s
Swap Cached     : %s

Active          : %s
Inactive        : %s

Active (Anon)   : %s
Inactive (Anon) : %s

Active (File)   : %s
Inactive (File) : %s

Dirty           : %s
Writeback       : %s

Mapped          : %s
Shared Memory   : %s

Commit Limit    : %s
Committed AS    : %s

=================================================
`,
		strings.TrimSpace(FormatSummary(memInfo, unit)),
		convertUnit(memInfo.Buffers, unit),
		convertUnit(memInfo.Cached, unit),
		convertUnit(memInfo.SwapCached, unit),
		convertUnit(memInfo.Active, unit),
		convertUnit(memInfo.Inactive, unit),
		convertUnit(memInfo.ActiveAnon, unit),
		convertUnit(memInfo.InactiveAnon, unit),
		convertUnit(memInfo.ActiveFile, unit),
		convertUnit(memInfo.InactiveFile, unit),
		convertUnit(memInfo.Dirty, unit),
		convertUnit(memInfo.Writeback, unit),
		convertUnit(memInfo.Mapped, unit),
		convertUnit(memInfo.Shmem, unit),
		convertUnit(memInfo.CommitLimit, unit),
		convertUnit(memInfo.CommittedAS, unit),
	)
}

func FormatKernel(memInfo MemoryInfo, unit string) string {
	return fmt.Sprintf(`
================ MEMORY KERNEL ================

%s

Slab            : %s
SReclaimable    : %s
SUnreclaim      : %s
KReclaimable    : %s

Kernel Stack    : %s
Page Tables     : %s
SecPageTables   : %s

Vmalloc Total   : %s
Vmalloc Used    : %s
Vmalloc Chunk   : %s

Percpu          : %s

HugePages Total : %d
HugePages Free  : %d
HugePages Rsvd  : %d
HugePages Surp  : %d
HugePage Size   : %s
Hugetlb         : %s

DirectMap4K     : %s
DirectMap2M     : %s
DirectMap1G     : %s

===============================================
`,
		strings.TrimSpace(FormatDetailed(memInfo, unit)),
		convertUnit(memInfo.Slab, unit),
		convertUnit(memInfo.SReclaimable, unit),
		convertUnit(memInfo.SUnreclaim, unit),
		convertUnit(memInfo.KReclaimable, unit),
		convertUnit(memInfo.KernelStack, unit),
		convertUnit(memInfo.PageTables, unit),
		convertUnit(memInfo.SecPageTables, unit),
		convertUnit(memInfo.VmallocTotal, unit),
		convertUnit(memInfo.VmallocUsed, unit),
		convertUnit(memInfo.VmallocChunk, unit),
		convertUnit(memInfo.Percpu, unit),
		memInfo.HugePagesTotal,
		memInfo.HugePagesFree,
		memInfo.HugePagesRsvd,
		memInfo.HugePagesSurp,
		convertUnit(memInfo.HugePageSize, unit),
		convertUnit(memInfo.Hugetlb, unit),
		convertUnit(memInfo.DirectMap4K, unit),
		convertUnit(memInfo.DirectMap2M, unit),
		convertUnit(memInfo.DirectMap1G, unit),
	)
}
