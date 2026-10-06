package net

import (
	"fmt"
	"strings"
	"path/filepath"
	"strconv"
)


func FormatXFRMStatistics(xfrm XFRMStatistics) {
	var b strings.Builder

	b.WriteString("┌─────────────────────────────────────────────────────┐\n")
	b.WriteString("│                    XFRM Statistics                  │\n")
	b.WriteString("├─────────────────────────────────────────────────────┤\n")

	// Inbound statistics
	b.WriteString("│ Inbound                                             │\n")
	b.WriteString("│ ─────────────────────────────────────────────────── │\n")

	fmt.Fprintf(&b, "│ Error              : %-30s │\n", xfrm.Inbound.Error)
	fmt.Fprintf(&b, "│ Buffer Error       : %-30s │\n", xfrm.Inbound.BufferError)
	fmt.Fprintf(&b, "│ Header Error       : %-30s │\n", xfrm.Inbound.HeaderError)
	fmt.Fprintf(&b, "│ No States          : %-30s │\n", xfrm.Inbound.NoStates)
	fmt.Fprintf(&b, "│ State Proto Error  : %-30s │\n", xfrm.Inbound.StateProtoError)
	fmt.Fprintf(&b, "│ State Mode Error   : %-30s │\n", xfrm.Inbound.StateModeError)
	fmt.Fprintf(&b, "│ State Seq Error    : %-30s │\n", xfrm.Inbound.StateSeqError)
	fmt.Fprintf(&b, "│ State Expired      : %-30s │\n", xfrm.Inbound.StateExpired)
	fmt.Fprintf(&b, "│ State Mismatch     : %-30s │\n", xfrm.Inbound.StateMismatch)
	fmt.Fprintf(&b, "│ State Invalid      : %-30s │\n", xfrm.Inbound.StateInvalid)
	fmt.Fprintf(&b, "│ State Dir Error    : %-30s │\n", xfrm.Inbound.StateDirError)
	fmt.Fprintf(&b, "│ Template Mismatch  : %-30s │\n", xfrm.Inbound.TemplateMismatch)
	fmt.Fprintf(&b, "│ No Policies        : %-30s │\n", xfrm.Inbound.NoPolicies)
	fmt.Fprintf(&b, "│ Policy Block       : %-30s │\n", xfrm.Inbound.PolicyBlock)
	fmt.Fprintf(&b, "│ Policy Error       : %-30s │\n", xfrm.Inbound.PolicyError)
	fmt.Fprintf(&b, "│ IPTFS Error        : %-30s │\n", xfrm.Inbound.IptfsError)

	b.WriteString("│                                                     │\n")

	// Outbound statistics
	b.WriteString("│ Outbound                                            │\n")
	b.WriteString("│ ─────────────────────────────────────────────────── │\n")

	fmt.Fprintf(&b, "│ Error              : %-30s │\n", xfrm.Outbound.Error)
	fmt.Fprintf(&b, "│ Bundle Gen Error   : %-30s │\n", xfrm.Outbound.BundleGenError)
	fmt.Fprintf(&b, "│ Bundle Check Error : %-30s │\n", xfrm.Outbound.BundleCheckError)
	fmt.Fprintf(&b, "│ No States          : %-30s │\n", xfrm.Outbound.NoStates)
	fmt.Fprintf(&b, "│ State Proto Error  : %-30s │\n", xfrm.Outbound.StateProtoError)
	fmt.Fprintf(&b, "│ State Mode Error   : %-30s │\n", xfrm.Outbound.StateModeError)
	fmt.Fprintf(&b, "│ State Seq Error    : %-30s │\n", xfrm.Outbound.StateSeqError)
	fmt.Fprintf(&b, "│ State Expired      : %-30s │\n", xfrm.Outbound.StateExpired)
	fmt.Fprintf(&b, "│ State Invalid      : %-30s │\n", xfrm.Outbound.StateInvalid)
	fmt.Fprintf(&b, "│ State Dir Error    : %-30s │\n", xfrm.Outbound.StateDirError)
	fmt.Fprintf(&b, "│ Policy Block       : %-30s │\n", xfrm.Outbound.PolicyBlock)
	fmt.Fprintf(&b, "│ Policy Dead        : %-30s │\n", xfrm.Outbound.PolicyDead)
	fmt.Fprintf(&b, "│ Policy Error       : %-30s │\n", xfrm.Outbound.PolicyError)
	fmt.Fprintf(&b, "│ No Queue Space     : %-30s │\n", xfrm.Outbound.NoQueueSpace)

	b.WriteString("│                                                     │\n")

	// Other XFRM statistics
	b.WriteString("│ Other                                               │\n")
	b.WriteString("│ ─────────────────────────────────────────────────── │\n")

	fmt.Fprintf(&b, "│ Forward Header Error: %-29s │\n", xfrm.ForwardHeaderError)
	fmt.Fprintf(&b, "│ Acquire Error       : %-29s │\n", xfrm.AcquireError)

	b.WriteString("└─────────────────────────────────────────────────────┘\n")

	fmt.Print(b.String())
}


