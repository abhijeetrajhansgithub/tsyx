package net 

// -------------------------------------------------------------------------
// ARP
// -------------------------------------------------------------------------

type NetARP struct {  // file: /proc/net/ARP
	ip_address string 
	hw_type string 
	hw_address string 
	device string
}

// -------------------------------------------------------------------------
// Conenctor
// -------------------------------------------------------------------------

type NetConnector struct {  // file: /proc/net/connector
	cn_proc string
}

// -------------------------------------------------------------------------
// if_inet6
// -------------------------------------------------------------------------

type NetIfINet6 struct {  // file: /proc/net/if_inet6
	ipv6_address string 
	interface_index string 
	pref_len string 
	scope string 
	flags string 
	interface_name string
}

// -------------------------------------------------------------------------
// IGMP
// -------------------------------------------------------------------------


type NetIGMP struct {  // file: /proc/net/igmp
	idx string 
	device string 
	count string 
	querier string 
	group string 
	users string 
	timer string
	reporter string 
}

// -------------------------------------------------------------------------
// IGMP6
// -------------------------------------------------------------------------

type NetIGMP6 struct {  // file: /proc/net/igmp6
	interface_index  string 
	interface_name string 
	multicast_address string
	users string 
	flags string 
	timer string
}

// -------------------------------------------------------------------------
// Netstat
// -------------------------------------------------------------------------

type NetNetStat struct {  // file: /proc/net/netstat
	// TcpExt
	SyncookiesSent              string
	SyncookiesRecv              string
	SyncookiesFailed            string
	EmbryonicRsts               string
	PruneCalled                 string
	RcvPruned                   string
	OfoPruned                   string
	OutOfWindowIcmps            string
	LockDroppedIcmps            string
	ArpFilter                   string
	TW                          string
	TWRecycled                  string
	TWKilled                    string
	PAWSActive                  string
	PAWSEstab                   string
	BeyondWindow                string
	TSEcrRejected               string
	PAWSOldAck                  string
	PAWSTimewait                string
	DelayedACKs                 string
	DelayedACKLocked            string
	DelayedACKLost              string
	ListenOverflows             string
	ListenDrops                 string
	TCPHPHits                   string
	TCPPureAcks                 string
	TCPHPAcks                   string
	TCPRenoRecovery             string
	TCPSackRecovery             string
	TCPSACKReneging             string
	TCPSACKReorder              string
	TCPRenoReorder              string
	TCPTSReorder                string
	TCPFullUndo                 string
	TCPPartialUndo              string
	TCPDSACKUndo                string
	TCPLossUndo                 string
	TCPLostRetransmit           string
	TCPRenoFailures             string
	TCPSackFailures             string
	TCPLossFailures             string
	TCPFastRetrans              string
	TCPSlowStartRetrans         string
	TCPTimeouts                 string
	TCPLossProbes               string
	TCPLossProbeRecovery        string
	TCPRenoRecoveryFail         string
	TCPSackRecoveryFail         string
	TCPRcvCollapsed             string
	TCPBacklogCoalesce          string
	TCPDSACKOldSent             string
	TCPDSACKOfoSent             string
	TCPDSACKRecv                string
	TCPDSACKOfoRecv             string
	TCPAbortOnData              string
	TCPAbortOnClose             string
	TCPAbortOnMemory            string
	TCPAbortOnTimeout           string
	TCPAbortOnLinger            string
	TCPAbortFailed              string
	TCPMemoryPressures          string
	TCPMemoryPressuresChrono    string
	TCPSACKDiscard              string
	TCPDSACKIgnoredOld          string
	TCPDSACKIgnoredNoUndo       string
	TCPSpuriousRTOs             string
	TCPMD5NotFound              string
	TCPMD5Unexpected            string
	TCPMD5Failure               string
	TCPSackShifted              string
	TCPSackMerged               string
	TCPSackShiftFallback        string
	TCPBacklogDrop              string
	PFMemallocDrop              string
	TCPMinTTLDrop               string
	TCPDeferAcceptDrop          string
	IPReversePathFilter         string
	TCPTimeWaitOverflow         string
	TCPReqQFullDoCookies        string
	TCPReqQFullDrop             string
	TCPRetransFail              string
	TCPRcvCoalesce              string
	TCPOFOQueue                 string
	TCPOFODrop                  string
	TCPOFOMerge                 string
	TCPChallengeACK             string
	TCPSYNChallenge             string
	TCPFastOpenActive           string
	TCPFastOpenActiveFail       string
	TCPFastOpenPassive          string
	TCPFastOpenPassiveFail      string
	TCPFastOpenListenOverflow   string
	TCPFastOpenCookieReqd       string
	TCPFastOpenBlackhole        string
	TCPSpuriousRtxHostQueues    string
	BusyPollRxPackets           string
	TCPAutoCorking              string
	TCPFromZeroWindowAdv        string
	TCPToZeroWindowAdv          string
	TCPWantZeroWindowAdv        string
	TCPSynRetrans               string
	TCPOrigDataSent             string
	TCPHystartTrainDetect       string
	TCPHystartTrainCwnd         string
	TCPHystartDelayDetect       string
	TCPHystartDelayCwnd         string
	TCPACKSkippedSynRecv        string
	TCPACKSkippedPAWS           string
	TCPACKSkippedSeq            string
	TCPACKSkippedFinWait2       string
	TCPACKSkippedTimeWait       string
	TCPACKSkippedChallenge      string
	TCPWinProbe                 string
	TCPKeepAlive                string
	TCPMTUPFail                 string
	TCPMTUPSuccess              string
	TCPDelivered                string
	TCPDeliveredCE              string
	TCPAckCompressed            string
	TCPZeroWindowDrop           string
	TCPRcvQDrop                 string
	TCPWqueueTooBig             string
	TCPFastOpenPassiveAltKey    string
	TcpTimeoutRehash            string
	TcpDuplicateDataRehash      string
	TCPDSACKRecvSegs            string
	TCPDSACKIgnoredDubious      string
	TCPMigrateReqSuccess        string
	TCPMigrateReqFailure        string
	TCPPLBRehash                string
	TCPAORequired               string
	TCPAOBad                    string
	TCPAOKeyNotFound            string
	TCPAOGood                   string
	TCPAODroppedIcmps           string

	// IpExt
	InNoRoutes       string
	InTruncatedPkts  string
	InMcastPkts      string
	OutMcastPkts     string
	InBcastPkts      string
	OutBcastPkts     string
	InOctets         string
	OutOctets        string
	InMcastOctets    string
	OutMcastOctets   string
	InBcastOctets    string
	OutBcastOctets   string
	InCsumErrors     string
	InNoECTPkts      string
	InECT1Pkts       string
	InECT0Pkts       string
	InCEPkts         string
	ReasmOverlaps    string
}

