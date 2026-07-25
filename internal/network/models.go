// Package net provides structured models for Linux networking data sourced
// from /proc, /sys, and /etc. Models are purely descriptive: no parsing,
// calculation, or presentation logic lives here.
package net

// =====================================================================================
// Interfaces
// =====================================================================================

// ARPEntry represents a single row of the kernel ARP cache.
type ARPEntry struct {
	IPAddress     string
	HardwareType  string
	HardwareAddr  string
	Flags         string
	Mask          string
	Device        string
}

// ARPTable holds all entries read from /proc/net/arp.
type ARPTable struct { // file: /proc/net/arp
	Entries []ARPEntry
}

// TrafficCounters holds counters common to both received and transmitted
// interface traffic.
type TrafficCounters struct {
	Bytes      string
	Packets    string
	Errors     string
	Drop       string
	Fifo       string
	Compressed string
}

// ReceivedTraffic holds inbound traffic counters for an interface.
type ReceivedTraffic struct {
	TrafficCounters
	Frame     string
	Multicast string
}

// TransmittedTraffic holds outbound traffic counters for an interface.
type TransmittedTraffic struct {
	TrafficCounters
	Collisions string
	Carrier    string
}

// InterfaceTraffic represents a single interface row of RX/TX counters.
type InterfaceTraffic struct {
	Interface   string
	Received    ReceivedTraffic
	Transmitted TransmittedTraffic
}

// NetworkDeviceStats holds per-interface traffic statistics.
type NetworkDeviceStats struct { // file: /proc/net/dev
	Interfaces []InterfaceTraffic
}

// IPv6Address represents a single IPv6 address assigned to an interface.
type IPv6Address struct {
	Address        string
	InterfaceIndex string
	PrefixLength   string
	Scope          string
	Flags          string
	InterfaceName  string
}

// IPv6AddressTable holds all entries read from /proc/net/if_inet6.
type IPv6AddressTable struct { // file: /proc/net/if_inet6
	Addresses []IPv6Address
}

// PacketTypeHandler represents a single registered packet handler.
type PacketTypeHandler struct {
	Type     string
	Device   string
	Function string
	Module   string
}

// PacketTypeHandlerTable holds all entries read from /proc/net/ptype.
type PacketTypeHandlerTable struct { // file: /proc/net/ptype
	Handlers []PacketTypeHandler
}

// MulticastGroup represents a single IPv4 multicast group membership.
type MulticastGroup struct {
	Index    string
	Device   string
	Count    string
	Querier  string
	Group    string
	Users    string
	Timer    string
	Reporter string
}

// MulticastGroupTable holds all entries read from /proc/net/igmp.
type MulticastGroupTable struct { // file: /proc/net/igmp
	Groups []MulticastGroup
}

// MulticastGroup6 represents a single IPv6 multicast group membership.
type MulticastGroup6 struct {
	InterfaceIndex string
	InterfaceName  string
	Address        string
	Users          string
	Flags          string
	Timer          string
}

// MulticastGroupTable6 holds all entries read from /proc/net/igmp6.
type MulticastGroupTable6 struct { // file: /proc/net/igmp6
	Groups []MulticastGroup6
}

// CarrierInfo describes the link-carrier state of a network interface.
type CarrierInfo struct {
	Status     string
	Changes    string
	DownCount  string
	UpCount    string
}

// DevicePowerInfo describes runtime power-management state of a device.
type DevicePowerInfo struct {
	Control               string
	RuntimeActiveTime     string
	RuntimeStatus         string
	RuntimeSuspendedTime  string
}

// DeviceSubsystemInfo describes subsystem-level driver behavior for a device.
type DeviceSubsystemInfo struct {
	DriversAutoprobe string
	Hibernation      string
}

// DeviceUeventInfo describes the uevent metadata reported for a device.
type DeviceUeventInfo struct {
	Driver   string
	Modalias string
}

