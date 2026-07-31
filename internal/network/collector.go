package net

import (
	// "fmt"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"reflect"
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

func LinuxCollectTLSStatistics(key string) (TLSStatistics, error) {
	tls := TLSStatistics{}

	__curr__ := TLSSessionCounts{}
	__total__ := TLSSessionCounts{}
	__errors__ := TLSErrorCounts{}
	__rekey__ := TLSRekeyCounts{}

	return tls, nil
}

func LinuxCollectProtocolStatisticsTable(key string) (ProtocolStatisticsTable, error) {
    pstab := ProtocolStatisticsTable{}

    capabilities := []string{
        "cl", "co", "di", "ac", "io", "in", "de", "sh", "ss",
        "gs", "se", "re", "bi", "br", "ha", "uh", "gp", "em",
    }

    content, err := os.ReadFile(DirMap[key])
    if err != nil {
        return pstab, err
    }

    lines := strings.Split(string(content), "\n")

    for idx, line := range lines {
        if idx == 0 {
            continue
        }

        fields := strings.Fields(line)
        if len(fields) < 8 {
            continue
        }

        caps := make(map[string]string)

        mapElems := fields[8:]

        for i, cap := range capabilities {
            if i >= len(mapElems) {
                break
            }
            caps[cap] = mapElems[i]
        }

        entry := ProtocolStatistics{
            Protocol:     fields[0],
            Size:         fields[1],
            Sockets:      fields[2],
            Memory:       fields[3],
            Press:        fields[4],
            MaxHeader:    fields[5],
            Slab:         fields[6],
            Module:       fields[7],
            Capabilities: caps,
        }

        pstab.Protocols = append(pstab.Protocols, entry)
    }

    return pstab, nil
}

func LinuxCollectSocketStatistics6(key string) (SocketStatistics6, error) {
	sock := SocketStatistics6{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return sock, err
	}

	lines := strings.Split(string(content), "\n")

	stc := SocketTypeCounters{}

	for _, line := range lines {
		__line__ := strings.TrimSpace(line)

		fields := strings.Fields(__line__)

		if len(fields) < 2 {
			continue
		}

		value := strings.Join(fields[1:], ",")

		if strings.HasPrefix(__line__, "TCP:") {
			stc.TCP = value			
		}

		if strings.HasPrefix(__line__, "UDP:") {
			stc.UDP = value			
		}

		if strings.HasPrefix(__line__, "UDPLITE:") {
			stc.UDPLite = value			
		}

		if strings.HasPrefix(__line__, "RAW:") {
			stc.RAW = value			
		}

		if strings.HasPrefix(__line__, "FRAG:") {
			stc.Frag = value			
		}
	}

	sock.Counters = stc

	return sock, nil
}

func LinuxCollectSocketStatistics(key string) (SocketStatistics, error) {
	sock := SocketStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return sock, err
	}

	lines := strings.Split(string(content), "\n")

	stc := SocketTypeCounters{}

	for _, line := range lines {
		__line__ := strings.TrimSpace(line)

		fields := strings.Fields(__line__)

		if len(fields) < 2 {
			continue
		}

		value := strings.Join(fields[1:], ",")

		if strings.HasPrefix(__line__, "sockets:") {
			sock.Sockets = value			
		}

		if strings.HasPrefix(__line__, "TCP:") {
			stc.TCP = value			
		}

		if strings.HasPrefix(__line__, "UDP:") {
			stc.UDP = value			
		}

		if strings.HasPrefix(__line__, "UDPLITE:") {
			stc.UDPLite = value			
		}

		if strings.HasPrefix(__line__, "RAW:") {
			stc.RAW = value			
		}

		if strings.HasPrefix(__line__, "FRAG:") {
			stc.Frag = value			
		}
	}

	sock.Counters = stc

	return sock, nil

}

func LinuxCollectSNMP6Statistics(key string) (SNMP6Statistics, error) {
	snmp6 := SNMP6Statistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return snmp6, err
	}

	var ans map[string]string

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) != 2 {
			continue
		}

		ans[fields[0]] = fields[1]
	}

	snmp6.Metrics = ans

	return snmp6, nil

}

