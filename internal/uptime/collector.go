package uptime

import (
	"os"
	"strconv"
	"strings"
)

func LinuxCollect() (DeviceUptime, error) {
	// Read the file
	content, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return DeviceUptime{}, err
	}

	// Split into fields (handles multiple spaces/newlines)
	fields := strings.Fields(string(content))

	// Parse uptime in seconds
	seconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return DeviceUptime{}, err
	}

	totalSeconds := int(seconds)

	uptime := DeviceUptime{
		Days:    totalSeconds / 86400,
		Hours:   (totalSeconds / 3600) % 24,
		Minutes: (totalSeconds / 60) % 60,
		Seconds: totalSeconds % 60,
	}

	return uptime, nil
}