// DeviceMonitorInfo describes paravirtual client/server monitor channel state.
type DeviceMonitorInfo struct {
	ClientConnectionID string
	ClientLatency      string
	ClientPending       string
	ServerConnectionID string
	ServerLatency       string
	ServerPending       string
}

// DeviceRingBuffer describes shared-memory ring buffer state for a device.
type DeviceRingBuffer struct {
	InInterruptMask     string
	InReadBytesAvail    string
	InReadIndex         string
	InWriteBytesAvail   string
	InWriteIndex        string
	OutInterruptMask    string
	OutReadBytesAvail   string
	OutReadIndex        string
	OutWriteBytesAvail  string
	OutWriteIndex       string
}

// InterfaceChannel describes a single hardware/software channel exposed by
// an interface's underlying device (e.g. a VMBus or NIC queue channel).
type InterfaceChannel struct {
	CPU           string
	Events        string
	InMask        string
	Interrupts    string
	IntrInFull    string
	IntrOutEmpty  string
	Latency       string
	MonitorID     string
	OutFullFirst  string
	OutFullTotal  string
	OutMask       string
	Pending       string
	ReadAvail     string
	SubchannelID  string
	WriteAvail    string
}

// InterfaceDevice describes the underlying device backing a network
// interface, as exposed under /sys/class/net/<iface>/device.
type InterfaceDevice struct {
	ID                string
	DeviceID          string
	ClassID           string
	DriverOverride    string
	Modalias          string
	NUMANode          string
	State             string
	Vendor            string
	Power             DevicePowerInfo
	Subsystem         DeviceSubsystemInfo
	Uevent            DeviceUeventInfo
	Monitor           DeviceMonitorInfo
	RingBuffer        DeviceRingBuffer
	Channels          map[string]InterfaceChannel
}

// NetworkInterface is a generic model for any interface exposed under
// /sys/class/net/<iface>, replacing per-interface-name structs (e.g. an
// "eth0"-specific struct) with a single reusable shape.
type NetworkInterface struct { // dir: /sys/class/net/<interface>
	Name             string
	AddressAssignType string
	AddressLength     string
	Address           string
	Broadcast         string
	Carrier           CarrierInfo
	DevID             string
	DevPort           string
	Device            InterfaceDevice
}

// NetworkInterfaces holds all interfaces discovered under /sys/class/net.
type NetworkInterfaces struct {
	Interfaces []NetworkInterface
}

// =====================================================================================
// Routing
// =====================================================================================

// RouteEntry represents a single row of the IPv4 routing table.
type RouteEntry struct {
	Interface   string
	Destination string
	Gateway     string
	Flags       string
	RefCount    string
	Use         string
	Metric      string
	Mask        string
	MTU         string
	Window      string
	IRTT        string
}

// RouteTable holds all entries read from /proc/net/route.
type RouteTable struct { // file: /proc/net/route
	Routes []RouteEntry
}

// RouteCacheEntry represents a single row of the (legacy) route cache.
type RouteCacheEntry struct {
	Interface        string
	Destination      string
	Gateway          string
	Flags            string
	RefCount         string
	Use              string
	Metric           string
	Source           string
	MTU              string
	Window           string
	IRTT             string
	TOS              string
	HardwareHeaderRef string
	HardwareHeaderUpToDate string
	SpecificDestination   string
}

// RouteCacheTable holds all entries read from /proc/net/rt_cache.
type RouteCacheTable struct { // file: /proc/net/rt_cache
	Routes []RouteCacheEntry
}

// FIBTrieCounters holds hit/miss counters for a FIB trie lookup structure.
type FIBTrieCounters struct {
	Gets                string
	Backtracks          string
	SemanticMatchPassed string
	SemanticMatchMiss   string
	NullNodeHit         string
	SkippedNodeResize   string
}

