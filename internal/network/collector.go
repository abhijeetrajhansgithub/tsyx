package net

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

func LinuxCollect() {
	
}