// internal/cpu/models.go
package cpu

type CPUInfo struct {
	Model             string
	Vendor            string
	Architecture      string

	PhysicalCores     int
	LogicalCores      int

	CurrentFrequency  float64
	MaxFrequency      float64

	Usage             float64
	Temperature       float64

	CacheSizeKB       string

	CPUFamily         int
	ModelID           int
	Stepping          int

	Flags             []string
}