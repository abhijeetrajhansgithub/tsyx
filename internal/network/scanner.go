package net

import (
	"os"
	"fmt"
	"strings"
	"path/filepath"
)

// const sysClassNetPath = "/sys/class/net/eth0/device/"
const sysClassNetPath = "/sys/class/net/"

func readSysFile(dir string, name string, trimBraces bool) (string, error) {
	content, err := os.ReadFile(sysClassNetPath + dir + "/device/" + name)
	if err != nil {
		return "", err
	}

	s := strings.TrimSpace(string(content))

	if trimBraces {
		s = strings.Trim(s, "{}")
	}

	return s, nil
}

func ReadSysId(dir string) (string, error) {
	return readSysFile(dir, "id", false)
}

func ReadSysDeviceId(dir string) (string, error) {
	return readSysFile(dir, "device_id", true)
}

func ReadSysClassId(dir string) (string, error) {
	return readSysFile(dir, "class_id", true)
}

func ReadSysDriverOverride(dir string) (string, error) {
	return readSysFile(dir, "driver_override", false)
}

func ReadSysModalias(dir string) (string, error) {
	return readSysFile(dir, "modalias", false)
}

func ReadSysNumaNode(dir string) (string, error) {
	return readSysFile(dir, "numa_node", false)
}

func ReadSysState(dir string) (string, error) {
	return readSysFile(dir, "state", false)
}

func ReadSysVendor(dir string) (string, error) {
	return readSysFile(dir, "vendor", false)
}

// ----------------------------------------------