func LinuxCollectSNMPStatistics(key string) (SNMPStatistics, error) {
	snmp := SNMPStatistics{}

	ip, err := LinuxCollectSNMPIPStatistics(key)
	if err != nil {
		return snmp, err
	}

	snmp.IP = ip 

	icmp, err := LinuxCollectSNMPICMPStatistics(key)
	if err != nil {
		return snmp, err
	}

	snmp.ICMP = icmp

	snmp.ICMPMsg = SNMPICMPMessageStatistics{}

	tcp, err := LinuxCollectSNMPTCPStatistics(key)
	if err != nil {
		return snmp, err
	}

	snmp.TCP = tcp

	udp, err := LinuxCollectSNMPUDPStatistics(key)
	if err != nil {
		return snmp, err
	}

	snmp.UDP = udp

	udp_lite, err := LinuxCollectSNMPUDP_Lite_Statistics(key)
	if err != nil {
		return snmp, err
	}

	snmp.UDPLite = udp_lite

	return snmp, nil

}

func LinuxCollectSNMPUDP_Lite_Statistics(key string) (UDPStatistics, error) {
	udp := UDPStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return udp, err
	}

	lines := strings.Split(string(content), "\n")

	var count int 

	for idx, line := range lines {
		if idx == 8 {
			fields := strings.Fields(line)
			count = len(fields)
			continue
		}

		if idx == 9 {
			fields := strings.Fields(line)

			if len(fields) != count {
				continue
			}
		}
	}

	return udp, nil
}

func LinuxCollectSNMPUDPStatistics(key string) (UDPStatistics, error) {
	udp := UDPStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return udp, err
	}

	lines := strings.Split(string(content), "\n")

	var count int 

	for idx, line := range lines {
		if idx == 6 {
			fields := strings.Fields(line)
			count = len(fields)
			continue
		}

		if idx == 7 {
			fields := strings.Fields(line)

			if len(fields) != count {
				continue
			}
		}
	}

	return udp, nil
}

func LinuxCollectSNMPTCPStatistics(key string) (SNMPTCPStatistics, error) {
	tcp := SNMPTCPStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return tcp, err
	}

	lines := strings.Split(string(content), "\n")

	var count int

	for idx, line := range lines {
		if idx == 4 {
			fields := strings.Fields(line)
			count = len(fields)
			continue
		}

		if idx == 5 {
			fields := strings.Fields(line)

			if len(fields) != count {
				continue
			}
		}
	}

	return tcp, nil
}

func LinuxCollectSNMPICMPStatistics(key string) (SNMPICMPStatistics, error) {
	snmpicmp := SNMPICMPStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return snmpicmp, err
	}

	lines := strings.Split(string(content), "\n")

	var count int

	for idx, line := range lines {
		if idx == 2 {
			fields := strings.Fields(line)
			count = len(fields)
			continue
		}

		if idx == 3 {
			fields := strings.Fields(line)

			if len(fields) != count {
				continue
			}
		}
	}

	return snmpicmp, nil
}

func LinuxCollectSNMPIPStatistics(key string) (SNMPIPStatistics, error) {
	snmp := SNMPIPStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return snmp, err
	}

	lines := strings.Split(string(content), "\n")

	var count int

	for idx, line := range lines {
		if idx == 0 {
			fields := strings.Fields(line)
			count = len(fields)
			continue
		}

		if idx == 1 {
			fields := strings.Fields(line)

			if len(fields) != count {
				continue
			}

			snmp = SNMPIPStatistics{
				Forwarding:      fields[1],
				DefaultTTL:      fields[2],
				InReceives:      fields[3],
				InHdrErrors:     fields[4],
				InAddrErrors:    fields[5],
				ForwDatagrams:   fields[6],
				InUnknownProtos: fields[7],
				InDiscards:      fields[8],
				InDelivers:      fields[9],
				OutRequests:     fields[10],
				OutDiscards:     fields[11],
				OutNoRoutes:     fields[12],
				ReasmTimeout:    fields[13],
				ReasmReqds:      fields[14],
				ReasmOKs:        fields[15],
				ReasmFails:      fields[16],
				FragOKs:         fields[17],
				FragFails:       fields[18],
				FragCreates:     fields[19],
				OutTransmits:    fields[20],
			}

			break
		}
	}

	return snmp, nil
}

