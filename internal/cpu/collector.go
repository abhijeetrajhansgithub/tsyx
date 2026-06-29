// internal/cpu/collector.go
package cpu

import (
	"os"
	"log"
	"strings"
	"strconv"
)

func LinuxCollect() CPUInfo {

	// initialize CPUInfo
	cpuInfo := CPUInfo{}

	// read file and parse it
	content, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
			case "model name":
				cpuInfo.Model = value
			case "vendor_id":
				cpuInfo.Vendor = value
			case "cpu family":
				cpuInfo.CPUFamily, _ = strconv.Atoi(value)
			case "cpu cores":
				cpuInfo.PhysicalCores, _ = strconv.Atoi(value)
			case "siblings":
				cpuInfo.LogicalCores, _ = strconv.Atoi(value)
			case "cpu MHz":
				cpuInfo.CurrentFrequency, _ = strconv.ParseFloat(value, 64)
			case "cache size":
				cpuInfo.CacheSizeKB = value
			case "flags":
				cpuInfo.Flags = strings.Split(value, " ")
		}
	}

	return cpuInfo
}

