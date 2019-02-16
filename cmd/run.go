package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/k-kinzal/aliases/pkg/logger"

	"github.com/k-kinzal/aliases/pkg/aliases/context"

	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/aliases/script"

	"github.com/k-kinzal/aliases/pkg/aliases/config"

	"github.com/urfave/cli"
)

// opt is docker run option for aliases run.
var opt = docker.RunOption{}

// boolFlag returns string from boolean flag.
func boolFlag(c *cli.Context, name string) *string {
	for _, n := range strings.Split(name, ",") {
		n = strings.Trim(n, " \t")
		if c.IsSet(n) {
			val := fmt.Sprintf("%t", c.Bool(n))
			return &val
		}
	}
	return nil
}

// stringFlag returns string from string flag.
func stringFlag(c *cli.Context, name string) *string {
	for _, n := range strings.Split(name, ",") {
		n = strings.Trim(n, " \t")
		if c.IsSet(n) {
			val := c.String(n)
			return &val
		}
	}
	return nil
}

// sliceFlag returns []string from string slice flag.
func sliceFlag(c *cli.Context, name string) []string {
	for _, n := range strings.Split(name, ",") {
		n = strings.Trim(n, " \t")
		if c.IsSet(n) {
			return c.StringSlice(n)
		}
	}
	return nil
}

// mapFlag returns map[string]string from string slice flag.
func mapFlag(c *cli.Context, name string) map[string]string {
	for _, n := range strings.Split(name, ",") {
		n = strings.Trim(n, " \t")
		if c.IsSet(n) {
			val := map[string]string{}
			for _, v := range c.StringSlice(n) {
				s := strings.Split(v, "=")
				val[s[0]] = strings.Join(s[1:], "=")
			}
			return val
		}
	}
	return nil
}

