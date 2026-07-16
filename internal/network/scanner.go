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

func ReadSysDeviceId() (string, error) {
	content, err := os.ReadFile("/sys/class/eth0/device/devide_id")
	if err != nil {
		return "", err
	}

	line := string(content)
	fields := strings.Fields(line)

	var str string = fields[0]
	str = strings.Trim(str, "{}")

	return str, nil


}

func ReadSysClassId() (string, error) {
	content, err := os.ReadFile("/sys/class/eth0/device/class_id")
	if err != nil {
		return "", err
	}

	line := string(content)
	fields := strings.Fields(line)

	var str string = fields[0]
	str = strings.Trim(str, "{}")

	return str, nil


}

func ReadSysDriverOverride() (string, error) {
	content, err := os.ReadFile("/sys/class/eth0/device/driver_override")
	if err != nil {
		return "", err
	}

	line := string(content)
	fields := strings.Fields(line)

	var str string

	if len(fields) > 0 {
		str = fields[0]
	} else {
		str = ""
	}

	return str, nil
}

