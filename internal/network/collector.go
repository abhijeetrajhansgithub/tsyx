package net

import (
	// "fmt"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var DirMap map[string]string = map[string]string{
	// =========================
	// /proc/net
	// =========================
	"proc_arp":        "/proc/net/arp",
	"proc_dev":        "/proc/net/dev",
	"proc_if_inet6":   "/proc/net/if_inet6",
	"proc_ptype":      "/proc/net/ptype",
	"proc_igmp":       "/proc/net/igmp",
	"proc_igmp6":      "/proc/net/igmp6",
	"proc_route":      "/proc/net/route",
	"proc_rt_cache":   "/proc/net/rt_cache",
	"proc_fib_triestat": "/proc/net/fib_triestat",
	"proc_tcp":        "/proc/net/tcp",
	"proc_tcp6":       "/proc/net/tcp6",
	"proc_udp":        "/proc/net/udp",
	"proc_udp6":       "/proc/net/udp6",
	"proc_unix":       "/proc/net/unix",
	"proc_packet":     "/proc/net/packet",
	"proc_netlink":    "/proc/net/netlink",
	"proc_connector":  "/proc/net/connector",
	"proc_netstat":    "/proc/net/netstat",
	"proc_snmp":       "/proc/net/snmp",
	"proc_snmp6":      "/proc/net/snmp6",
	"proc_sockstat":   "/proc/net/sockstat",
	"proc_sockstat6":  "/proc/net/sockstat6",
	"proc_protocols":  "/proc/net/protocols",
	"proc_tls_stat":   "/proc/net/tls_stat",
	"proc_xfrm_stat":  "/proc/net/xfrm_stat",

	// =========================
	// /sys
	// =========================
	"sys_class_net": "/sys/class/net",

	// =========================
	// /etc
	// =========================
	"etc_os_release": "/etc/os-release",
	"etc_hostname":   "/etc/hostname",
	"etc_hosts":      "/etc/hosts",
	"etc_resolv_conf": "/etc/resolv.conf",
	"etc_nsswitch":   "/etc/nsswitch.conf",
	"etc_passwd":     "/etc/passwd",
	"etc_shells":     "/etc/shells",
	"etc_protocols":  "/etc/protocols",
	"etc_services":   "/etc/services",
}

func LinuxCollectTCP(key string) (TCPConnectionTable, error) {  // key: tcp ot tcp6
	tcp := TCPConnectionTable{}

	return tcp, nil
}

func LinuxCollectFIBTrieStatistics(key string) (FIBTrieStatistics, error) {
	if key != "proc_fib_triestat" {
		return FIBTrieStatistics{}, fmt.Errorf("wrong key for LinuxCollectFIBTrieStatistics")
	}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return FIBTrieStatistics{}, err
	}

	var fib FIBTrieStatistics
	data := string(content)

	// Parse basic info
	_, err = fmt.Sscanf(
		data,
		"Basic info: size of leaf: %s bytes, size of tnode: %s bytes.",
		&fib.LeafSizeBytes,
		&fib.TNodeSizeBytes,
	)
	if err != nil {
		return FIBTrieStatistics{}, err
	}

	// Parse Main section
	fib.Main, err = ReadFibTrieSection("Main", data)
	if err != nil {
		return FIBTrieStatistics{}, err
	}

	// Parse Local section
	fib.Local, err = ReadFibTrieSection("Local", data)
	if err != nil {
		return FIBTrieStatistics{}, err
	}

	return fib, nil
}

func LinuxCollectRouteCacheTable(key string) (RouteCacheTable, error) {
	routeCacheTable := RouteCacheTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return routeCacheTable, err
	}

	lines := strings.Split(string(content), "\n") 

	for i, line := range lines {
		if i == 0 {
			continue
		}

		fields := strings.Fields(line)

		entry := RouteCacheEntry{
			Interface: fields[0],
			Destination: fields[1],
			Gateway: fields[2],
			Flags: fields[3],
			RefCount: fields[4],
			Use: fields[5],
			Metric: fields[6],
			Source: fields[7],
			MTU: fields[8],
			Window: fields[9],
			IRTT: fields[10],
			TOS: fields[11],
			HardwareHeaderRef: fields[12],
			HardwareHeaderUpToDate: fields[13],
			SpecificDestination: fields[14],
		}

		routeCacheTable.Routes = append(routeCacheTable.Routes, entry)
	}

	return routeCacheTable, nil
}

func LinuxCollectRouteTable(key string) (RouteTable, error) {
	routeTable := RouteTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return routeTable, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i == 0 {
			continue
		}
		// TODO
		fields := strings.Fields(line)

		entry := RouteEntry{
			Interface: fields[0],
			Destination: fields[1],
			Gateway: fields[2],
			Flags: fields[3],
			RefCount: fields[4],
			Use: fields[5],
			Metric: fields[6],
			Mask: fields[7],
			MTU: fields[8],
			Window: fields[9],
			IRTT: fields[10],
		}

		routeTable.Routes = append(routeTable.Routes, entry)

	}

	return routeTable, nil
}