// -------------------------------------------------------------------------
// FIB-Trie Stat
// -------------------------------------------------------------------------

type NetFibTrieStat struct {  // file: /proc/net/fib_triestat
	LeafSizeBytes  string
	TNodeSizeBytes string

	Main  FibTrieSection
	Local FibTrieSection
}

type FibTrieSection struct {
	AverageDepth string
	MaxDepth     string
	Leaves       string
	Prefixes     string
	InternalNodes string

	Depth1 string
	Depth2 string
	Depth3 string

	Pointers string
	NullPtrs string
	TotalSizeKB string

	Counters FibTrieCounters
}

type FibTrieCounters struct {
	Gets                 string
	Backtracks           string
	SemanticMatchPassed  string
	SemanticMatchMiss    string
	NullNodeHit          string
	SkippedNodeResize    string
}


// -------------------------------------------------------------------------
// Netlink
// -------------------------------------------------------------------------

type NetNetlink struct {  // file: /proc/net/netlink
	sk string 
	eth string 
	pid string 
	groups string 
	rmem string 
	wmem string 
	dump string 
	locks string 
	drops string 
	inode string
}


// -------------------------------------------------------------------------
// Packet
// -------------------------------------------------------------------------

type NetPacket struct {  // file: /proc/net/packet
	sk string 
	refcnt string 
	type_ string 
	proto string 
	iface string
	r string 
	rmem string
	user string 
	inode string
}

// -------------------------------------------------------------------------
// Protocols
// -------------------------------------------------------------------------

type NetProtocols struct {
	Protocols []NetProtocol
}

