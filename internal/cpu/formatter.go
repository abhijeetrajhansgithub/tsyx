package cpu

import "fmt"

func Format(cpuInfo CPUInfo) string {
	return fmt.Sprintf(`
==================== CPU INFORMATION ====================

Hardware
---------------------------------------------------------
Model               : %s
Vendor              : %s
Architecture        : %s

Topology
---------------------------------------------------------
Physical Cores      : %d
Logical Cores       : %d

Performance
---------------------------------------------------------
Current Frequency   : %.2f MHz
Max Frequency       : %.2f MHz
Usage               : %.2f%%
Temperature         : %.2f °C

Cache
---------------------------------------------------------
Cache Size            : %s

=========================================================
`,
		cpuInfo.Model,
		cpuInfo.Vendor,
		cpuInfo.Architecture,
		cpuInfo.PhysicalCores,
		cpuInfo.LogicalCores,
		cpuInfo.CurrentFrequency,
		cpuInfo.MaxFrequency,
		cpuInfo.Usage,
		cpuInfo.Temperature,
		cpuInfo.CacheSizeKB,
	)
}