func LinuxCollectNetworkInterface(key string) (NetworkInterfaces, error) {
	netif := NetworkInterfaces{}
	//TODO

	fmt.Println("[ LinuxCollectNetworkInterface ] KEY: " + key)

	// get the directories
	basePath := "/sys/class/net"

	// inside folders: eth0, lo
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return netif, err
	}

	for _, dir := range entries {
		if !dir.IsDir() {
			continue
		}

		fullDirPath := filepath.Join(basePath, dir.Name())

		name := dir.Name()

		addr_assign_type, err := ReadGenericData(fullDirPath, "addr_assign_type")
		if err != nil {
			return netif, err
		}

		addr_len, err := ReadGenericData(fullDirPath, "addr_len")
		if err != nil {
			return netif, err
		}

		address, err := ReadGenericData(fullDirPath, "address")
		if err != nil {
			return netif, err
		}

		broadcast, err := ReadGenericData(fullDirPath, "broadcast")
		if err != nil {
			return netif, err
		}

		carrier_info, err := LinucCollectCarrierInfo(name)
		if err != nil {
			return netif, err
		}

		dev_id, err := ReadGenericData(fullDirPath, "dev_id")
		if err != nil {
			return netif, err
		}

		dev_port, err := ReadGenericData(fullDirPath, "dev_port") 
		if err != nil {
			return netif, err
		}

		device, err := LinuxCollectInterfaceDevice(name)
		if err != nil {
			return netif, err
		}

		entry := NetworkInterface{
			Name: name,
			AddressAssignType: addr_assign_type,
			AddressLength: addr_len,
			Address: address,
			Broadcast: broadcast,
			Carrier: carrier_info,
			DevID: dev_id,
			DevPort: dev_port,
			Device: device,

		}

		netif.Interfaces = append(netif.Interfaces, entry)

	}

	return netif, nil

}

func LinucCollectCarrierInfo(dir string) (CarrierInfo, error) {
	carrier := CarrierInfo{}

	basePath := "/sys/class/net/" + dir

	status, err := os.ReadFile(filepath.Join(basePath, "carrier"))
	if err != nil {
		return carrier, err
	}

	carrier_changes, err := os.ReadFile(filepath.Join(basePath, "carrier_changes"))
	if err != nil {
		return carrier, err
	}

	carrier_up_count, err := os.ReadFile(filepath.Join(basePath, "carrier_up_count"))
	if err != nil {
		return carrier, err
	}

	carrier_down_count, err := os.ReadFile(filepath.Join(basePath, "carrier_down_count"))
	if err != nil {
		return carrier, err
	}

	carrier.Status = string(status)
	carrier.Changes = string(carrier_changes)
	carrier.UpCount = string(carrier_up_count)
	carrier.DownCount = string(carrier_down_count)

	return carrier, nil
}

func LinuxCollectInterfaceDevice(dir string) (InterfaceDevice, error) {
	intDevice := InterfaceDevice{}
	// Collect Id
	id, err := ReadSysId(dir)
	if err != nil {
		return intDevice, err
	}

	// collect device_id
	device_id, err := ReadSysDeviceId(dir)
	if err != nil {
		return intDevice, err
	}

	// collect class_id
	class_id, err := ReadSysClassId(dir)
	if err != nil {
		return intDevice, err
	}

	// collect driver_override
	driver_override, err := ReadSysDriverOverride(dir)
	if err != nil {
		return intDevice, err
	}

	// collect mod_alias
	modalias, err := ReadSysModalias(dir)
	if err != nil {
		return intDevice, err
	}

	// collect numa_node
	numa_node, err := ReadSysNumaNode(dir)
	if err != nil {
		return intDevice, err
	}

	// collect state
	state, err := ReadSysState(dir)
	if err != nil {
		return intDevice, err
	}

	// collect vendor
	vendor, err := ReadSysVendor(dir)
	if err != nil {
		return intDevice, err
	}

	// power
	power, err := ReadSysPower(dir)
	if err != nil {
		return intDevice, err
	}

	subsystem, err := ReadSysSubsystem(dir)
	if err != nil {
		return intDevice, err
	}

	uevent, err := ReadSysUevent(dir)
	if err != nil {
		return intDevice, err
	}

	monitoring, err := ReadSysMonitor(dir)
	if err != nil {
		return intDevice, err
	}

	ring_buffer, err := ReadSysRingBuffer(dir)
	if err != nil {
		return intDevice, err
	}

	channels, err := ReadSysInterfaceChannel(dir)
	if err != nil {
		return intDevice, err
	}

	intDevice.ID = id
	intDevice.DeviceID = device_id 
	intDevice.ClassID = class_id 
	intDevice.DriverOverride = driver_override 
	intDevice.Modalias = modalias 
	intDevice.NUMANode = numa_node
	intDevice.State = state 
	intDevice.Vendor = vendor 
	intDevice.Power = power 
	intDevice.Subsystem = subsystem 
	intDevice.Uevent = uevent 
	intDevice.Monitor = monitoring 
	intDevice.RingBuffer = ring_buffer
	intDevice.Channels = channels 

	return intDevice, nil
}