// RunCommand returns `aliases run` command.
//
// `aliases run` is a command for run the command defined in aliases.yaml.
// By specifying the same option as `docker run`, overwrite the command defined in aliases.yaml and run it.
func RunCommand() cli.Command {
	return cli.Command{
		Name:      "run",
		Usage:     "Run aliases command",
		UsageText: "aliases [global options] run [command options] command [arguments...]",
		Flags: []cli.Flag{
			// see: https://github.com/docker/cli/blob/18.09/cli/command/container/run.go
			cli.BoolFlag{Name: "detach, d", Usage: "run container in background and print container ID"},
			cli.BoolFlag{Name: "sig-roxy", Usage: "proxy received signals to the process"},
			cli.StringFlag{Name: "name", Usage: "assign a name to the container"},
			cli.StringFlag{Name: "detach-keys", Usage: "override the key sequence for detaching a container"},
			cli.StringFlag{Name: "platform", Usage: "set platform if server is multi-platform capable"},
			cli.BoolFlag{Name: "disable-content-trust", Usage: "skip image verification"},
			// see: https://github.com/docker/cli/blob/18.09/cli/command/container/opts.go
			// General purpose flags
			cli.StringSliceFlag{Name: "attach, a", Usage: "attach to STDIN, STDOUT or STDERR"},
			cli.StringSliceFlag{Name: "device-cgroup-rule", Usage: "add a rule to the cgroup allowed devices list"},
			cli.StringSliceFlag{Name: "device", Usage: "add a host device to the container"},
			cli.StringSliceFlag{Name: "env, e", Usage: "set environment variables"},
			cli.StringSliceFlag{Name: "env-file", Usage: "read in a file of environment variables"},
			cli.StringFlag{Name: "entrypoint", Usage: "overwrite the default ENTRYPOINT of the image"},
			cli.StringSliceFlag{Name: "group-add", Usage: "add additional groups to join"},
			cli.StringFlag{Name: "hostname, h", Usage: "container host name"},
			cli.StringFlag{Name: "domainname", Usage: "container NIS domain name"},
			cli.BoolFlag{Name: "interactive, i", Usage: "keep STDIN open even if not attached"},
			cli.StringSliceFlag{Name: "label, l", Usage: "set meta data on a container"},
			cli.StringSliceFlag{Name: "label-file", Usage: "read in a line delimited file of labels"},
			cli.BoolFlag{Name: "read-only", Usage: "mount the container's root filesystem as read only"},
			cli.StringFlag{Name: "restart", Usage: "restart policy to apply when a container exits"},
			cli.StringFlag{Name: "stop-signal", Usage: "signal to stop a container"},
			cli.BoolFlag{Name: "stop-timeout", Usage: "timeout (in seconds) to stop a container"},
			cli.StringSliceFlag{Name: "sysctl", Usage: "sysctl options"},
			cli.BoolFlag{Name: "tty, t", Usage: "allocate a pseudo-TTY"},
			cli.StringSliceFlag{Name: "ulimit", Usage: "ulimit options"},
			cli.StringFlag{Name: "user, u", Usage: "username or UID (format: <name|uid>[:<group|gid>])"},
			cli.StringFlag{Name: "workdir, w", Usage: "working directory inside the container"},
			cli.BoolFlag{Name: "rm", Usage: "automatically remove the container when it exits"},
			// Security
			cli.StringSliceFlag{Name: "cap-add", Usage: "add Linux capabilities"},
			cli.StringSliceFlag{Name: "cap-drop", Usage: "drop Linux capabilities"},
			cli.BoolFlag{Name: "privileged", Usage: "give extended privileges to this container"},
			cli.StringSliceFlag{Name: "security-opt", Usage: "security Options"},
			cli.StringFlag{Name: "userns", Usage: "user namespace to use"},
			// Network and port publishing flag
			cli.StringSliceFlag{Name: "add-host", Usage: "add a custom host-to-IP mapping (host:ip)"},
			cli.StringSliceFlag{Name: "dns", Usage: "set custom DNS servers"},
			cli.StringSliceFlag{Name: "dns-option, dns-opt", Usage: "set DNS options"},
			cli.StringSliceFlag{Name: "dns-search", Usage: "set custom DNS search domains"},
			cli.StringSliceFlag{Name: "expose", Usage: "expose a port or a range of ports"},
			cli.StringFlag{Name: "ip", Usage: "IPv4 address (e.g., 172.30.100.104)"},
			cli.StringFlag{Name: "ip6", Usage: "IPv6 address (e.g., 2001:db8::33)"},
			cli.StringSliceFlag{Name: "link", Usage: "add link to another container"},
			cli.StringSliceFlag{Name: "link-local-ip", Usage: "container IPv4/IPv6 link-local addresses"},
			cli.StringFlag{Name: "mac-address", Usage: "container MAC address (e.g., 92:d0:c6:0a:29:33)"},
			cli.StringSliceFlag{Name: "publish, p", Usage: "publish a container's port(s) to the host"},
			cli.BoolFlag{Name: "publish-all, P", Usage: "publish all exposed ports to random ports"},
			cli.StringFlag{Name: "network, net", Usage: "connect a container to a network"},
			cli.StringSliceFlag{Name: "network-alias, net-alias", Usage: "add network-scoped alias for the container"},
			// Logging and storage
			cli.StringFlag{Name: "log-driver", Usage: "logging driver for the container"},
			cli.StringFlag{Name: "volume-driver", Usage: "optional volume driver for the container"},
			cli.StringSliceFlag{Name: "log-opt", Usage: "log driver options"},
			cli.StringSliceFlag{Name: "storage-opt", Usage: "storage driver options for the container"},
			cli.StringSliceFlag{Name: "tmpfs", Usage: "mount a tmpfs directory"},
			cli.StringSliceFlag{Name: "volumes-from", Usage: "mount volumes from the specified container(s)"},
			cli.StringSliceFlag{Name: "volume, v", Usage: "bind mount a volume"},
			cli.StringSliceFlag{Name: "mount", Usage: "attach a filesystem mount to the container"},
			// Health-checking
			cli.StringFlag{Name: "health-cmd", Usage: "command to run to check health"},
			cli.StringFlag{Name: "health-interval", Usage: "time between running the check (ms|s|m|h) (default 0s)"},
			cli.StringFlag{Name: "health-retries", Usage: "consecutive failures needed to report unhealthy"},
			cli.StringFlag{Name: "health-timeout", Usage: "maximum time to allow one check to run (ms|s|m|h) (default 0s)"},
			cli.StringFlag{Name: "health-start-period", Usage: "start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)"},
			cli.BoolFlag{Name: "no-healthcheck", Usage: "disable any container-specified HEALTHCHECK"},
			// Resource management
			cli.StringFlag{Name: "blkio-weight", Usage: "block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)"},
			cli.BoolFlag{Name: "blkio-weight-device", Usage: "block IO weight (relative device weight)"},
			cli.StringFlag{Name: "cidfile", Usage: "write the container ID to the file"},
			cli.StringFlag{Name: "cpuset-cpus", Usage: "CPUs in which to allow execution (0-3, 0,1)"},
			cli.StringFlag{Name: "cpuset-mems", Usage: "MEMs in which to allow execution (0-3, 0,1)"},
			cli.StringFlag{Name: "cpu-period", Usage: "limit CPU CFS (Completely Fair Scheduler) period"},
			cli.StringFlag{Name: "cpu-quota", Usage: "limit CPU CFS (Completely Fair Scheduler) quota"},
			cli.StringFlag{Name: "cpu-rt-period", Usage: "limit CPU real-time period in microseconds"},
			cli.StringFlag{Name: "cpu-rt-runtime", Usage: "limit CPU real-time runtime in microseconds"},
			cli.StringFlag{Name: "cpu-shares, c", Usage: "CPU shares (relative weight)"},
			cli.StringFlag{Name: "cpus", Usage: "number of CPUs"},
			cli.StringSliceFlag{Name: "device-read-bps", Usage: "limit read rate (bytes per second) from a device"},
			cli.StringSliceFlag{Name: "device-read-iops", Usage: "limit read rate (IO per second) from a device"},
			cli.StringSliceFlag{Name: "device-write-bps", Usage: "limit write rate (bytes per second) to a device"},
			cli.StringSliceFlag{Name: "device-write-iops", Usage: "limit write rate (IO per second) to a device"},
			cli.StringFlag{Name: "kernel-memory", Usage: "kernel memory limit"},
			cli.StringFlag{Name: "memory", Usage: "memory limit"},
			cli.StringFlag{Name: "memory-reservation", Usage: "memory soft limit"},
			cli.StringFlag{Name: "memory-swap", Usage: "swap limit equal to memory plus swap: '-1' to enable unlimited swap"},
			cli.StringFlag{Name: "memory-swappiness", Usage: "tune container memory swappiness (0 to 100)"},
			cli.BoolFlag{Name: "oom-kill-disable", Usage: "disable OOM Killer"},
			cli.StringFlag{Name: "oom-score-adj", Usage: "tune host's OOM preferences (-1000 to 1000)"},
			cli.StringFlag{Name: "pids-limit", Usage: "tune container pids limit (set -1 for unlimited)"},
			// Low-level execution (cgroups, namespaces, ...)
			cli.StringFlag{Name: "cgroup-parent", Usage: "optional parent cgroup for the container"},
			cli.StringFlag{Name: "ipc", Usage: "IPC mode to use"},
			cli.StringFlag{Name: "isolation", Usage: "container isolation technology"},
			cli.StringFlag{Name: "pid", Usage: "PID namespace to use"},
			cli.StringFlag{Name: "shm-size", Usage: "size of /dev/shm"},
			cli.StringFlag{Name: "uts", Usage: "UTS namespace to use"},
			cli.StringFlag{Name: "runtime", Usage: "runtime to use for this container"},
			cli.BoolFlag{Name: "init", Usage: "run an init inside the container that forwards signals and reaps processes"},
		},
		Before: func(c *cli.Context) error {
			// FIXME: ٩(ˊᗜˋ*)و.
			// since the type missmatch and destination is not supported, mapping is done with cli.Command.Before.
			opt.Detach = boolFlag(c, "detach, d")
			opt.SigProxy = boolFlag(c, "sig-roxy")
			opt.Name = stringFlag(c, "name")
			opt.DetachKeys = stringFlag(c, "detach-keys")
			opt.Platform = stringFlag(c, "platform")
			opt.DisableContentTrust = boolFlag(c, "disable-content-trust")
			opt.Attach = sliceFlag(c, "attach, a")
			opt.DeviceCgroupRule = sliceFlag(c, "device-cgroup-rule")
			opt.Device = sliceFlag(c, "device")
			opt.Env = mapFlag(c, "env, e")
			opt.EnvFile = sliceFlag(c, "env-file")
			opt.Entrypoint = stringFlag(c, "entrypoint")
			opt.GroupAdd = sliceFlag(c, "group-add")
			opt.Hostname = stringFlag(c, "hostname, h")
			opt.Domainname = stringFlag(c, "domainname")
			opt.Interactive = boolFlag(c, "interactive, i")
			opt.Label = mapFlag(c, "label, l")
			opt.LabelFile = sliceFlag(c, "label-file")
			opt.ReadOnly = boolFlag(c, "read-only")
			opt.Restart = stringFlag(c, "restart")
			opt.StopSignal = stringFlag(c, "stop-signal")
			opt.StopTimeout = boolFlag(c, "stop-timeout")
			opt.Sysctl = mapFlag(c, "sysctl")
			opt.TTY = boolFlag(c, "tty, t")
			opt.Ulimit = mapFlag(c, "ulimit")
			opt.User = stringFlag(c, "user, u")
			opt.Workdir = stringFlag(c, "workdir, w")
			opt.Rm = boolFlag(c, "rm")
			opt.CapAdd = sliceFlag(c, "cap-add")
			opt.CapDrop = sliceFlag(c, "cap-drop")
			opt.Privileged = boolFlag(c, "privileged")
			opt.SecurityOpt = mapFlag(c, "security-opt")
			opt.Userns = stringFlag(c, "userns")
			opt.AddHost = sliceFlag(c, "add-host")
			opt.DNS = sliceFlag(c, "dns")
			opt.DNSOption = sliceFlag(c, "dns-option, dns-opt")
			opt.DNSSearch = sliceFlag(c, "dns-search")
			opt.Expose = sliceFlag(c, "expose")
			opt.IP = stringFlag(c, "ip")
			opt.IP6 = stringFlag(c, "ip6")
			opt.Link = sliceFlag(c, "link")
			opt.LinkLocalIP = sliceFlag(c, "link-local-ip")
			opt.MacAddress = stringFlag(c, "mac-address")
			opt.Publish = sliceFlag(c, "publish, p")
			opt.PublishAll = boolFlag(c, "publish-all, P")
			opt.Network = stringFlag(c, "network, net")
			opt.NetworkAlias = sliceFlag(c, "network-alias, net-alias")
			opt.LogDriver = stringFlag(c, "log-driver")
			opt.VolumeDriver = stringFlag(c, "volume-driver")
			opt.LogOpt = mapFlag(c, "log-opt")
			opt.StorageOpt = mapFlag(c, "storage-opt")
			opt.Tmpfs = sliceFlag(c, "tmpfs")
			opt.VolumesFrom = sliceFlag(c, "volumes-from")
			opt.Volume = sliceFlag(c, "volume, v")
			opt.Mount = mapFlag(c, "mount")
			opt.HealthCmd = stringFlag(c, "health-cmd")
			opt.HealthInterval = stringFlag(c, "health-interval")
			opt.HealthRetries = stringFlag(c, "health-retries")
			opt.HealthTimeout = stringFlag(c, "health-timeout")
			opt.HealthStartPeriod = stringFlag(c, "health-start-period")
			opt.NoHealthcheck = boolFlag(c, "no-healthcheck")
			opt.BlkioWeight = stringFlag(c, "blkio-weight")
			opt.BlkioWeight = boolFlag(c, "blkio-weight-device")
			opt.CIDFile = stringFlag(c, "cidfile")
			opt.CPUsetCPUs = stringFlag(c, "cpuset-cpus")
			opt.CPUsetMems = stringFlag(c, "cpuset-mems")
			opt.CPUPeriod = stringFlag(c, "cpu-period")
			opt.CPUQuota = stringFlag(c, "cpu-quota")
			opt.CPURtPeriod = stringFlag(c, "cpu-rt-period")
			opt.CPURtRuntime = stringFlag(c, "cpu-rt-runtime")
			opt.CPUShares = stringFlag(c, "cpu-shares, c")
			opt.CPUs = stringFlag(c, "cpus")
			opt.DeviceReadBPS = sliceFlag(c, "device-read-bps")
			opt.DeviceReadIOPS = sliceFlag(c, "device-read-iops")
			opt.DeviceWriteBPS = sliceFlag(c, "device-write-bps")
			opt.DeviceWriteIOPS = sliceFlag(c, "device-write-iops")
			opt.KernelMemory = stringFlag(c, "kernel-memory")
			opt.Memory = stringFlag(c, "memory")
			opt.MemoryReservation = stringFlag(c, "memory-reservation")
			opt.MemorySwap = stringFlag(c, "memory-swap")
			opt.MemorySwappiness = stringFlag(c, "memory-swappiness")
			opt.OOMKillDisable = boolFlag(c, "oom-kill-disable")
			opt.OOMScoreAdj = stringFlag(c, "oom-score-adj")
			opt.PidsLimit = stringFlag(c, "pids-limit")
			opt.CgroupParent = stringFlag(c, "cgroup-parent")
			opt.IPC = stringFlag(c, "ipc")
			opt.Isolation = stringFlag(c, "isolation")
			opt.PID = stringFlag(c, "pid")
			opt.ShmSize = stringFlag(c, "shm-size")
			opt.UTS = stringFlag(c, "uts")
			opt.Runtime = stringFlag(c, "runtime")
			opt.Init = boolFlag(c, "init")
			return nil
		},
		Action:                 RunAction,
		SkipArgReorder:         true,
		UseShortOptionHandling: true,
	}
}

// RunAction is the action of `aliases run`.
func RunAction(c *cli.Context) error {
	index := c.Args().Get(0)
	if index == "" {
		return cli.ShowCommandHelp(c, "run")
	}
	args := c.Args()
	if c.NArg() > 1 {
		args = args[1:]
	}

	dir, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		return err
	}
	if err := context.ChangeExportPath(dir); err != nil {
		return err
	}

	client, err := docker.NewClient()
	if err != nil {
		return err
	}

	conf, err := config.LoadConfig(context.ConfPath())
	if err != nil {
		return err
	}

	option, err := conf.Get(index)
	if err != nil {
		return err
	}

	cmd := script.NewScript(client, *option)
	logger.Debug(cmd.StringWithOverride(args, opt))

	if err := cmd.Run(args, opt); err != nil {
		return err
	}

	return nil
}
