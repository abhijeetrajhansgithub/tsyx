# Running TSYX from a cloned repository

TSYX is a Go command-line application. Run these commands from the repository
root, where `go.mod` and `main.go` are located.

## Prerequisites

- Go 1.25 or later, as specified in `go.mod`.
- Linux for the system information commands currently implemented.

## Clone and run

```sh
git clone https://github.com/abhijeetrajhansgithub/tsyx.git
cd tsyx
go run . --help
```

`go run .` compiles and runs the module's main package. Go downloads required
modules on the first run if they are not already available locally. No binary
installation is required.

The general form is:

```sh
go run . <command> [flags]
```

For example, run the CPU, disk, memory, network, or uptime command:

```sh
go run . cpu
go run . disk
go run . mem
go run . net
go run . uptime
```

The network command defaults to the `proc` branch and `arp` file.

## Help

Use the root help to list commands, or request help for a specific command:

```sh
go run . --help
go run . -h
go run . help
go run . help mem
go run . net --help
```

## Disk

Choose the unit used to display disk sizes with `--unit` or `-u`:

```sh
go run . disk --unit auto
go run . disk --unit gb
go run . disk -u mb
```

Accepted units are `auto`, `b`, `kb`, `mb`, `gb`, and `tb`. Short forms are
also accepted for the prefixed units: `k`, `m`, `g`, and `t`.

## Memory

Choose a view with `--view` or `-v`, and a display unit with `--unit` or `-u`:

```sh
go run . mem
go run . mem --view summary --unit auto
go run . mem -v detailed -u mb
go run . mem --view kernel --unit gb
```

Views are `summary`, `detailed`, and `kernel`; their short forms are `s`, `d`,
and `k`. Units are `auto`, `b`, `kb`, `mb`, `gb`, and `tb`, with `k`, `m`, `g`,
and `t` accepted as short forms.

## Network

Select a branch with `--branch` or `-b`, and a file with `--file` or `-f`:

```sh
go run . net --branch proc --file arp
go run . net -b p -f dev
go run . net --branch proc --file if_inet6
go run . net --branch etc --file shells
go run . net -b e -f services
go run . net --branch sys --file class_net
```

Branch names and aliases are `proc`/`p`, `sys`/`s`, and `etc`/`e`.

The CLI currently accepts these file names:

- `proc`: `arp`, `dev`, `if_inet6`, `ptype`, `igmp`, `igmp6`, `route`,
  `rt_cache`, `fib_triestat`, `tcp`, `tcp6`, `udp`, `udp6`, `unix`, `packet`,
  `netlink`, `connector`, `netstat`, `snmp`, `snmp6`, `sockstat`, `sockstat6`,
  `protocols`, `tls_stat`, `xfrm_stat`
- `sys`: `class_net`
- `etc`: `os_release`, `hostname`, `hosts`, `resolv_conf`, `nsswitch`,
  `passwd`, `shells`, `protocols`, `services`

Of these, data collection and formatting are currently wired for `proc` files
`arp`, `dev`, `if_inet6`, `igmp`, `igmp6`, and `ptype`, and `etc` files
`passwd`, `shells`, `protocols`, and `services`. The other listed names pass
input validation but do not yet have a collection/formatting handler, so they
may only print the selected branch and file key.

## Platform support

The current CPU, disk, memory, network, and uptime command implementations are
Linux-only. Running them on another operating system returns an unsupported-OS
error.