// models.go for linux
package net

// -------------------------------------------------------------------------
// ARP
// -------------------------------------------------------------------------

type NetARP struct { // file: /proc/net/ARP
	ip_address string
	hw_type    string
	hw_address string
	device     string
}


// -------------------------------------------------------------------------
// dev
// -------------------------------------------------------------------------

type NetDev struct { // file: /proc/net/dev
	received  NetDevDataR
	transmitted NetDevDataT
}

type NetDevDataR struct {
	interface_ [2]string
	bytes [2]string 
	packets [2]string 
	err [2]string 
	drop [2]string 
	fifo [2]string 
	frame [2]string 
	compressed [2]string 
	multicast [2]string 
}

type NetDevDataT struct {
	interface_ [2]string
	bytes [2]string 
	packets [2]string 
	err [2]string 
	drop [2]string 
	fifo [2]string 
	colls [2]string
	carrier [2]string
	compressed [2]string 
}


// -------------------------------------------------------------------------
// connector
// -------------------------------------------------------------------------

type NetConnector struct { // file: /proc/net/connector
	cn_proc string
}

// -------------------------------------------------------------------------
// if_inet6
// -------------------------------------------------------------------------

type NetIfINet6 struct { // file: /proc/net/if_inet6
	ipv6_address    string
	interface_index string
	pref_len        string
	scope           string
	flags           string
	interface_name  string
}

// -------------------------------------------------------------------------
// IGMP
// -------------------------------------------------------------------------

type NetIGMP struct { // file: /proc/net/igmp
	idx      string
	device   string
	count    string
	querier  string
	group    string
	users    string
	timer    string
	reporter string
}

// -------------------------------------------------------------------------
// IGMP6
// -------------------------------------------------------------------------

type NetIGMP6 struct { // file: /proc/net/igmp6
	interface_index   string
	interface_name    string
	multicast_address string
	users             string
	flags             string
	timer             string
}

// -------------------------------------------------------------------------
// Netstat
// -------------------------------------------------------------------------

