package net

import (
	// "fmt"
	"os"
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

func LinuxCollectPtype(key string) (PacketTypeHandlerTable, error) {
	
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