func FormatNetworkInterfaces(interfaces NetworkInterfaces) {
	var b strings.Builder

	b.WriteString("┌─────────────────────────────────────────────────────┐\n")
	b.WriteString("│                  Network Interfaces                 │\n")
	b.WriteString("├─────────────────────────────────────────────────────┤\n")

	for i, iface := range interfaces.Interfaces {
		fmt.Fprintf(&b, "│ Interface %-42d │\n", i)
		b.WriteString("├─────────────────────────────────────────────────────┤\n")

		fmt.Fprintf(&b, "│ Name               : %-29s │\n", iface.Name)
		fmt.Fprintf(&b, "│ Address Assign Type: %-29s │\n", iface.AddressAssignType)
		fmt.Fprintf(&b, "│ Address Length     : %-29s │\n", iface.AddressLength)
		fmt.Fprintf(&b, "│ Address            : %-29s │\n", iface.Address)
		fmt.Fprintf(&b, "│ Broadcast          : %-29s │\n", iface.Broadcast)
		fmt.Fprintf(&b, "│ Dev ID             : %-29s │\n", iface.DevID)
		fmt.Fprintf(&b, "│ Dev Port           : %-29s │\n", iface.DevPort)

		// Carrier information
		b.WriteString("│                                                     │\n")
		b.WriteString("│ Carrier                                             │\n")
		b.WriteString("│ ─────────────────────────────────────────────────── │\n")

		fmt.Fprintf(&b, "│ Status             : %-29s │\n", iface.Carrier.Status)
		fmt.Fprintf(&b, "│ Changes            : %-29s │\n", iface.Carrier.Changes)
		fmt.Fprintf(&b, "│ Down Count         : %-29s │\n", iface.Carrier.DownCount)
		fmt.Fprintf(&b, "│ Up Count           : %-29s │\n", iface.Carrier.UpCount)

		// Device information
		b.WriteString("│                                                     │\n")
		b.WriteString("│ Device                                              │\n")
		b.WriteString("│ ─────────────────────────────────────────────────── │\n")

		fmt.Fprintf(&b, "│ ID                 : %-29s │\n", iface.Device.ID)
		fmt.Fprintf(&b, "│ Device ID          : %-29s │\n", iface.Device.DeviceID)
		fmt.Fprintf(&b, "│ Class ID           : %-29s │\n", iface.Device.ClassID)
		fmt.Fprintf(&b, "│ Driver Override    : %-29s │\n", iface.Device.DriverOverride)
		fmt.Fprintf(&b, "│ Modalias           : %-29s │\n", iface.Device.Modalias)
		fmt.Fprintf(&b, "│ NUMA Node          : %-29s │\n", iface.Device.NUMANode)
		fmt.Fprintf(&b, "│ State              : %-29s │\n", iface.Device.State)
		fmt.Fprintf(&b, "│ Vendor             : %-29s │\n", iface.Device.Vendor)

		// Power
		b.WriteString("│                                                     │\n")
		b.WriteString("│ Device Power                                        │\n")
		b.WriteString("│ ─────────────────────────────────────────────────── │\n")

		fmt.Fprintf(&b, "│ Control            : %-29s │\n", iface.Device.Power.Control)
		fmt.Fprintf(&b, "│ Active Time        : %-29s │\n", iface.Device.Power.RuntimeActiveTime)
		fmt.Fprintf(&b, "│ Status             : %-29s │\n", iface.Device.Power.RuntimeStatus)
		fmt.Fprintf(&b, "│ Suspended Time     : %-29s │\n", iface.Device.Power.RuntimeSuspendedTime)

		// Subsystem
		b.WriteString("│                                                     │\n")
		b.WriteString("│ Device Subsystem                                    │\n")
		b.WriteString("│ ─────────────────────────────────────────────────── │\n")

		fmt.Fprintf(&b, "│ Drivers Autoprobe  : %-29s │\n", iface.Device.Subsystem.DriversAutoprobe)
		fmt.Fprintf(&b, "│ Hibernation        : %-29s │\n", iface.Device.Subsystem.Hibernation)

		// Uevent
		b.WriteString("│                                                     │\n")
		b.WriteString("│ Device Uevent                                       │\n")
		b.WriteString("│ ─────────────────────────────────────────────────── │\n")

		fmt.Fprintf(&b, "│ Driver             : %-29s │\n", iface.Device.Uevent.Driver)
		fmt.Fprintf(&b, "│ Modalias           : %-29s │\n", iface.Device.Uevent.Modalias)

		// Monitor
		b.WriteString("│                                                     │\n")
		b.WriteString("│ Device Monitor                                      │\n")
		b.WriteString("│ ─────────────────────────────────────────────────── │\n")

		fmt.Fprintf(&b, "│ Client Conn. ID    : %-29s │\n", iface.Device.Monitor.ClientConnectionID)
		fmt.Fprintf(&b, "│ Client Latency     : %-29s │\n", iface.Device.Monitor.ClientLatency)
		fmt.Fprintf(&b, "│ Client Pending     : %-29s │\n", iface.Device.Monitor.ClientPending)
		fmt.Fprintf(&b, "│ Server Conn. ID    : %-29s │\n", iface.Device.Monitor.ServerConnectionID)
		fmt.Fprintf(&b, "│ Server Latency     : %-29s │\n", iface.Device.Monitor.ServerLatency)
		fmt.Fprintf(&b, "│ Server Pending     : %-29s │\n", iface.Device.Monitor.ServerPending)

		// Ring buffer
		b.WriteString("│                                                     │\n")
		b.WriteString("│ Device Ring Buffer                                   │\n")
		b.WriteString("│ ─────────────────────────────────────────────────── │\n")

		fmt.Fprintf(&b, "│ In Interrupt Mask  : %-29s │\n", iface.Device.RingBuffer.InInterruptMask)
		fmt.Fprintf(&b, "│ In Read Bytes Avail: %-29s │\n", iface.Device.RingBuffer.InReadBytesAvail)
		fmt.Fprintf(&b, "│ In Read Index      : %-29s │\n", iface.Device.RingBuffer.InReadIndex)
		fmt.Fprintf(&b, "│ In Write Bytes Avail: %-28s │\n", iface.Device.RingBuffer.InWriteBytesAvail)
		fmt.Fprintf(&b, "│ In Write Index     : %-29s │\n", iface.Device.RingBuffer.InWriteIndex)
		fmt.Fprintf(&b, "│ Out Interrupt Mask : %-29s │\n", iface.Device.RingBuffer.OutInterruptMask)
		fmt.Fprintf(&b, "│ Out Read Bytes Avail: %-28s │\n", iface.Device.RingBuffer.OutReadBytesAvail)
		fmt.Fprintf(&b, "│ Out Read Index     : %-29s │\n", iface.Device.RingBuffer.OutReadIndex)
		fmt.Fprintf(&b, "│ Out Write Bytes Avail: %-27s │\n", iface.Device.RingBuffer.OutWriteBytesAvail)
		fmt.Fprintf(&b, "│ Out Write Index    : %-29s │\n", iface.Device.RingBuffer.OutWriteIndex)

		// Channels
		if len(iface.Device.Channels) > 0 {
			b.WriteString("│                                                     │\n")
			b.WriteString("│ Interface Channels                                  │\n")
			b.WriteString("│ ─────────────────────────────────────────────────── │\n")

			for name, channel := range iface.Device.Channels {
				fmt.Fprintf(&b, "│ Channel: %-39s │\n", name)

				fmt.Fprintf(&b, "│   CPU              : %-27s │\n", channel.CPU)
				fmt.Fprintf(&b, "│   Events           : %-27s │\n", channel.Events)
				fmt.Fprintf(&b, "│   In Mask          : %-27s │\n", channel.InMask)
				fmt.Fprintf(&b, "│   Interrupts       : %-27s │\n", channel.Interrupts)
				fmt.Fprintf(&b, "│   Intr In Full     : %-27s │\n", channel.IntrInFull)
				fmt.Fprintf(&b, "│   Intr Out Empty   : %-27s │\n", channel.IntrOutEmpty)
				fmt.Fprintf(&b, "│   Latency          : %-27s │\n", channel.Latency)
				fmt.Fprintf(&b, "│   Monitor ID       : %-27s │\n", channel.MonitorID)
				fmt.Fprintf(&b, "│   Out Full First   : %-27s │\n", channel.OutFullFirst)
				fmt.Fprintf(&b, "│   Out Full Total   : %-27s │\n", channel.OutFullTotal)
				fmt.Fprintf(&b, "│   Out Mask         : %-27s │\n", channel.OutMask)
				fmt.Fprintf(&b, "│   Pending          : %-27s │\n", channel.Pending)
				fmt.Fprintf(&b, "│   Read Avail       : %-27s │\n", channel.ReadAvail)
				fmt.Fprintf(&b, "│   Subchannel ID    : %-27s │\n", channel.SubchannelID)
				fmt.Fprintf(&b, "│   Write Avail      : %-27s │\n", channel.WriteAvail)
			}
		}

		b.WriteString("├─────────────────────────────────────────────────────┤\n")
	}

	b.WriteString("└─────────────────────────────────────────────────────┘\n")

	fmt.Print(b.String())
}

