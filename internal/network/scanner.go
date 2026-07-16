package net

import (
	"os"
	"strings"
)


func ReadSysId() (string, error){
	content, err := os.ReadFile("/sys/class/eth0/device/id")
	if err != nil {
		return "", err
	}

	line := string(content)
	fields := strings.Fields(line)

	return fields[0], nil
}