package net

import (
	"os"
	"strings"
	"path/filepath"
)

const sysDevicePath = "/sys/class/net/eth0/device/"

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

// ----------------------------------------------

func ReadSysPower() (DevicePowerInfo, error) {
	dpInfo := DevicePowerInfo{}

	// control
	control, err := readSysFile(
		"power/control",
		false,
	)

	if err != nil {
		return dpInfo, err
	}

	// runtime_active_time
	runtime_active_time, err := readSysFile(
		"power/runtime_active_time",
		false,
	)

	if err != nil {
		return dpInfo, err
	}

	// runtime_status
	runtime_status, err := readSysFile(
		"power/runtime_status",
		false,
	)

	if err != nil {
		return dpInfo, err
	}

	// runtime_suspended_time
	runtime_suspended_time, err := readSysFile(
		"power/runtime_suspended_time",
		false,
	)

	if err != nil {
		return dpInfo, err
	}

	dpInfo.Control = control
	dpInfo.RuntimeActiveTime = runtime_active_time
	dpInfo.RuntimeStatus = runtime_status
	dpInfo.RuntimeSuspendedTime = runtime_suspended_time

	return dpInfo, nil

}


func ReadSysSubsystem() (DeviceSubsystemInfo, error) {
	dsinfo := DeviceSubsystemInfo{}

	drivers_autoprobe, err := readSysFile(
		"subsystem/drivers_autoprobe",
		false,
	)

	if err != nil {
		return dsinfo, err
	}

	hibernation, err := readSysFile(
		"subsystem/hibernation",
		false,
	)

	if err != nil {
		return dsinfo, err
	}

	dsinfo.DriversAutoprobe = drivers_autoprobe
	dsinfo.Hibernation = hibernation

	return dsinfo, nil
}


func ReadSysUevent() (DeviceUeventInfo, error) {
	uevent := DeviceUeventInfo{}

	// TODO: Replace "uevent" with the absolute path, e.g., "/sys/class/net/eth0/device/uevent"
	data, err := os.ReadFile("uevent")
	if err != nil {
		return uevent, err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		// Skip empty lines (like the trailing newline at the end of the file)
		if line == "" {
			continue
		}

		fields := strings.Split(line, "=")

		// Prevent index out of bounds panic if a line doesn't contain an "="
		if len(fields) != 2 {
			continue
		}

		// Correct way to make a string uppercase in Go
		key := strings.ToUpper(fields[0])
		val := fields[1]

		if key == "DRIVER" {
			uevent.Driver = val
		} else if key == "MODALIAS" {
			uevent.Modalias = val
		}
	}

	return uevent, nil
}

func ReadSysMonitor() (DeviceMonitorInfo, error) {
	dmi := DeviceMonitorInfo{}

	client_monitor_conn_id, err := readSysFile(
		"client_monitor_conn_id",
		false,
	)

	if err != nil {
		return dmi, err
	}

	client_monitor_latency, err := readSysFile(
		"client_monitor_latency",
		false,
	)

	if err != nil {
		return dmi, err
	}

	client_monitor_pending, err := readSysFile(
		"client_monitor_pending",
		false,
	)

	if err != nil {
		return dmi, err
	}

	server_monitor_conn_id, err := readSysFile(
		"server_monitor_conn_id",
		false,
	)

	if err != nil {
		return dmi, err
	}

	server_monitor_latency, err := readSysFile(
		"server_monitor_latency",
		false,
	)

	if err != nil {
		return dmi, err
	}

	server_monitor_pending, err := readSysFile(
		"server_monitor_pending",
		false,
	)

	if err != nil {
		return dmi, err
	}

	// assign values if no error
	dmi.ClientConnectionID = client_monitor_conn_id
	dmi.ClientLatency = client_monitor_latency
	dmi.ClientPending = client_monitor_pending
	dmi.ServerConnectionID = server_monitor_conn_id
	dmi.ServerLatency = server_monitor_latency
	dmi.ServerPending = server_monitor_pending

	return dmi, nil
}

