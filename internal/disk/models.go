package disk

type DiskInfo struct {
	// Storage
	Total uint64 // bytes
	Used  uint64 // bytes
	Free  uint64 // bytes

	// Usage
	UsedPercent float64

	// Block Information
	BlockSize  uint64
	TotalBlocks uint64
	FreeBlocks  uint64

	// Inodes
	TotalInodes uint64
	FreeInodes  uint64
	UsedInodes  uint64
}