func ReadSysPower(dir string) (DevicePowerInfo, error) {
	dpInfo := DevicePowerInfo{}

	// control
	control, err := readSysFile(
		dir, 
		"power/control",
		false,
	)

	if err != nil {
		return dpInfo, err
	}

	// runtime_active_time
	runtime_active_time, err := readSysFile(
		dir,
		"power/runtime_active_time",
		false,
	)

	if err != nil {
		return dpInfo, err
	}

	// runtime_status
	runtime_status, err := readSysFile(
		dir,
		"power/runtime_status",
		false,
	)

	if err != nil {
		return dpInfo, err
	}

	// runtime_suspended_time
	runtime_suspended_time, err := readSysFile(
		dir,
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


func ReadSysSubsystem(dir string) (DeviceSubsystemInfo, error) {
	dsinfo := DeviceSubsystemInfo{}

	drivers_autoprobe, err := readSysFile(
		dir,
		"subsystem/drivers_autoprobe",
		false,
	)

	if err != nil {
		return dsinfo, err
	}

	hibernation, err := readSysFile(
		dir,
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


func ReadSysUevent(dir string) (DeviceUeventInfo, error) {
	uevent := DeviceUeventInfo{}

	// TODO: Replace "uevent" with the absolute path, e.g., "/sys/class/net/eth0/device/uevent"
	data, err := os.ReadFile(sysClassNetPath + dir + "/device/uevent")
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

func ReadSysMonitor(dir string) (DeviceMonitorInfo, error) {
	dmi := DeviceMonitorInfo{}

	client_monitor_conn_id, err := readSysFile(
		dir,
		"client_monitor_conn_id",
		false,
	)

	if err != nil {
		return dmi, err
	}

	client_monitor_latency, err := readSysFile(
		dir,
		"client_monitor_latency",
		false,
	)

	if err != nil {
		return dmi, err
	}

	client_monitor_pending, err := readSysFile(
		dir,
		"client_monitor_pending",
		false,
	)

	if err != nil {
		return dmi, err
	}

	server_monitor_conn_id, err := readSysFile(
		dir,
		"server_monitor_conn_id",
		false,
	)

	if err != nil {
		return dmi, err
	}

	server_monitor_latency, err := readSysFile(
		dir,
		"server_monitor_latency",
		false,
	)

	if err != nil {
		return dmi, err
	}

	server_monitor_pending, err := readSysFile(
		dir,
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

func ReadSysRingBuffer(dir string) (DeviceRingBuffer, error) {
	drb := DeviceRingBuffer{}

	in_intr_mask, err := readSysFile(
		dir,
		"in_intr_mask",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_read_bytes_avail, err := readSysFile(
		dir,
		"in_read_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_read_index, err := readSysFile(
		dir,
		"in_read_index",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_write_bytes_avail, err := readSysFile(
		dir,
		"in_write_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	in_write_index, err := readSysFile(
		dir,
		"in_write_index",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_intr_mask, err := readSysFile(
		dir,
		"out_intr_mask",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_read_bytes_avail, err := readSysFile(
		dir,
		"out_read_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_read_index, err := readSysFile(
		dir,
		"out_read_index",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_write_bytes_avail, err := readSysFile(
		dir,
		"out_write_bytes_avail",
		false,
	)

	if err != nil {
		return drb, err
	}

	out_write_index, err := readSysFile(
		dir,
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

func ReadSysInterfaceChannel(dir string) (map[string]InterfaceChannel, error) {
	interfaceMap := make(map[string]InterfaceChannel)

	path := "/sys/class/net/" + dir + "/device/channels"

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

func ReadGenericData(path string, file string) (string, error) {
	content, err := os.ReadFile(filepath.Join(path, file))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}

// =============================================================================================
// =============================================================================================
// =============================================================================================
// Fib Trie Reader Function

func ReadFibTrieSection(section string, content string) (FIBTrieSection, error) {
	fib := FIBTrieSection{}

	var sectionContent string

	switch section {
	case "Main":
		parts := strings.Split(content, "Main:")
		if len(parts) < 2 {
			return fib, fmt.Errorf("Main section not found")
		}

		sectionContent = strings.Split(parts[1], "Local:")[0]

	case "Local":
		parts := strings.Split(content, "Local:")
		if len(parts) < 2 {
			return fib, fmt.Errorf("Local section not found")
		}

		sectionContent = parts[1]

	default:
		return fib, fmt.Errorf("unknown section %q", section)
	}

	parts := strings.Split(sectionContent, "Counters:")
	if len(parts) != 2 {
		return fib, fmt.Errorf("Counters section missing")
	}

	mainPart := parts[0]
	counterPart := parts[1]

	// ---------------- Main ----------------
	for _, line := range strings.Split(mainPart, "\n") {
		line = strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(line, "Aver depth:"):
			fib.AverageDepth = strings.Fields(line)[2]

		case strings.HasPrefix(line, "Max depth:"):
			fib.MaxDepth = strings.Fields(line)[2]

		case strings.HasPrefix(line, "Leaves:"):
			fib.Leaves = strings.Fields(line)[1]

		case strings.HasPrefix(line, "Prefixes:"):
			fib.Prefixes = strings.Fields(line)[1]

		case strings.HasPrefix(line, "Internal nodes:"):
			fib.InternalNodes = strings.Fields(line)[2]

		case strings.HasPrefix(line, "1:"):
			fields := strings.Fields(line)
			// 1: 1  2: 4  3: 1
			fib.Depth1 = fields[1]
			fib.Depth2 = fields[3]
			fib.Depth3 = fields[5]

		case strings.HasPrefix(line, "Pointers:"):
			fib.Pointers = strings.Fields(line)[1]

		case strings.HasPrefix(line, "Null ptrs:"):
			fib.NullPtrs = strings.Fields(line)[2]

		case strings.HasPrefix(line, "Total size:"):
			fields := strings.Fields(line)
			fib.TotalSizeKB = fields[2]
		}
	}

	// ---------------- Counters ----------------

	for _, line := range strings.Split(counterPart, "\n") {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "---------") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		switch key {
		case "gets":
			fib.Counters.Gets = value

		case "backtracks":
			fib.Counters.Backtracks = value

		case "semantic match passed":
			fib.Counters.SemanticMatchPassed = value

		case "semantic match miss":
			fib.Counters.SemanticMatchMiss = value

		case "null node hit":
			fib.Counters.NullNodeHit = value

		case "skipped node resize":
			fib.Counters.SkippedNodeResize = value
		}
	}

	return fib, nil
}