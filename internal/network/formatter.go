package net

import (
	"fmt"
	"strings"
	"path/filepath"
)

func FormatShellsSummary(sf ShellsFile) string {
	var b strings.Builder

	b.WriteString("┌─────────────────────────────────────────────┐\n")
	b.WriteString("│               SHELLS SUMMARY                │\n")
	b.WriteString("├─────────────────────────────────────────────┤\n")
	b.WriteString(fmt.Sprintf("  Installed Shells : %d\n", len(sf.Paths)))
	b.WriteString("└─────────────────────────────────────────────┘")

	return b.String()
}

func FormatShellsDetailed(sf ShellsFile) string {
	var b strings.Builder

	b.WriteString(strings.TrimSpace(FormatShellsSummary(sf)))
	b.WriteString("\n\n")

	b.WriteString(fmt.Sprintf("%-20s %s\n", "SHELL", "PATH"))
	b.WriteString(strings.Repeat("─", 70))
	b.WriteString("\n")

	for _, path := range sf.Paths {
		b.WriteString(fmt.Sprintf("%-20s %s\n",
			filepath.Base(path),
			path,
		))
	}

	return strings.TrimSpace(b.String())
}

func FormatProtocolRegistrySummary(pr ProtocolRegistry) string {
	var aliasCount int

	for _, p := range pr.Protocols {
		aliasCount += len(p.Aliases)
	}

	var b strings.Builder

	b.WriteString("┌─────────────────────────────────────────────┐\n")
	b.WriteString("│             PROTOCOLS SUMMARY               │\n")
	b.WriteString("├─────────────────────────────────────────────┤\n")
	b.WriteString(fmt.Sprintf("  Total Protocols : %d\n", len(pr.Protocols)))
	b.WriteString(fmt.Sprintf("  Total Aliases   : %d\n", aliasCount))
	b.WriteString("└─────────────────────────────────────────────┘")

	return b.String()
}

func FormatProtocolRegistryDetailed(pr ProtocolRegistry) string {
	var b strings.Builder

	b.WriteString(strings.TrimSpace(FormatProtocolRegistrySummary(pr)))
	b.WriteString("\n\n")

	b.WriteString(fmt.Sprintf("%-24s %-10s %s\n",
		"NAME", "NUMBER", "ALIASES"))
	b.WriteString(strings.Repeat("─", 70))
	b.WriteString("\n")

	for _, p := range pr.Protocols {
		aliases := "-"
		if len(p.Aliases) > 0 {
			aliases = strings.Join(p.Aliases, ", ")
		}

		b.WriteString(fmt.Sprintf("%-24s %-10s %s\n",
			p.Name,
			p.Number,
			aliases,
		))
	}

	return strings.TrimSpace(b.String())
}

func FormatServicesSummary(sf ServicesFile) string {
	var tcpCount, udpCount int

	for _, s := range sf.Entries {
		switch strings.ToLower(s.Protocol) {
		case "tcp":
			tcpCount++
		case "udp":
			udpCount++
		}
	}

	var b strings.Builder

	b.WriteString("┌─────────────────────────────────────────────┐\n")
	b.WriteString("│              SERVICES SUMMARY               │\n")
	b.WriteString("├─────────────────────────────────────────────┤\n")
	b.WriteString(fmt.Sprintf("  Total Services : %d\n", len(sf.Entries)))
	b.WriteString(fmt.Sprintf("  TCP Services   : %d\n", tcpCount))
	b.WriteString(fmt.Sprintf("  UDP Services   : %d\n", udpCount))
	b.WriteString("└─────────────────────────────────────────────┘")

	return b.String()
}

func FormatServicesDetailed(sf ServicesFile) string {
	var b strings.Builder

	b.WriteString(strings.TrimSpace(FormatServicesSummary(sf)))
	b.WriteString("\n\n")

	b.WriteString(fmt.Sprintf("%-24s %-8s %-8s %s\n",
		"NAME", "PORT", "PROTO", "ALIASES"))
	b.WriteString(strings.Repeat("─", 80))
	b.WriteString("\n")

	for _, s := range sf.Entries {
		aliases := "-"
		if len(s.Aliases) > 0 {
			aliases = strings.Join(s.Aliases, ", ")
		}

		b.WriteString(fmt.Sprintf("%-24s %-8s %-8s %s\n",
			s.Name,
			s.Port,
			strings.ToUpper(s.Protocol),
			aliases,
		))
	}

	return strings.TrimSpace(b.String())
}

func FormatPType(handlers []PacketTypeHandler) {
	if len(handlers) == 0 {
		fmt.Println("No packet type handlers found.")
		return
	}

	fmt.Printf(
		"%-10s %-20s %-30s %-20s\n",
		"Type",
		"Device",
		"Function",
		"Module",
	)

	for _, handler := range handlers {
		device := handler.Device
		if device == "" {
			device = "*"
		}

		module := handler.Module
		if module == "" {
			module = "-"
		}

		fmt.Printf(
			"%-10s %-20s %-30s %-20s\n",
			handler.Type,
			device,
			handler.Function,
			module,
		)
	}
}

func FormatIGMP6(table MulticastGroupTable6) {
	if len(table.Groups) == 0 {
		fmt.Println("No IGMP entries found.")
		return
	}

	fmt.Printf(
		"%-18s %-20s %-45s %-20s %-20s %-10s\n",
		"InterfaceIndex", 
		"InterfaceName",  
		"Address",        
		"Users",          
		"Flags",          
		"Timer",          
	)

	for _, irow := range table.Groups {
		
		fmt.Printf(
			"%-18s %-20s %-45s %-20s %-20s %-10s\n",
			irow.InterfaceIndex,   
			irow.InterfaceName,  
			irow.Address,
			irow.Users,
			irow.Flags,    
			irow.Timer,    
		)
		
	}
}


func FormatIGMP(table MulticastGroupTable) {
	if len(table.Groups) == 0 {
		fmt.Println("No IGMP entries found.")
		return
	}

	fmt.Printf(
		"%-35s %-20s %-15s %-20s %-20s %-10s %-15s %-15s\n",
		"Index",    
		"Device",   
		"Count",    
		"Querier",  
		"Group",    
		"Users",    
		"Timer",    
		"Reporter",
	)

	for _, irow := range table.Groups {
		
		fmt.Printf(
			"%-35s %-20s %-15s %-20s %-20s %-10s %-15s %-15s\n",
			irow.Index ,   
			irow.Device ,  
			irow.Count,
			irow.Querier,
			irow.Group,    
			irow.Users,    
			irow.Timer,    
			irow.Reporter,
		)
		
	}




}

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