// FIBTrieSection describes a single FIB trie (e.g. Main or Local).
type FIBTrieSection struct {
	AverageDepth  string
	MaxDepth      string
	Leaves        string
	Prefixes      string
	InternalNodes string
	Depth1        string
	Depth2        string
	Depth3        string
	Pointers      string
	NullPtrs      string
	TotalSizeKB   string
	Counters      FIBTrieCounters
}

// FIBTrieStatistics holds statistics for the kernel's FIB trie routing
// tables.
type FIBTrieStatistics struct { // file: /proc/net/fib_triestat
	LeafSizeBytes  string
	TNodeSizeBytes string
	Main           FIBTrieSection
	Local          FIBTrieSection
}

// =====================================================================================
// Connections
// =====================================================================================

// TimerInfo describes the retransmission/expiry timer state of a socket.
type TimerInfo struct {
	Type        string
	Expires     string
	Retransmits string
}

// ConnectionEntry represents a single row shared by the TCP and UDP socket
// tables (v4 and v6), since /proc/net/{tcp,tcp6,udp,udp6} share this base
// layout.
type ConnectionEntry struct {
	Slot          string
	LocalAddress  string
	RemoteAddress string
	State         string
	TxQueue       string
	RxQueue       string
	Timer         TimerInfo
	UID           string
	Timeout       string
	Inode         string
}

// TCPConnectionTable holds all entries read from /proc/net/tcp or
// /proc/net/tcp6.
type TCPConnectionTable struct { // files: /proc/net/tcp, /proc/net/tcp6
	Connections []ConnectionEntry
}

// UDPConnectionEntry extends ConnectionEntry with fields specific to UDP
// sockets.
type UDPConnectionEntry struct {
	connection ConnectionEntry
	RefCount      string
	MemoryPointer string
	Drops         string
}

// UDPConnectionTable holds all entries read from /proc/net/udp or
// /proc/net/udp6.
type UDPConnectionTable struct { // files: /proc/net/udp, /proc/net/udp6
	Connections []UDPConnectionEntry
}

// UnixSocketEntry represents a single row of the Unix domain socket table.
type UnixSocketEntry struct {
	Num       string
	RefCount  string
	Protocol  string
	Flags     string
	Type      string
	State     string
	Inode     string
	Path      string
}

// UnixSocketTable holds all entries read from /proc/net/unix.
type UnixSocketTable struct { // file: /proc/net/unix
	Sockets []UnixSocketEntry
}

// PacketSocketEntry represents a single row of the raw packet socket table.
type PacketSocketEntry struct {
	Socket     string
	RefCount   string
	Type       string
	Protocol   string
	Interface  string
	RecvQueue  string
	RecvMem    string
	User       string
	Inode      string
}

// PacketSocketTable holds all entries read from /proc/net/packet.
type PacketSocketTable struct { // file: /proc/net/packet
	Sockets []PacketSocketEntry
}

// NetlinkSocket represents a single row of the netlink socket table.
type NetlinkSocket struct {
	Socket    string
	Protocol  string
	PID       string
	Groups    string
	RecvMem   string
	SendMem   string
	Dump      string
	Locks     string
	Drops     string
	Inode     string
}

// NetlinkTable holds all entries read from /proc/net/netlink.
type NetlinkTable struct { // file: /proc/net/netlink
	Sockets []NetlinkSocket
}

// ProcessConnector holds data read from the kernel's process event
// connector interface.
type ProcessConnector struct { // file: /proc/net/connector
	Name string
	ID string
}

// =====================================================================================
// Statistics
// =====================================================================================

