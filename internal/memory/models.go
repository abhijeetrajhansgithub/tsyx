package memory

type MemoryInfo struct {
	// Main Memory
	MemTotal        uint64 // kB
	MemFree         uint64 // kB
	MemAvailable    uint64 // kB

	// Buffers & Cache
	Buffers         uint64 // kB
	Cached          uint64 // kB
	SwapCached      uint64 // kB

	// Active / Inactive
	Active          uint64 // kB
	Inactive        uint64 // kB
	ActiveAnon      uint64 // kB
	InactiveAnon    uint64 // kB
	ActiveFile      uint64 // kB
	InactiveFile    uint64 // kB

	// Locked Memory
	Unevictable     uint64 // kB
	Mlocked         uint64 // kB

	// Swap
	SwapTotal       uint64 // kB
	SwapFree        uint64 // kB

	// Dirty Pages
	Dirty           uint64 // kB
	Writeback       uint64 // kB
	WritebackTmp    uint64 // kB

	// Process Memory
	AnonPages       uint64 // kB
	Mapped          uint64 // kB
	Shmem           uint64 // kB

	// Kernel Memory
	KReclaimable    uint64 // kB
	Slab            uint64 // kB
	SReclaimable    uint64 // kB
	SUnreclaim      uint64 // kB
	KernelStack     uint64 // kB
	PageTables      uint64 // kB
	SecPageTables   uint64 // kB

	// Commit Accounting
	CommitLimit     uint64 // kB
	CommittedAS     uint64 // kB

	// Virtual Memory
	VmallocTotal    uint64 // kB
	VmallocUsed     uint64 // kB
	VmallocChunk    uint64 // kB

	// CPU-specific Memory
	Percpu          uint64 // kB

	// Hardware
	HardwareCorrupted uint64 // kB

	// Huge Pages
	AnonHugePages   uint64 // kB
	ShmemHugePages  uint64 // kB
	ShmemPmdMapped  uint64 // kB
	FileHugePages   uint64 // kB
	FilePmdMapped   uint64 // kB

	Unaccepted      uint64 // kB
	Balloon         uint64 // kB

	HugePagesTotal  uint64
	HugePagesFree   uint64
	HugePagesRsvd   uint64
	HugePagesSurp   uint64

	HugePageSize    uint64 // kB
	Hugetlb         uint64 // kB

	// Direct Mapping
	DirectMap4K     uint64 // kB
	DirectMap2M     uint64 // kB
	DirectMap1G     uint64 // kB
}