type NetNetStat struct { // file: /proc/net/netstat
	// TcpExt
	SyncookiesSent            string
	SyncookiesRecv            string
	SyncookiesFailed          string
	EmbryonicRsts             string
	PruneCalled               string
	RcvPruned                 string
	OfoPruned                 string
	OutOfWindowIcmps          string
	LockDroppedIcmps          string
	ArpFilter                 string
	TW                        string
	TWRecycled                string
	TWKilled                  string
	PAWSActive                string
	PAWSEstab                 string
	BeyondWindow              string
	TSEcrRejected             string
	PAWSOldAck                string
	PAWSTimewait              string
	DelayedACKs               string
	DelayedACKLocked          string
	DelayedACKLost            string
	ListenOverflows           string
	ListenDrops               string
	TCPHPHits                 string
	TCPPureAcks               string
	TCPHPAcks                 string
	TCPRenoRecovery           string
	TCPSackRecovery           string
	TCPSACKReneging           string
	TCPSACKReorder            string
	TCPRenoReorder            string
	TCPTSReorder              string
	TCPFullUndo               string
	TCPPartialUndo            string
	TCPDSACKUndo              string
	TCPLossUndo               string
	TCPLostRetransmit         string
	TCPRenoFailures           string
	TCPSackFailures           string
	TCPLossFailures           string
	TCPFastRetrans            string
	TCPSlowStartRetrans       string
	TCPTimeouts               string
	TCPLossProbes             string
	TCPLossProbeRecovery      string
	TCPRenoRecoveryFail       string
	TCPSackRecoveryFail       string
	TCPRcvCollapsed           string
	TCPBacklogCoalesce        string
	TCPDSACKOldSent           string
	TCPDSACKOfoSent           string
	TCPDSACKRecv              string
	TCPDSACKOfoRecv           string
	TCPAbortOnData            string
	TCPAbortOnClose           string
	TCPAbortOnMemory          string
	TCPAbortOnTimeout         string
	TCPAbortOnLinger          string
	TCPAbortFailed            string
	TCPMemoryPressures        string
	TCPMemoryPressuresChrono  string
	TCPSACKDiscard            string
	TCPDSACKIgnoredOld        string
	TCPDSACKIgnoredNoUndo     string
	TCPSpuriousRTOs           string
	TCPMD5NotFound            string
	TCPMD5Unexpected          string
	TCPMD5Failure             string
	TCPSackShifted            string
	TCPSackMerged             string
	TCPSackShiftFallback      string
	TCPBacklogDrop            string
	PFMemallocDrop            string
	TCPMinTTLDrop             string
	TCPDeferAcceptDrop        string
	IPReversePathFilter       string
	TCPTimeWaitOverflow       string
	TCPReqQFullDoCookies      string
	TCPReqQFullDrop           string
	TCPRetransFail            string
	TCPRcvCoalesce            string
	TCPOFOQueue               string
	TCPOFODrop                string
	TCPOFOMerge               string
	TCPChallengeACK           string
	TCPSYNChallenge           string
	TCPFastOpenActive         string
	TCPFastOpenActiveFail     string
	TCPFastOpenPassive        string
	TCPFastOpenPassiveFail    string
	TCPFastOpenListenOverflow string
	TCPFastOpenCookieReqd     string
	TCPFastOpenBlackhole      string
	TCPSpuriousRtxHostQueues  string
	BusyPollRxPackets         string
	TCPAutoCorking            string
	TCPFromZeroWindowAdv      string
	TCPToZeroWindowAdv        string
	TCPWantZeroWindowAdv      string
	TCPSynRetrans             string
	TCPOrigDataSent           string
	TCPHystartTrainDetect     string
	TCPHystartTrainCwnd       string
	TCPHystartDelayDetect     string
	TCPHystartDelayCwnd       string
	TCPACKSkippedSynRecv      string
	TCPACKSkippedPAWS         string
	TCPACKSkippedSeq          string
	TCPACKSkippedFinWait2     string
	TCPACKSkippedTimeWait     string
	TCPACKSkippedChallenge    string
	TCPWinProbe               string
	TCPKeepAlive              string
	TCPMTUPFail               string
	TCPMTUPSuccess            string
	TCPDelivered              string
	TCPDeliveredCE            string
	TCPAckCompressed          string
	TCPZeroWindowDrop         string
	TCPRcvQDrop               string
	TCPWqueueTooBig           string
	TCPFastOpenPassiveAltKey  string
	TcpTimeoutRehash          string
	TcpDuplicateDataRehash    string
	TCPDSACKRecvSegs          string
	TCPDSACKIgnoredDubious    string
	TCPMigrateReqSuccess      string
	TCPMigrateReqFailure      string
	TCPPLBRehash              string
	TCPAORequired             string
	TCPAOBad                  string
	TCPAOKeyNotFound          string
	TCPAOGood                 string
	TCPAODroppedIcmps         string

	// IpExt
	InNoRoutes      string
	InTruncatedPkts string
	InMcastPkts     string
	OutMcastPkts    string
	InBcastPkts     string
	OutBcastPkts    string
	InOctets        string
	OutOctets       string
	InMcastOctets   string
	OutMcastOctets  string
	InBcastOctets   string
	OutBcastOctets  string
	InCsumErrors    string
	InNoECTPkts     string
	InECT1Pkts      string
	InECT0Pkts      string
	InCEPkts        string
	ReasmOverlaps   string
}

// -------------------------------------------------------------------------
// FIB-Trie Stat
// -------------------------------------------------------------------------

type NetFibTrieStat struct { // file: /proc/net/fib_triestat
	LeafSizeBytes  string
	TNodeSizeBytes string

	Main  FibTrieSection
	Local FibTrieSection
}

type FibTrieSection struct {
	AverageDepth  string
	MaxDepth      string
	Leaves        string
	Prefixes      string
	InternalNodes string

	Depth1 string
	Depth2 string
	Depth3 string

	Pointers    string
	NullPtrs    string
	TotalSizeKB string

	Counters FibTrieCounters
}

type FibTrieCounters struct {
	Gets                string
	Backtracks          string
	SemanticMatchPassed string
	SemanticMatchMiss   string
	NullNodeHit         string
	SkippedNodeResize   string
}

// -------------------------------------------------------------------------
// Netlink
// -------------------------------------------------------------------------

type NetNetlink struct { // file: /proc/net/netlink
	sk     string
	eth    string
	pid    string
	groups string
	rmem   string
	wmem   string
	dump   string
	locks  string
	drops  string
	inode  string
}

