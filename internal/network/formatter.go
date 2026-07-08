package net

import (
	"fmt"
	"strings"
)
func FormatDev(stats NetworkDeviceStats) {
	if len(stats.Interfaces) == 0 {
		fmt.Println("No network interface statistics found.")
		return
	}

	fmt.Printf(
		"%-10s %-12s %-10s %-6s %-6s %-6s %-6s %-12s %-12s %-12s %-10s %-6s %-6s %-6s %-6s %-10s %-10s\n",
		"INTERFACE",
		"RX_BYTES",
		"RX_PKTS",
		"ERRS",
		"DROP",
		"FIFO",
		"FRAME",
		"RX_COMP",
		"RX_MULTI",
		"TX_BYTES",
		"TX_PKTS",
		"ERRS",
		"DROP",
		"FIFO",
		"COLLS",
		"CARRIER",
		"TX_COMP",
	)

	fmt.Println(strings.Repeat("-", 170))

	for _, iface := range stats.Interfaces {
		fmt.Printf(
			"%-10s %-12s %-10s %-6s %-6s %-6s %-6s %-12s %-12s %-12s %-10s %-6s %-6s %-6s %-6s %-10s %-10s\n",
			iface.Interface,

			iface.Received.Bytes,
			iface.Received.Packets,
			iface.Received.Errors,
			iface.Received.Drop,
			iface.Received.Fifo,
			iface.Received.Frame,
			iface.Received.Compressed,
			iface.Received.Multicast,

			iface.Transmitted.Bytes,
			iface.Transmitted.Packets,
			iface.Transmitted.Errors,
			iface.Transmitted.Drop,
			iface.Transmitted.Fifo,
			iface.Transmitted.Collisions,
			iface.Transmitted.Carrier,
			iface.Transmitted.Compressed,
		)
	}
}

func FormatIfNet6(table IPv6AddressTable) {
	if len(table.Addresses) == 0 {
		fmt.Println("No IPv6 entries found.")
		return
	}

	fmt.Printf(
		"%-35s %-20s %-15s %-20s %-20s %-10s\n",
		"IP ADDRESS",
		"Interface Index",
		"Prefix Length",
		"Scope",
		"Flags",
		"Interface Name",
	)

	fmt.Println(strings.Repeat("-", 135))

	for _, entry := range table.Addresses {
		fmt.Printf(
			"%-35s %-20s %-15s %-20s %-20s %-10s\n",
			entry.Address,
			entry.InterfaceIndex,
			entry.PrefixLength,
			entry.Scope,
			entry.Flags,
			entry.InterfaceName,
		)
	}


}

func FormatARP(table ARPTable) {
	if len(table.Entries) == 0 {
		fmt.Println("No ARP entries found.")
		return
	}

	fmt.Printf(
		"%-18s %-8s %-8s %-20s %-18s %-10s\n",
		"IP ADDRESS",
		"HW TYPE",
		"FLAGS",
		"HW ADDRESS",
		"MASK",
		"DEVICE",
	)

	fmt.Println(strings.Repeat("-", 90))

	for _, entry := range table.Entries {
		fmt.Printf(
			"%-18s %-8s %-8s %-20s %-18s %-10s\n",
			entry.IPAddress,
			entry.HardwareType,
			entry.Flags,
			entry.HardwareAddr,
			entry.Mask,
			entry.Device,
		)
	}
}