func FormatOSRelease(info OSRelease) {
	var b strings.Builder

	b.WriteString("┌─────────────────────────────────────────────────────┐\n")
	b.WriteString("│                       os-release                    │\n")
	b.WriteString("├─────────────────────────────────────────────────────┤\n")

	fmt.Fprintf(&b, "│ Pretty Name        : %-30s │\n", info.PrettyName)
	fmt.Fprintf(&b, "│ Name               : %-30s │\n", info.Name)
	fmt.Fprintf(&b, "│ Version ID         : %-30s │\n", info.VersionID)
	fmt.Fprintf(&b, "│ Version            : %-30s │\n", info.Version)
	fmt.Fprintf(&b, "│ Version Codename   : %-30s │\n", info.VersionCodename)
	fmt.Fprintf(&b, "│ ID                 : %-30s │\n", info.ID)
	fmt.Fprintf(&b, "│ ID Like            : %-30s │\n", info.IDLike)
	fmt.Fprintf(&b, "│ Home URL           : %-30s │\n", info.HomeURL)
	fmt.Fprintf(&b, "│ Support URL        : %-30s │\n", info.SupportURL)
	fmt.Fprintf(&b, "│ Bug Report URL     : %-30s │\n", info.BugReportURL)
	fmt.Fprintf(&b, "│ Privacy Policy URL : %-30s │\n", info.PrivacyPolicyURL)
	fmt.Fprintf(&b, "│ Ubuntu Codename    : %-30s │\n", info.UbuntuCodename)
	fmt.Fprintf(&b, "│ Logo               : %-30s │\n", info.Logo)

	b.WriteString("└─────────────────────────────────────────────────────┘\n")

	fmt.Print(b.String())
}