func LinuxCollectNetworkExtendedStatistics(key string) (NetworkExtendedStatistics, error) {
	netstats := NetworkExtendedStatistics{}

	tcp, err := LinuxCollectTCPExtendedStatistics(key)
	if err != nil {
		return netstats, err
	}

	ip, err := LinuxCollectIPExtendedStatistics(key)
	if err != nil {
		return netstats, err
	}

	netstats.TCP = tcp 
	netstats.IP = ip

	return netstats, nil
}

func LinuxCollectIPExtendedStatistics(key string) (IPExtendedStatistics, error) {
	ies := IPExtendedStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return ies, err
	}

	lines := strings.Split(string(content), "\n")

	for i := 0; i < len(lines)-1; i++ {
		if strings.HasPrefix(lines[i], "IpExt:") {
			values := strings.Fields(lines[i+1])

			if len(values) < 19 { // "IpExt:" + 18 values
				return ies, fmt.Errorf("invalid IpExt statistics")
			}

			ies.InNoRoutes = values[1]
			ies.InTruncatedPkts = values[2]
			ies.InMcastPkts = values[3]
			ies.OutMcastPkts = values[4]
			ies.InBcastPkts = values[5]
			ies.OutBcastPkts = values[6]
			ies.InOctets = values[7]
			ies.OutOctets = values[8]
			ies.InMcastOctets = values[9]
			ies.OutMcastOctets = values[10]
			ies.InBcastOctets = values[11]
			ies.OutBcastOctets = values[12]
			ies.InCsumErrors = values[13]
			ies.InNoECTPkts = values[14]
			ies.InECT1Pkts = values[15]
			ies.InECT0Pkts = values[16]
			ies.InCEPkts = values[17]
			ies.ReasmOverlaps = values[18]

			break
		}
	}

	return ies, nil
}

