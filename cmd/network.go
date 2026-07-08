package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/abhijeetrajhansgithub/tsyx/internal/network"
)

var (
	branchNet string
	fileNet string
)

var netCmd = &cobra.Command{
	Use: "net",
	Short: "Display network information",
	RunE: func(cmd *cobra.Command, args []string) error {
		if runtime.GOOS != "linux" {
			return fmt.Errorf("memory is not yet supported on %s", runtime.GOOS)
		}

		// normalize branch
		switch branchNet {
		case "proc", "p":
			branchNet = "proc"
		case "sys", "s":
			branchNet = "sys"
		case "etc", "e":
			branchNet = "etc"
			
		}

		// validate branch
		switch branchNet {
		case "proc", "p", "sys", "s", "etc", "e":
			// valid
		default:
			return fmt.Errorf("invalid branch %q (expected: proc, sys, etc)", branchNet)
		}

		// validate file
		switch branchNet {
		case "proc", "p":
			switch fileNet {
			case "arp", "dev", "if_inet6", "ptype", "igmp", "igmp6",
				"route", "rt_cache", "fib_triestat",
				"tcp", "tcp6", "udp", "udp6",
				"unix", "packet", "netlink", "connector",
				"netstat", "snmp", "snmp6",
				"sockstat", "sockstat6",
				"protocols", "tls_stat", "xfrm_stat":
				// valid
			default:
				return fmt.Errorf("invalid proc file %q", fileNet)
			}

		case "sys", "s":
			switch fileNet {
			case "class_net":
				// valid
			default:
				return fmt.Errorf("invalid sys file %q", fileNet)
			}

		case "etc", "e":
			switch fileNet {
			case "os_release", "hostname", "hosts",
				"resolv_conf", "nsswitch", "passwd",
				"shells", "protocols", "services":
				// valid
			default:
				return fmt.Errorf("invalid etc file %q", fileNet)
			}
		}

		fmt.Printf("Branch: %s, File: %s\n", branchNet, fileNet)

		var b strings.Builder

		b.WriteString(branchNet)
		b.WriteString("_")
		b.WriteString(fileNet)

		key := b.String()
		fmt.Printf("key: %s\n", key)	

		switch key {
		case "proc_arp":
			arpTableInfo, err := net.LinuxCollectArp(key)
			if err != nil {
				return err
			}

			net.FormatARP(arpTableInfo)
		case "proc_dev":
			devStructInfo, err := net.LinuxCollectDev(key)
			if err != nil {
				return err
			}

			net.FormatDev(devStructInfo)
		case "proc_if_inet6":
			ifnetInfo, err := net.LinuxCollectIfInet6(key)
			if err != nil {
				return err
			}

			net.FormatIfNet6(ifnetInfo)
		} 


		return nil
	},
}

		

func init() {
	rootCmd.AddCommand(netCmd)

	netCmd.Flags().StringVarP(
		&branchNet,
		"branch",
		"b",
		"proc",
		`Branch:
  proc (p)
  sys  (s)
  etc  (e)`,
	)

	netCmd.Flags().StringVarP(
		&fileNet,
		"file",
		"f",
		"arp",
		`File:

proc:
  arp
  dev
  if_inet6
  ptype
  igmp
  igmp6
  route
  rt_cache
  fib_triestat
  tcp
  tcp6
  udp
  udp6
  unix
  packet
  netlink
  connector
  netstat
  snmp
  snmp6
  sockstat
  sockstat6
  protocols
  tls_stat
  xfrm_stat

sys:
  class_net

etc:
  os_release
  hostname
  hosts
  resolv_conf
  nsswitch
  passwd
  shells
  protocols
  services`,
	)
}