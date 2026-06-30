package disk

import "golang.org/x/sys/unix"

func LinuxCollect() (DiskInfo, error) {
	var stat unix.Statfs_t

	// Get filesystem statistics for the root filesystem
	if err := unix.Statfs("/", &stat); err != nil {
		return DiskInfo{}, err
	}

	// Linux recommends using Frsize (fundamental block size)
	// Fall back to Bsize if Frsize is unavailable.
	blockSize := uint64(stat.Frsize)
	if blockSize == 0 {
		blockSize = uint64(stat.Bsize)
	}

	// Storage (bytes)
	total := stat.Blocks * blockSize
	free := stat.Bfree * blockSize
	used := total - free

	// Usage percentage
	var usedPercent float64
	if total > 0 {
		usedPercent = (float64(used) / float64(total)) * 100
	}

	// Inodes
	totalInodes := stat.Files
	freeInodes := stat.Ffree
	usedInodes := totalInodes - freeInodes

	return DiskInfo{
		// Storage
		Total: total,
		Used:  used,
		Free:  free,

		// Usage
		UsedPercent: usedPercent,

		// Blocks
		BlockSize:   blockSize,
		TotalBlocks: stat.Blocks,
		FreeBlocks:  stat.Bfree,

		// Inodes
		TotalInodes: totalInodes,
		FreeInodes:  freeInodes,
		UsedInodes:  usedInodes,
	}, nil
}