// -------------------------------------------------------------------------
// Packet
// -------------------------------------------------------------------------

type NetPacket struct { // file: /proc/net/packet
	sk     string
	refcnt string
	type_  string
	proto  string
	iface  string
	r      string
	rmem   string
	user   string
	inode  string
}

// -------------------------------------------------------------------------
// Protocols
// -------------------------------------------------------------------------

type NetProtocols struct { // file: /proc/net/protocols
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

type NetPtype struct { // file: /proc/net/ptype
	type_    string
	device   string
	function string
}

// -------------------------------------------------------------------------
// Route
// -------------------------------------------------------------------------

type NetRoute struct { // file: /proc/net/route
	iface       string
	destination string
	gateway     string
	flags       string
	refcnt      string
	use         string
	metric      string
	mask        string
	mtu         string
	window      string
	irit        string
}

// -------------------------------------------------------------------------
// rt_cache
// -------------------------------------------------------------------------

type NetRtCache struct { // file: /proc/net/rt_cache
	iface       string
	destination string
	gateway     string
	flags       string
	refcnt      string
	use         string
	metric      string
	source      string
	mtu         string
	window      string
	irit        string
	tos         string
	hhref       string
	hhuptod     string
	specdest    string
}

// -------------------------------------------------------------------------
// SNMP
// -------------------------------------------------------------------------

type NetSNMP struct { // file: /proc/net/snmp
	IP      NetSNMPIP
	ICMP    NetSNMPICMP
	ICMPMsg NetSNMPICMPMsg
	TCP     NetSNMPTCP
	UDP     NetSNMPUDP
	UDPLite NetSNMPUDPLite
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
	InDatagrams  string
	NoPorts      string
	InErrors     string
	OutDatagrams string
	RcvbufErrors string
	SndbufErrors string
	InCsumErrors string
	IgnoredMulti string
	MemErrors    string
}

type NetSNMPUDPLite struct {
	InDatagrams  string
	NoPorts      string
	InErrors     string
	OutDatagrams string
	RcvbufErrors string
	SndbufErrors string
	InCsumErrors string
	IgnoredMulti string
	MemErrors    string
}

// -------------------------------------------------------------------------
// SNMP
// -------------------------------------------------------------------------

type NetSNMP6 struct { // file: /proc/net/snmp6
	Metrics map[string]string
}

// -------------------------------------------------------------------------
// Sockstat
// -------------------------------------------------------------------------

type NetSockstat struct { // file: /proc/net/sockstat
	sockets  string
	tcp      string
	udp      string
	udp_lite string
	raw      string
	frag     string
}

type NetSockstat6 struct { // file: /proc/net/sockstat6
	tcp6      string
	udp6      string
	udp_lite6 string
	raw6      string
	frag6     string
}

// -------------------------------------------------------------------------
// TCP
// -------------------------------------------------------------------------

type NetTcp struct { // file: /proc/net/tcp
	sl            []string
	local_address []string
	rem_Address   []string
	st            []string
	tx_queue      []string
	rx_queue      []string
	tr            []string
	tm_when       []string
	retrnsmt      []string
	uid           []string
	timeout       []string
	inode         []string
}

// -------------------------------------------------------------------------
// TLS Stat
// -------------------------------------------------------------------------

type NetTLSStat struct { // file: /proc/net/tls_stat
	TlsCurrTxSw         string
	TlsCurrRxSw         string
	TlsCurrTxDevice     string
	TlsCurrRxDevice     string
	TlsTxSw             string
	TlsRxSw             string
	TlsTxDevice         string
	TlsRxDevice         string
	TlsDecryptError     string
	TlsRxDeviceResync   string
	TlsDecryptRetry     string
	TlsRxNoPadViolation string
	TlsRxRekeyOk        string
	TlsRxRekeyError     string
	TlsTxRekeyOk        string
	TlsTxRekeyError     string
	TlsRxRekeyReceived  string
}

// -------------------------------------------------------------------------
// UDP
// -------------------------------------------------------------------------

type NetUdp struct { // file: /proc/net/udp
	sl            []string
	local_address []string
	rem_Address   []string
	st            []string
	tx_queue      []string
	rx_queue      []string
	tr            []string
	tm_when       []string
	retrnsmt      []string
	uid           []string
	timeout       []string
	inode         []string
	ref           []string
	pointer       []string
	drops         []string
}

// -------------------------------------------------------------------------
// Unix
// -------------------------------------------------------------------------

type NetUnix struct { // file: /proc/net/unix
	num       []string
	ref_count []string
	protocol  []string
	flags     []string
	type_     []string
	st        []string
	inode     []string
	path      []string
}

// -------------------------------------------------------------------------
// xfrm_stat
// -------------------------------------------------------------------------

type NetXFRMStat struct { // file: /proc/net/xfrm_stat
	XfrmInError           string
	XfrmInBufferError     string
	XfrmInHdrError        string
	XfrmInNoStates        string
	XfrmInStateProtoError string
	XfrmInStateModeError  string
	XfrmInStateSeqError   string
	XfrmInStateExpired    string
	XfrmInStateMismatch   string
	XfrmInStateInvalid    string
	XfrmInTmplMismatch    string
	XfrmInNoPols          string
	XfrmInPolBlock        string
	XfrmInPolError        string

	XfrmOutError            string
	XfrmOutBundleGenError   string
	XfrmOutBundleCheckError string
	XfrmOutNoStates         string
	XfrmOutStateProtoError  string
	XfrmOutStateModeError   string
	XfrmOutStateSeqError    string
	XfrmOutStateExpired     string
	XfrmOutPolBlock         string
	XfrmOutPolDead          string
	XfrmOutPolError         string

	XfrmFwdHdrError      string
	XfrmOutStateInvalid  string
	XfrmAcquireError     string
	XfrmOutStateDirError string
	XfrmInStateDirError  string
	XfrmInIptfsError     string
	XfrmOutNoQueueSpace  string
}

// =====================================================================================
// =====================================================================================
// Search: /sys
// =====================================================================================
// =====================================================================================

type NetEth0 struct {
	addr_assign_type string
	addr_len string 
	address string
	broadcast string
	carrier string 
	carrier_changes string 
	carrier_down_count string 
	carrier_up_count string
	dev_id string 
	dev_port string 
	device NetEth0Device

}

type NetEth0_ struct {
	addr_assign_type string
	addr_len string 
	address string
	broadcast string
	carrier string 
	carrier_changes string 
	carrier_down_count string 
	carrier_up_count string
	dev_id string 
	dev_port string 

}

type NetEth0Device struct {
	channel_vp_mapping []string 
	channel NetEth0DeviceChannel
	class_id string
	client_monitor_conn_id string 
	client_monitor_latency string 
	client_monitor_pending string
	device string 
	device_id string 
	driver_override string 
	id string 
	in_intr_mask string
	in_read_bytes_avail string 
	in_read_index string
	in_write_bytes_avail string 
	in_write_index string 
	modalias string 
	monitor_id string
	
	net NetEth0_
	numa_node string 
	out_intr_mask string 
	out_read_bytes_avail string 
	out_read_index string 
	out_write_bytes_avail string 

	out_write_index string 
	power NetEth0DevicePower
	server_monitor_conn_id string 
	server_monitor_latency string

	server_monitor_pending string
	state string 
	subsystem NetEth0DeviceSubsystem
	
	uevent NetEth0DeviceUevent
	vendor string

}

type NetEth0DeviceNet struct {

}

type NetEth0DeviceSubsystem struct {
	drivers_autoprobe string
	hibernation string
}

type NetEth0DevicePower struct {
	control string
	runtime_active_time string 
	runtime_status string
	runtime_suspended_time string
}

type NetEth0DeviceChannel struct {
	channel map[string]NetEth0DeviceChannelSingle  // channels: 15  16  17  18  19  20  21  22  23  24  25  26  // note: my pc
}

type NetEth0DeviceUevent struct {
	driver string 
	modalias string
}

type NetEth0DeviceChannelSingle struct {
	cpu string 
	events string 
	in_mask string 
	interrupts string 
	intr_in_full string 
	intr_out_empty string 
	latency string 
	monitor_id string 
	out_full_first string 
	out_full_total string 
	out_mask string 
	pending string
	read_avail string 
	subchannel_id string 
	write_avail string
}