func FormatHostName(name Hostname) {
	var b strings.Builder

	b.WriteString(name.Name)
	fmt.Print(b.String())
}


func FormatHostsFile(hosts HostsFile) {
	var b strings.Builder

	b.WriteString("┌─────────────────────────────────────────────────────────────────┐\n")
	b.WriteString("│                            HostsFile                            │\n")
	b.WriteString("├─────────────────────────────────────────────────────────────────┤\n")

	for i, host := range hosts.Entries {
		b.WriteString("│ Entry ")
		b.WriteString(strconv.Itoa(i))
		b.WriteString("                                                         │\n")
		b.WriteString("├─────────────────────────────────────────────────────────────────┤\n")

		b.WriteString("│ Database        : ")
		b.WriteString(host.IPAddress)
		b.WriteString("\n")

		b.WriteString("│ Hostnames       : ")
		b.WriteString(strings.Join(host.Hostnames, ", "))
		b.WriteString("\n")
	}

	fmt.Print(b.String())
}


func FormatResolvConf(res ResolvConf) {
	var b strings.Builder

	b.WriteString("┌───────────────────────────────────────────────┐\n")
	b.WriteString("│                  ResolvConf                   │\n")
	b.WriteString("├───────────────────────────────────────────────┤\n")

	b.WriteString(res.Content)

	fmt.Print(b.String())

}