type NetProtocol struct {
	Protocol string

	Size    string
	Sockets string
	Memory  string
	Press   string
	MaxHdr  string
	Slab    string
	Module  string

	Capabilities map[string]string
}

// -------------------------------------------------------------------------
// ptype
// -------------------------------------------------------------------------

type NetPtype struct {
	type_ string 
	device string
	function string 
}

// -------------------------------------------------------------------------
// Route
// -------------------------------------------------------------------------

type NetRoute struct {
	iface string 
	destination string 
	gateway string 
	flags string
	refcnt string 
	use string 
	metric string 
	mask string 
	mtu string
	window string 
	irit string
}


// -------------------------------------------------------------------------
// rt_cache
// -------------------------------------------------------------------------

type NetRtCache struct {
	iface string 
	destination string 
	gateway string 
	flags string
	refcnt string 
	use string 
	metric string 
	source string 
	mtu string
	window string 
	irit string
	tos string 
	hhref string 
	hhuptod string 
	specdest string
}

// -------------------------------------------------------------------------
// SNMP
// -------------------------------------------------------------------------

type NetSNMP struct {
	IP       NetSNMPIP
	ICMP     NetSNMPICMP
	ICMPMsg  NetSNMPICMPMsg
	TCP      NetSNMPTCP
	UDP      NetSNMPUDP
	UDPLite  NetSNMPUDPLite
}

type NetSNMPIP struct {
	Forwarding      string
	DefaultTTL      string
	InReceives      string
	InHdrErrors     string
	InAddrErrors    string
	ForwDatagrams   string
	InUnknownProtos string
	InDiscards      string
	InDelivers      string
	OutRequests     string
	OutDiscards     string
	OutNoRoutes     string
	ReasmTimeout    string
	ReasmReqds      string
	ReasmOKs        string
	ReasmFails      string
	FragOKs         string
	FragFails       string
	FragCreates     string
	OutTransmits    string
}

type NetSNMPICMP struct {
	InMsgs             string
	InErrors           string
	InCsumErrors       string
	InDestUnreachs     string
	InTimeExcds        string
	InParmProbs        string
	InSrcQuenchs       string
	InRedirects        string
	InEchos            string
	InEchoReps         string
	InTimestamps       string
	InTimestampReps    string
	InAddrMasks        string
	InAddrMaskReps     string
	OutMsgs            string
	OutErrors          string
	OutRateLimitGlobal string
	OutRateLimitHost   string
	OutDestUnreachs    string
	OutTimeExcds       string
	OutParmProbs       string
	OutSrcQuenchs      string
	OutRedirects       string
	OutEchos           string
	OutEchoReps        string
	OutTimestamps      string
	OutTimestampReps   string
	OutAddrMasks       string
	OutAddrMaskReps    string
}

type NetSNMPICMPMsg struct {
	Types map[string]string
}

type NetSNMPTCP struct {
	RtoAlgorithm string
	RtoMin       string
	RtoMax       string
	MaxConn      string
	ActiveOpens  string
	PassiveOpens string
	AttemptFails string
	EstabResets  string
	CurrEstab    string
	InSegs       string
	OutSegs      string
	RetransSegs  string
	InErrs       string
	OutRsts      string
	InCsumErrors string
}

type NetSNMPUDP struct {
	InDatagrams   string
	NoPorts       string
	InErrors      string
	OutDatagrams  string
	RcvbufErrors  string
	SndbufErrors  string
	InCsumErrors  string
	IgnoredMulti  string
	MemErrors     string
}

type NetSNMPUDPLite struct {
	InDatagrams   string
	NoPorts       string
	InErrors      string
	OutDatagrams  string
	RcvbufErrors  string
	SndbufErrors  string
	InCsumErrors  string
	IgnoredMulti  string
	MemErrors     string
}

// -------------------------------------------------------------------------
// SNMP
// -------------------------------------------------------------------------

type NetSNMP6 struct {
	Metrics map[string]string
}

// -------------------------------------------------------------------------
// Sockstat
// -------------------------------------------------------------------------

type NetSockstat struct {
	sockets string 
	tcp string 
	udp string 
	udp_lite string 
	raw string 
	frag string
}