// tcpExtFieldMap maps the column names as they appear in the "TcpExt:" header
// line of /proc/net/netstat to the corresponding TCPExtendedStatistics field.
// Using a name-based map (instead of positional indices) means the parser
// keeps working even if the kernel adds, removes, or reorders counters.
var tcpExtFieldMap = map[string]string{
	"SyncookiesSent":            "SyncookiesSent",
	"SyncookiesRecv":            "SyncookiesRecv",
	"SyncookiesFailed":          "SyncookiesFailed",
	"EmbryonicRsts":             "EmbryonicRsts",
	"PruneCalled":               "PruneCalled",
	"RcvPruned":                 "RcvPruned",
	"OfoPruned":                 "OfoPruned",
	"OutOfWindowIcmps":          "OutOfWindowICMPs",
	"LockDroppedIcmps":          "LockDroppedICMPs",
	"ArpFilter":                 "ArpFilter",
	"TW":                        "TimeWait",
	"TWRecycled":                "TimeWaitRecycled",
	"TWKilled":                  "TimeWaitKilled",
	"PAWSActive":                "PAWSActive",
	"PAWSEstab":                 "PAWSEstab",
	"BeyondWindow":              "BeyondWindow",
	"TSEcrRejected":             "TSEcrRejected",
	"PAWSOldAck":                "PAWSOldAck",
	"PAWSTimewait":              "PAWSTimewait",
	"DelayedACKs":               "DelayedACKs",
	"DelayedACKLocked":          "DelayedACKLocked",
	"DelayedACKLost":            "DelayedACKLost",
	"ListenOverflows":           "ListenOverflows",
	"ListenDrops":               "ListenDrops",
	"TCPHPHits":                 "HPHits",
	"TCPPureAcks":               "PureAcks",
	"TCPHPAcks":                 "HPAcks",
	"TCPRenoRecovery":           "RenoRecovery",
	"TCPSackRecovery":           "SackRecovery",
	"TCPSACKReneging":           "SACKReneging",
	"TCPSACKReorder":            "SACKReorder",
	"TCPRenoReorder":            "RenoReorder",
	"TCPTSReorder":              "TSReorder",
	"TCPFullUndo":               "FullUndo",
	"TCPPartialUndo":            "PartialUndo",
	"TCPDSACKUndo":              "DSACKUndo",
	"TCPLossUndo":               "LossUndo",
	"TCPLostRetransmit":         "LostRetransmit",
	"TCPRenoFailures":           "RenoFailures",
	"TCPSackFailures":           "SackFailures",
	"TCPLossFailures":           "LossFailures",
	"TCPFastRetrans":            "FastRetrans",
	"TCPSlowStartRetrans":       "SlowStartRetrans",
	"TCPTimeouts":               "Timeouts",
	"TCPLossProbes":             "LossProbes",
	"TCPLossProbeRecovery":      "LossProbeRecovery",
	"TCPRenoRecoveryFail":       "RenoRecoveryFail",
	"TCPSackRecoveryFail":       "SackRecoveryFail",
	"TCPRcvCollapsed":           "RcvCollapsed",
	"TCPBacklogCoalesce":        "BacklogCoalesce",
	"TCPDSACKOldSent":           "DSACKOldSent",
	"TCPDSACKOfoSent":           "DSACKOfoSent",
	"TCPDSACKRecv":              "DSACKRecv",
	"TCPDSACKOfoRecv":           "DSACKOfoRecv",
	"TCPAbortOnData":            "AbortOnData",
	"TCPAbortOnClose":           "AbortOnClose",
	"TCPAbortOnMemory":          "AbortOnMemory",
	"TCPAbortOnTimeout":         "AbortOnTimeout",
	"TCPAbortOnLinger":          "AbortOnLinger",
	"TCPAbortFailed":            "AbortFailed",
	"TCPMemoryPressures":        "MemoryPressures",
	"TCPMemoryPressuresChrono":  "MemoryPressuresChrono",
	"TCPSACKDiscard":            "SACKDiscard",
	"TCPDSACKIgnoredOld":        "DSACKIgnoredOld",
	"TCPDSACKIgnoredNoUndo":     "DSACKIgnoredNoUndo",
	"TCPSpuriousRTOs":           "SpuriousRTOs",
	"TCPMD5NotFound":            "MD5NotFound",
	"TCPMD5Unexpected":          "MD5Unexpected",
	"TCPMD5Failure":             "MD5Failure",
	"TCPSackShifted":            "SackShifted",
	"TCPSackMerged":             "SackMerged",
	"TCPSackShiftFallback":      "SackShiftFallback",
	"TCPBacklogDrop":            "BacklogDrop",
	"PFMemallocDrop":            "PFMemallocDrop",
	"TCPMinTTLDrop":             "MinTTLDrop",
	"TCPDeferAcceptDrop":        "DeferAcceptDrop",
	"IPReversePathFilter":       "IPReversePathFilter",
	"TCPTimeWaitOverflow":       "TimeWaitOverflow",
	"TCPReqQFullDoCookies":      "ReqQFullDoCookies",
	"TCPReqQFullDrop":           "ReqQFullDrop",
	"TCPRetransFail":            "RetransFail",
	"TCPRcvCoalesce":            "RcvCoalesce",
	"TCPOFOQueue":               "OFOQueue",
	"TCPOFODrop":                "OFODrop",
	"TCPOFOMerge":               "OFOMerge",
	"TCPChallengeACK":           "ChallengeACK",
	"TCPSYNChallenge":           "SYNChallenge",
	"TCPFastOpenActive":         "FastOpenActive",
	"TCPFastOpenActiveFail":     "FastOpenActiveFail",
	"TCPFastOpenPassive":        "FastOpenPassive",
	"TCPFastOpenPassiveFail":    "FastOpenPassiveFail",
	"TCPFastOpenListenOverflow": "FastOpenListenOverflow",
	"TCPFastOpenCookieReqd":     "FastOpenCookieReqd",
	"TCPFastOpenBlackhole":      "FastOpenBlackhole",
	"TCPSpuriousRtxHostQueues":  "SpuriousRtxHostQueues",
	"BusyPollRxPackets":         "BusyPollRxPackets",
	"TCPAutoCorking":            "AutoCorking",
	"TCPFromZeroWindowAdv":      "FromZeroWindowAdv",
	"TCPToZeroWindowAdv":        "ToZeroWindowAdv",
	"TCPWantZeroWindowAdv":      "WantZeroWindowAdv",
	"TCPSynRetrans":             "SynRetrans",
	"TCPOrigDataSent":           "OrigDataSent",
	"TCPHystartTrainDetect":     "HystartTrainDetect",
	"TCPHystartTrainCwnd":       "HystartTrainCwnd",
	"TCPHystartDelayDetect":     "HystartDelayDetect",
	"TCPHystartDelayCwnd":       "HystartDelayCwnd",
	"TCPACKSkippedSynRecv":      "ACKSkippedSynRecv",
	"TCPACKSkippedPAWS":         "ACKSkippedPAWS",
	"TCPACKSkippedSeq":          "ACKSkippedSeq",
	"TCPACKSkippedFinWait2":     "ACKSkippedFinWait2",
	"TCPACKSkippedTimeWait":     "ACKSkippedTimeWait",
	"TCPACKSkippedChallenge":    "ACKSkippedChallenge",
	"TCPWinProbe":               "WinProbe",
	"TCPKeepAlive":              "KeepAlive",
	"TCPMTUPFail":               "MTUPFail",
	"TCPMTUPSuccess":            "MTUPSuccess",
	"TCPDelivered":              "Delivered",
	"TCPDeliveredCE":            "DeliveredCE",
	"TCPAckCompressed":          "AckCompressed",
	"TCPZeroWindowDrop":         "ZeroWindowDrop",
	"TCPRcvQDrop":               "RcvQDrop",
	"TCPWqueueTooBig":           "WqueueTooBig",
	"TCPFastOpenPassiveAltKey":  "FastOpenPassiveAltKey",
	"TcpTimeoutRehash":          "TimeoutRehash",
	"TcpDuplicateDataRehash":    "DuplicateDataRehash",
	"TCPDSACKRecvSegs":          "DSACKRecvSegs",
	"TCPDSACKIgnoredDubious":    "DSACKIgnoredDubious",
	"TCPMigrateReqSuccess":      "MigrateReqSuccess",
	"TCPMigrateReqFailure":      "MigrateReqFailure",
	"TCPPLBRehash":              "PLBRehash",
	"TCPAORequired":             "AORequired",
	"TCPAOBad":                  "AOBad",
	"TCPAOKeyNotFound":          "AOKeyNotFound",
	"TCPAOGood":                 "AOGood",
	"TCPAODroppedIcmps":         "AODroppedICMPs",
}