func ReadSysRingBuffer() (DeviceRingBuffer, error) {
	drb := DeviceRingBuffer{}

	in_intr_mask, err := readSysFile(
		"in_intr_mask",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_read_bytes_avail, err := readSysFile(
		"in_read_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_read_index, err := readSysFile(
		"in_read_index",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_write_bytes_avail, err := readSysFile(
		"in_write_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_write_index, err := readSysFile(
		"in_write_index",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_intr_mask, err := readSysFile(
		"out_intr_mask",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_read_bytes_avail, err := readSysFile(
		"out_read_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_read_index, err := readSysFile(
		"out_read_index",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_write_bytes_avail, err := readSysFile(
		"out_write_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_write_index, err := readSysFile(
		"out_write_index",
		false,
	)

	if err != nil {
		return drb, err
	}

	// assign the values
	drb.InInterruptMask = in_intr_mask
	drb.InReadBytesAvail = in_read_bytes_avail
	drb.InReadIndex = in_read_index
	drb.InWriteBytesAvail = in_write_bytes_avail
	drb.InWriteIndex = in_write_index

	drb.OutInterruptMask = out_intr_mask
	drb.OutReadBytesAvail = out_read_bytes_avail
	drb.OutReadIndex = out_read_index
	drb.OutWriteBytesAvail = out_write_bytes_avail
	drb.OutWriteIndex = out_write_index


	return drb, nil
}

func readSysInterfaceChannelElements(key string, path string) (string, error) {
	content, err := os.ReadFile(path + "/" + key)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}

func ReadSysInterfaceChannel() (map[string]InterfaceChannel, error) {
	interfaceMap := make(map[string]InterfaceChannel)

	path := "/sys/class/net/eth0/device/channels"

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		innerDirName := entry.Name()
		innerDirPath := filepath.Join(path, innerDirName)

		_cpu, err := readSysInterfaceChannelElements("cpu", innerDirPath)
		if err != nil {
			return nil, err
		}

		_events, err := readSysInterfaceChannelElements("events", innerDirPath)
		if err != nil {
			return nil, err
		}

		_inMask, err := readSysInterfaceChannelElements("in_mask", innerDirPath)
		if err != nil {
			return nil, err
		}

		_interrupts, err := readSysInterfaceChannelElements("interrupts", innerDirPath)
		if err != nil {
			return nil, err
		}

		_intrInFull, err := readSysInterfaceChannelElements("intr_in_full", innerDirPath)
		if err != nil {
			return nil, err
		}

		_intrOutEmpty, err := readSysInterfaceChannelElements("intr_out_empty", innerDirPath)
		if err != nil {
			return nil, err
		}

		_latency, err := readSysInterfaceChannelElements("latency", innerDirPath)
		if err != nil {
			return nil, err
		}

		_monitorID, err := readSysInterfaceChannelElements("monitor_id", innerDirPath)
		if err != nil {
			return nil, err
		}

		_outFullFirst, err := readSysInterfaceChannelElements("out_full_first", innerDirPath)
		if err != nil {
			return nil, err
		}

		_outFullTotal, err := readSysInterfaceChannelElements("out_full_total", innerDirPath)
		if err != nil {
			return nil, err
		}

		_outMask, err := readSysInterfaceChannelElements("out_mask", innerDirPath)
		if err != nil {
			return nil, err
		}

		_pending, err := readSysInterfaceChannelElements("pending", innerDirPath)
		if err != nil {
			return nil, err
		}

		_readAvail, err := readSysInterfaceChannelElements("read_avail", innerDirPath)
		if err != nil {
			return nil, err
		}

		_subchannelID, err := readSysInterfaceChannelElements("subchannel_id", innerDirPath)
		if err != nil {
			return nil, err
		}

		_writeAvail, err := readSysInterfaceChannelElements("write_avail", innerDirPath)
		if err != nil {
			return nil, err
		}

		channel := InterfaceChannel{
			CPU:          _cpu,
			Events:       _events,
			InMask:       _inMask,
			Interrupts:   _interrupts,
			IntrInFull:   _intrInFull,
			IntrOutEmpty: _intrOutEmpty,
			Latency:      _latency,
			MonitorID:    _monitorID,
			OutFullFirst: _outFullFirst,
			OutFullTotal: _outFullTotal,
			OutMask:      _outMask,
			Pending:      _pending,
			ReadAvail:    _readAvail,
			SubchannelID: _subchannelID,
			WriteAvail:   _writeAvail,
		}

		interfaceMap[innerDirName] = channel
	}

	return interfaceMap, nil
}

// =============================================================================================
// =============================================================================================
// =============================================================================================
// Generic Data Reader Function

func ReadGenericData(path, file string) (string, error) {
	content, err := os.ReadFile(filepath.Join(path, file))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}