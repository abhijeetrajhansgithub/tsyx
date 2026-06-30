package disk

import (
	"fmt"
	"strings"
)

func convertUnit(value uint64, unit string) string {
	unit = strings.ToLower(unit)

	valuef := float64(value)

	switch unit {
	case "b":
		return fmt.Sprintf("%.2f B", valuef)

	case "kb", "k":
		return fmt.Sprintf("%.2f KB", valuef/1024)

	case "mb", "m":
		return fmt.Sprintf("%.2f MB", valuef/(1024*1024))

	case "gb", "g":
		return fmt.Sprintf("%.2f GB", valuef/(1024*1024*1024))

	case "tb", "t":
		return fmt.Sprintf("%.2f TB", valuef/(1024*1024*1024*1024))

	case "auto":
		fallthrough
	default:
		switch {
		case value >= 1024*1024*1024*1024:
			return fmt.Sprintf("%.2f TB", valuef/(1024*1024*1024*1024))
		case value >= 1024*1024*1024:
			return fmt.Sprintf("%.2f GB", valuef/(1024*1024*1024))
		case value >= 1024*1024:
			return fmt.Sprintf("%.2f MB", valuef/(1024*1024))
		case value >= 1024:
			return fmt.Sprintf("%.2f KB", valuef/1024)
		default:
			return fmt.Sprintf("%.2f B", valuef)
		}
	}
}

func Format(diskInfo DiskInfo, unit string) string {
	return fmt.Sprintf(`
==================== DISK INFORMATION ====================

Storage
---------------------------------------------------------
Total          : %s
Used           : %s
Free           : %s
Used Percent   : %.2f%%

Blocks
---------------------------------------------------------
Block Size     : %.0f KB
Total Blocks   : %d
Free Blocks    : %d

Inodes
---------------------------------------------------------
Total Inodes   : %d
Free Inodes    : %d
Used Inodes    : %d

=========================================================
`,
		convertUnit(diskInfo.Total, unit),
		convertUnit(diskInfo.Used, unit),
		convertUnit(diskInfo.Free, unit),
		diskInfo.UsedPercent,

		float64(diskInfo.BlockSize)/1024,

		diskInfo.TotalBlocks,
		diskInfo.FreeBlocks,

		diskInfo.TotalInodes,
		diskInfo.FreeInodes,
		diskInfo.UsedInodes,
	)
}