func LinuxCollectTCPExtendedStatistics(key string) (TCPExtendedStatistics, error) {
	tes := TCPExtendedStatistics{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return tes, err
	}

	lines := strings.Split(string(content), "\n")

	for i := 0; i < len(lines)-1; i++ {
		if !strings.HasPrefix(lines[i], "TcpExt:") {
			continue
		}

		headers := strings.Fields(lines[i])[1:]  // drop the "TcpExt:" label
		values := strings.Fields(lines[i+1])[1:] // drop the "TcpExt:" label

		if len(headers) != len(values) {
			return tes, fmt.Errorf(
				"invalid TcpExt statistics: %d headers but %d values",
				len(headers), len(values),
			)
		}

		v := reflect.ValueOf(&tes).Elem()

		for idx, name := range headers {
			fieldName, known := tcpExtFieldMap[name]
			if !known {
				// A counter the kernel added that we don't track yet.
				// Skip it instead of failing the whole parse.
				continue
			}

			f := v.FieldByName(fieldName)
			if !f.IsValid() || !f.CanSet() {
				continue
			}

			f.SetString(values[idx])
		}

		return tes, nil
	}

	return tes, fmt.Errorf("TcpExt section not found in %s", DirMap[key])
}

func LinuxCollectProcessConnector(key string) (ProcessConnector, error) {
	pcn := ProcessConnector{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return pcn, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i==0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		pcn.Name = fields[0]
		pcn.ID = fields[1]

		if pcn.Name != "" && pcn.ID != "" {
			break
		}


	}

	return pcn, nil


}