func LinuxCollectIGMP6 (key string) (MulticastGroupTable6, error) {
	igmpTab := MulticastGroupTable6{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return igmpTab, err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		length := len(fields)

		if length != 6 {
			continue
		}

		entry := MulticastGroup6{
			InterfaceIndex: fields[0],
			InterfaceName: fields[1],
			Address: fields[2],
			Users: fields[3],
			Flags: fields[4],
			Timer: fields[5],
		}

		igmpTab.Groups = append(igmpTab.Groups, entry)
	}

	return igmpTab, nil
}

func LinuxCollectIGMP(key string) (MulticastGroupTable, error) {
	igmpTab := MulticastGroupTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return igmpTab, err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		length := len(fields)

		entry := MulticastGroup{}

		switch length {
		case 5:	
			entry.Index = fields[0]
			entry.Device = fields[1]
			entry.Count = fields[3]
			entry.Querier = fields[4]

			entry.Group = ""
			entry.Users = ""
			entry.Timer = ""
			entry.Reporter = ""
		case 4:
			entry.Index = ""
			entry.Device = ""
			entry.Count = ""
			entry.Querier = ""

			entry.Group = fields[0]
			entry.Users = fields[1]
			entry.Timer = fields[2]
			entry.Reporter = fields[3]
		}

		igmpTab.Groups = append(igmpTab.Groups, entry)

	}

	return igmpTab, nil

}

func LinuxCollectPtype(key string) (PacketTypeHandlerTable, error) {
	ptypeTab := PacketTypeHandlerTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return ptypeTab, err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and header
		if line == "" || strings.HasPrefix(line, "Type") {
			continue
		}

		fields := strings.Fields(line)

		var handler PacketTypeHandler

		switch len(fields) {

		// Example:
		// 0800 ip_rcv
		case 2:
			handler.Type = fields[0]
			handler.Function = fields[1]

		// Example:
		// 0004 llc_rcv [llc]
		// or
		// 0800 eth0 ip_rcv
		case 3:
			handler.Type = fields[0]

			if strings.HasPrefix(fields[2], "[") {
				handler.Function = fields[1]
				handler.Module = strings.Trim(fields[2], "[]")
			} else {
				handler.Device = fields[1]
				handler.Function = fields[2]
			}

		// Example:
		// 0800 eth0 ip_rcv [llc]
		case 4:
			handler.Type = fields[0]
			handler.Device = fields[1]
			handler.Function = fields[2]
			handler.Module = strings.Trim(fields[3], "[]")

		default:
			continue
		}

		ptypeTab.Handlers = append(ptypeTab.Handlers, handler)

	}

	return ptypeTab, nil
}


func LinuxCollectIfInet6(key string) (IPv6AddressTable, error) {
	ipv6AddrTab := IPv6AddressTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return ipv6AddrTab, err
	}

	// raw print
	// fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) != 6 {
			continue
		}

		entry := IPv6Address{
			Address: fields[0],
			InterfaceIndex: fields[1],
			PrefixLength: fields[2],
			Scope: fields[3],
			Flags: fields[4],
			InterfaceName: fields[5],
		}

		ipv6AddrTab.Addresses = append(ipv6AddrTab.Addresses, entry)
	}

	return ipv6AddrTab, nil
	}


func LinuxCollectDev(key string) (NetworkDeviceStats, error) {
	networkDeviceStats := NetworkDeviceStats{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return networkDeviceStats, err
	}

	// raw print
	// fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {

		// Skip the two header lines and blank lines
		if i < 2 || strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 17 {
			continue
		}

		entry := InterfaceTraffic{
			Interface: strings.TrimSuffix(fields[0], ":"),

			Received: ReceivedTraffic{
				TrafficCounters: TrafficCounters{
					Bytes:      fields[1],
					Packets:    fields[2],
					Errors:     fields[3],
					Drop:       fields[4],
					Fifo:       fields[5],
					Compressed: fields[7],
				},
				Frame:     fields[6],
				Multicast: fields[8],
			},

			Transmitted: TransmittedTraffic{
				TrafficCounters: TrafficCounters{
					Bytes:      fields[9],
					Packets:    fields[10],
					Errors:     fields[11],
					Drop:       fields[12],
					Fifo:       fields[13],
					Compressed: fields[16],
				},
				Collisions: fields[14],
				Carrier:    fields[15],
			},
		}

		networkDeviceStats.Interfaces = append(networkDeviceStats.Interfaces, entry)
	}

	return networkDeviceStats, nil

}

func LinuxCollectArp(key string) (ARPTable, error) {
	arpTable := ARPTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return arpTable, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}
		
		fields := strings.Fields(line)

		if len(fields) < 6 {
			continue
		}

		entry := ARPEntry{
			IPAddress:    fields[0],
			HardwareType: fields[1],
			Flags:        fields[2],
			HardwareAddr: fields[3],
			Mask:         fields[4],
			Device:       fields[5],
		}

		arpTable.Entries = append(arpTable.Entries, entry)
	}

	return arpTable, nil
}