func FormatNsswitchConfiguration(nsw NsswitchConfiguration) {
	var b strings.Builder

	b.WriteString("┌───────────────────────────────────────────────┐\n")
	b.WriteString("│                 nsswitch.conf                 │\n")
	b.WriteString("├───────────────────────────────────────────────┤\n")

	for i, item := range nsw.Entries {
		b.WriteString("│ Entry ")
		b.WriteString(strconv.Itoa(i))
		b.WriteString("                                      │\n")
		b.WriteString("├───────────────────────────────────────────────┤\n")

		b.WriteString("│ Database       : ")
		b.WriteString(item.Database)
		b.WriteString("\n")

		b.WriteString("│ Sources        : ")
		b.WriteString(strings.Join(item.Sources, ", "))
		b.WriteString("\n")

		b.WriteString("├───────────────────────────────────────────────┤\n")
	}

	b.WriteString("└───────────────────────────────────────────────┘\n")

	fmt.Print(b.String())
}

func FormatPasswdFile(pswd PasswdFile) {
	var b strings.Builder

	b.WriteString("┌───────────────────────────────────────────────┐\n")
	b.WriteString("│                   PasswdFile                  │\n")
	b.WriteString("├───────────────────────────────────────────────┤\n")

	for i, p := range pswd.Entries {
		b.WriteString("│ Entry ")
		b.WriteString(strconv.Itoa(i))
		b.WriteString("                                      │\n")
		b.WriteString("├───────────────────────────────────────────────┤\n")

		b.WriteString("│ Username       : ")
		b.WriteString(p.Username)
		b.WriteString("\n")
		b.WriteString("│ Password       : ")
		b.WriteString(p.Password)
		b.WriteString("\n")
		b.WriteString("│ UID            : ")
		b.WriteString(p.UID)
		b.WriteString("\n")
		b.WriteString("│ GID            : ")
		b.WriteString(p.GID)
		b.WriteString("\n")
		b.WriteString("│ Comment        : ")
		b.WriteString(p.Comment)
		b.WriteString("\n")
		b.WriteString("│ Home Directory : ")
		b.WriteString(p.HomeDirectory)
		b.WriteString("\n")
		b.WriteString("│ Shell          : ")
		b.WriteString(p.Shell)
		b.WriteByte('\n')

		b.WriteString("└───────────────────────────────────────────────┘\n")
	}

	fmt.Print(b.String())
}

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