func LinuxCollectNetlinkTable(key string) (NetlinkTable, error) {
	netl := NetlinkTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return netl, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i==0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 10 {
			continue
		}

		entry := NetlinkSocket{
			Socket: fields[0],
			Protocol: fields[1],
			PID: fields[2],
			Groups: fields[3],
			RecvMem: fields[4],
			SendMem: fields[5],
			Dump: fields[6],
			Locks: fields[7],
			Drops: fields[8],
			Inode: fields[9],
		}

		netl.Sockets = append(netl.Sockets, entry)
	}

	return netl, nil
}

func LinuxCollectPacketSocketTable(key string) (PacketSocketTable, error) {
	pack := PacketSocketTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return pack, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i==0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 9 {
			continue
		}

		entry := PacketSocketEntry{
			Socket: fields[0],
			RefCount: fields[1],
			Type: fields[2],
			Protocol: fields[3],
			Interface: fields[4],
			RecvQueue: fields[5],
			RecvMem: fields[6],
			User: fields[7],
			Inode: fields[8],
		}

		pack.Sockets = append(pack.Sockets, entry)


	}

	return pack, nil
}

func LinuxCollectUnixSocketTable(key string) (UnixSocketTable, error) {
	unix := UnixSocketTable{}

	content, err := os.ReadFile(DirMap[key]) 
	if err != nil {
		return unix, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i==0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 8 {
			continue
		}

		entry := UnixSocketEntry{
			Num: fields[0],
			RefCount: fields[1],
			Protocol: fields[2],
			Flags: fields[3],
			Type: fields[4],
			State: fields[5],
			Inode: fields[6],
			Path: fields[7],
		}

		unix.Sockets = append(unix.Sockets, entry)
	}

	return unix, nil
}

func LinuxCollectUDP(key string) (UDPConnectionTable, error) {
	udp := UDPConnectionTable{}

	content, err := os.ReadFile(DirMap[key]) // key: udp or udp6
	if err != nil {
		return udp, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i==0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		if len(fields) < 13 {
			continue
		}

		txrx := strings.SplitN(fields[4], ":", 2)
		timer := strings.SplitN(fields[5], ":", 2)

		entry := UDPConnectionEntry{
			connection: ConnectionEntry{
				Slot: fields[0],
				LocalAddress: fields[1],
				RemoteAddress: fields[2],
				State: fields[3],
				TxQueue: txrx[0],
				RxQueue: txrx[1],
				Timer: TimerInfo{
					Type: timer[0],
					Expires: timer[1],
					Retransmits: fields[6],
				},
				UID: fields[7],
				Timeout: fields[8],
				Inode: fields[9],
			},

			RefCount: fields[10],
			MemoryPointer: fields[11],
			Drops: fields[12],
		}

		udp.Connections = append(udp.Connections, entry)
	}

	return udp, nil
}

func LinuxCollectTCP(key string) (TCPConnectionTable, error) {  // key: tcp or tcp6
	tcp := TCPConnectionTable{}

	content, err := os.ReadFile(DirMap[key])
	if err != nil {
		return tcp, err
	}

	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if i==0 {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		if len(fields) < 10 {
			continue 
		}

		inode := strings.Join(fields[9:], " ")

		txrx := strings.SplitN(fields[4], ":", 2)
		timer := strings.SplitN(fields[5], ":", 2)

		entry := ConnectionEntry{
			Slot: fields[0],
			LocalAddress: fields[1],
			RemoteAddress: fields[2],
			State: fields[3],
			TxQueue: txrx[0],
			RxQueue: txrx[1],

			Timer: TimerInfo{
				Type: timer[0],
				Expires: timer[1],
				Retransmits: fields[6],
			},

			UID: fields[7],
			Timeout: fields[8],
			Inode: inode,
		}

		tcp.Connections = append(tcp.Connections, entry)
	}

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