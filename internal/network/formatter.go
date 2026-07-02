package net

import (
	"fmt"
	"strings"
)

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