// TCPExtendedStatistics holds Linux-specific extended TCP counters, as
// reported under the "TcpExt" section of /proc/net/netstat.
type TCPExtendedStatistics struct {
	SyncookiesSent            string
	SyncookiesRecv            string
	SyncookiesFailed          string
	EmbryonicRsts             string
	PruneCalled               string
	RcvPruned                 string
	OfoPruned                 string
	OutOfWindowICMPs          string
	LockDroppedICMPs          string
	ArpFilter                 string
	TimeWait                  string
	TimeWaitRecycled          string
	TimeWaitKilled            string
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
	HPHits                    string
	PureAcks                  string
	HPAcks                    string
	RenoRecovery              string
	SackRecovery              string
	SACKReneging              string
	SACKReorder               string
	RenoReorder               string
	TSReorder                 string
	FullUndo                  string
	PartialUndo                string
	DSACKUndo                 string
	LossUndo                  string
	LostRetransmit            string
	RenoFailures              string
	SackFailures              string
	LossFailures              string
	FastRetrans               string
	SlowStartRetrans          string
	Timeouts                  string
	LossProbes                string
	LossProbeRecovery         string
	RenoRecoveryFail          string
	SackRecoveryFail          string
	RcvCollapsed              string
	BacklogCoalesce           string
	DSACKOldSent              string
	DSACKOfoSent              string
	DSACKRecv                 string
	DSACKOfoRecv              string
	AbortOnData               string
	AbortOnClose              string
	AbortOnMemory             string
	AbortOnTimeout            string
	AbortOnLinger             string
	AbortFailed               string
	MemoryPressures           string
	MemoryPressuresChrono     string
	SACKDiscard               string
	DSACKIgnoredOld           string
	DSACKIgnoredNoUndo        string
	SpuriousRTOs              string
	MD5NotFound               string
	MD5Unexpected             string
	MD5Failure                string
	SackShifted               string
	SackMerged                string
	SackShiftFallback         string
	BacklogDrop               string
	PFMemallocDrop            string
	MinTTLDrop                string
	DeferAcceptDrop           string
	IPReversePathFilter       string
	TimeWaitOverflow          string
	ReqQFullDoCookies         string
	ReqQFullDrop              string
	RetransFail               string
	RcvCoalesce               string
	OFOQueue                  string
	OFODrop                   string
	OFOMerge                  string
	ChallengeACK              string
	SYNChallenge              string
	FastOpenActive            string
	FastOpenActiveFail        string
	FastOpenPassive           string
	FastOpenPassiveFail       string
	FastOpenListenOverflow    string
	FastOpenCookieReqd        string
	FastOpenBlackhole         string
	FastOpenPassiveAltKey     string
	SpuriousRtxHostQueues     string
	BusyPollRxPackets         string
	AutoCorking               string
	FromZeroWindowAdv         string
	ToZeroWindowAdv           string
	WantZeroWindowAdv         string
	SynRetrans                string
	OrigDataSent              string
	HystartTrainDetect        string
	HystartTrainCwnd          string
	HystartDelayDetect        string
	HystartDelayCwnd          string
	ACKSkippedSynRecv         string
	ACKSkippedPAWS            string
	ACKSkippedSeq             string
	ACKSkippedFinWait2        string
	ACKSkippedTimeWait        string
	ACKSkippedChallenge       string
	WinProbe                  string
	KeepAlive                 string
	MTUPFail                  string
	MTUPSuccess               string
	Delivered                 string
	DeliveredCE               string
	AckCompressed             string
	ZeroWindowDrop            string
	RcvQDrop                  string
	WqueueTooBig              string
	TimeoutRehash             string
	DuplicateDataRehash       string
	DSACKRecvSegs             string
	DSACKIgnoredDubious       string
	MigrateReqSuccess         string
	MigrateReqFailure         string
	PLBRehash                 string
	AORequired                string
	AOBad                     string
	AOKeyNotFound             string
	AOGood                    string
	AODroppedICMPs            string
}

