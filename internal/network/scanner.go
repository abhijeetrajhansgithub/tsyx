package net

import (
	"os"
	"strings"
)

const sysDevicePath = "/sys/class/eth0/device/"

func readSysFile(name string, trimBraces bool) (string, error) {
	content, err := os.ReadFile(sysDevicePath + name)
	if err != nil {
		return "", err
	}

	s := strings.TrimSpace(string(content))

	if trimBraces {
		s = strings.Trim(s, "{}")
	}

	return s, nil
}

func ReadSysId() (string, error) {
	return readSysFile("id", false)
}

func ReadSysDeviceId() (string, error) {
	return readSysFile("device_id", true)
}

func ReadSysClassId() (string, error) {
	return readSysFile("class_id", true)
}

func ReadSysDriverOverride() (string, error) {
	return readSysFile("driver_override", false)
}

func ReadSysModalias() (string, error) {
	return readSysFile("modalias", false)
}

func ReadSysNumaNode() (string, error) {
	return readSysFile("numa_node", false)
}

func ReadSysState() (string, error) {
	return readSysFile("state", false)
}

func ReadSysVendor() (string, error) {
	return readSysFile("vendor", false)
}