// IPExtendedStatistics holds extended IP counters, as reported under the
// "IpExt" section of /proc/net/netstat.
type IPExtendedStatistics struct {
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

// NetworkExtendedStatistics holds extended TCP/IP statistics read from
// /proc/net/netstat.
type NetworkExtendedStatistics struct { // file: /proc/net/netstat
	TCP TCPExtendedStatistics
	IP  IPExtendedStatistics
}

// SNMPIPStatistics holds IP-layer counters from /proc/net/snmp.
type SNMPIPStatistics struct {
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

// SNMPICMPStatistics holds ICMP counters from /proc/net/snmp.
type SNMPICMPStatistics struct {
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

// SNMPICMPMessageStatistics holds per-message-type ICMP counters, which are
// dynamic and thus modeled as a map rather than a fixed schema.
type SNMPICMPMessageStatistics struct {
	Types map[string]string
}

// SNMPTCPStatistics holds TCP-layer counters from /proc/net/snmp.
type SNMPTCPStatistics struct {
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

// UDPStatistics holds counters shared by the UDP and UDP-Lite sections of
// /proc/net/snmp, since both protocols report an identical set of fields.
type UDPStatistics struct {
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

// SNMPStatistics holds protocol statistics read from /proc/net/snmp.
type SNMPStatistics struct { // file: /proc/net/snmp
	IP      SNMPIPStatistics
	ICMP    SNMPICMPStatistics
	ICMPMsg SNMPICMPMessageStatistics
	TCP     SNMPTCPStatistics
	UDP     UDPStatistics
	UDPLite UDPStatistics
}

// SNMP6Statistics holds IPv6 protocol statistics read from
// /proc/net/snmp6. The kernel reports a variable, version-dependent set of
// metrics, so they are modeled as a map rather than a fixed schema.
type SNMP6Statistics struct { // file: /proc/net/snmp6
	Metrics map[string]string
}

// SocketTypeCounters holds per-protocol socket usage counters shared by the
// IPv4 and IPv6 socket summary files.
type SocketTypeCounters struct {
	TCP     string
	UDP     string
	UDPLite string
	RAW     string
	Frag    string
}

// SocketStatistics holds IPv4 socket usage summary data.
type SocketStatistics struct { // file: /proc/net/sockstat
	Sockets  string
	Counters SocketTypeCounters
}

// SocketStatistics6 holds IPv6 socket usage summary data.
type SocketStatistics6 struct { // file: /proc/net/sockstat6
	Counters SocketTypeCounters
}

// ProtocolStatistics represents a single protocol row read from
// /proc/net/protocols. Capabilities are modeled as a map since the
// available capability flags vary by kernel build.
type ProtocolStatistics struct {
	Protocol     string
	Size         string
	Sockets      string
	Memory       string
	Press        string
	MaxHeader    string
	Slab         string
	Module       string
	Capabilities map[string]string
}

// ProtocolStatisticsTable holds all entries read from /proc/net/protocols.
type ProtocolStatisticsTable struct { // file: /proc/net/protocols
	Protocols []ProtocolStatistics
}

// TLSSessionCounts holds a count of TLS sessions broken down by direction
// and offload mode. Reused for both current (in-progress) and cumulative
// totals.
type TLSSessionCounts struct {
	TxSoftware string
	RxSoftware string
	TxDevice   string
	RxDevice   string
}

// TLSErrorCounts holds TLS decryption/validation error counters.
type TLSErrorCounts struct {
	DecryptError     string
	RxDeviceResync   string
	DecryptRetry     string
	RxNoPadViolation string
}

// TLSRekeyCounts holds TLS session rekey counters.
type TLSRekeyCounts struct {
	RxOK       string
	RxError    string
	TxOK       string
	TxError    string
	RxReceived string
}

// TLSStatistics holds kernel TLS offload statistics read from
// /proc/net/tls_stat.
type TLSStatistics struct { // file: /proc/net/tls_stat
	Current TLSSessionCounts
	Total   TLSSessionCounts
	Errors  TLSErrorCounts
	Rekey   TLSRekeyCounts
}

// =====================================================================================
// Security
// =====================================================================================

// XFRMInboundStatistics holds inbound IPsec/XFRM error and drop counters.
type XFRMInboundStatistics struct {
	Error             string
	BufferError       string
	HeaderError       string
	NoStates          string
	StateProtoError   string
	StateModeError    string
	StateSeqError     string
	StateExpired      string
	StateMismatch     string
	StateInvalid      string
	StateDirError     string
	TemplateMismatch  string
	NoPolicies        string
	PolicyBlock       string
	PolicyError       string
	IptfsError        string
}

// XFRMOutboundStatistics holds outbound IPsec/XFRM error and drop counters.
type XFRMOutboundStatistics struct {
	Error             string
	BundleGenError    string
	BundleCheckError  string
	NoStates          string
	StateProtoError   string
	StateModeError    string
	StateSeqError     string
	StateExpired      string
	StateInvalid      string
	StateDirError     string
	PolicyBlock       string
	PolicyDead        string
	PolicyError       string
	NoQueueSpace      string
}

// XFRMStatistics holds IPsec/XFRM subsystem statistics read from
// /proc/net/xfrm_stat.
type XFRMStatistics struct { // file: /proc/net/xfrm_stat
	Inbound            XFRMInboundStatistics
	Outbound           XFRMOutboundStatistics
	ForwardHeaderError string
	AcquireError       string
}

// =====================================================================================
// Configuration
// =====================================================================================

// OSRelease holds parsed key/value data from /etc/os-release.
type OSRelease struct { // file: /etc/os-release
	PrettyName        string
	Name              string
	VersionID         string
	Version           string
	VersionCodename   string
	ID                string
	IDLike            string
	HomeURL           string
	SupportURL        string
	BugReportURL      string
	PrivacyPolicyURL  string
	UbuntuCodename    string
	Logo              string
}

// Hostname holds the system hostname from /etc/hostname.
type Hostname struct { // file: /etc/hostname
	Name string
}

// HostsEntry represents a single host-to-address mapping.
type HostsEntry struct {
	IPAddress string
	Hostnames []string
}

// HostsFile holds all entries parsed from /etc/hosts.
type HostsFile struct { // file: /etc/hosts
	Entries []HostsEntry
}

// ResolvConf holds parsed DNS resolver configuration from
// /etc/resolv.conf.
type ResolvConf struct { // file: /etc/resolv.conf
	Nameservers []string
	Search      []string
	Domain      string
	Options     []string
}

// NsswitchEntry represents a single database-to-source mapping from
// /etc/nsswitch.conf.
type NsswitchEntry struct {
	Database string
	Sources  []string
}

// NsswitchConfiguration holds all entries parsed from /etc/nsswitch.conf.
type NsswitchConfiguration struct { // file: /etc/nsswitch.conf
	Entries []NsswitchEntry
}

// PasswdEntry represents a single user account row from /etc/passwd.
type PasswdEntry struct {
	Username      string
	Password      string
	UID           string
	GID           string
	Comment       string
	HomeDirectory string
	Shell         string
}

// PasswdFile holds all entries parsed from /etc/passwd.
type PasswdFile struct { // file: /etc/passwd
	Entries []PasswdEntry
}

// ShellsFile holds the list of valid login shells from /etc/shells.
type ShellsFile struct { // file: /etc/shells
	Paths []string
}

// ProtocolDefinition represents a single protocol entry from
// /etc/protocols, distinct from the runtime ProtocolStatistics reported by
// /proc/net/protocols.
type ProtocolDefinition struct {
	Name    string
	Number  string
	Aliases []string
}

// ProtocolRegistry holds all entries parsed from /etc/protocols.
type ProtocolRegistry struct { // file: /etc/protocols
	Protocols []ProtocolDefinition
}

// ServiceEntry represents a single service-to-port mapping from
// /etc/services.
type ServiceEntry struct {
	Name     string
	Port     string
	Protocol string
	Aliases  []string
}

// ServicesFile holds all entries parsed from /etc/services.
type ServicesFile struct { // file: /etc/services
	Entries []ServiceEntry
}