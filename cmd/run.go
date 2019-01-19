package cmd

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/k-kinzal/aliases/pkg/posix"

	"github.com/k-kinzal/aliases/pkg/aliases"
	"github.com/k-kinzal/aliases/pkg/export"
	"github.com/k-kinzal/aliases/pkg/logger"
	"github.com/urfave/cli"
)

type runContext struct {
	aliases.Context
	flags      helper
	exportPath string
}

func (ctx *runContext) ExportPath() string {
	return ctx.exportPath
}

func (ctx *runContext) GetCommandShema() *aliases.Schema {
	flags := ctx.flags
	arguments := flags.Args()

	if len(arguments) == 0 {
		return nil
	}
	index := arguments[0]

	var command *string
	if len(arguments) > 1 {
		command = &arguments[1]
	}
	var args []string
	if len(arguments) > 2 {
		args = arguments[2:]
	}

	schema := aliases.Schema{
		index,
		path.Base(index),
		aliases.BinarySchema{"", ""},
		nil,
		flags.bool("detach", "d"),
		flags.bool("sig-proxy"),
		flags.string("name"),
		flags.string("detach-keys"),
		flags.string("platform"),
		flags.bool("disable-content-trust"),
		flags.stringSlice("attach", "a"),
		flags.stringSlice("device-cgroup-rule"),
		flags.stringSlice("device"),
		flags.stringMap("env", "e"),
		flags.stringSlice("env-file"),
		flags.string("entrypoint"),
		flags.stringSlice("group-add"),
		flags.string("hostname"),
		flags.string("domainname"),
		flags.bool("interactive", "i"),
		flags.stringMap("label", "l"),
		flags.stringSlice("label-file"),
		flags.bool("read-only"),
		flags.string("restart"),
		flags.string("stop-signal"),
		flags.int("stop-timeout"),
		flags.stringMap("sysctl"),
		flags.bool("tty", "t"),
		flags.stringMap("ulimit"),
		flags.string("user", "u"),
		flags.string("work-dir", "w"),
		flags.bool("rm"),
		flags.stringSlice("cap-add"),
		flags.stringSlice("cap-drop"),
		flags.string("privileged"),
		flags.stringMap("security-opt"),
		flags.string("userns"),
		flags.stringSlice("add-host"),
		flags.stringSlice("dns"),
		flags.stringSlice("dns-option", "dns-opt"),
		flags.stringSlice("dns-search"),
		flags.stringSlice("expose"),
		flags.string("ip"),
		flags.string("ip6"),
		flags.stringSlice("link"),
		flags.stringSlice("link-local-ip"),
		flags.string("mac-address"),
		flags.stringSlice("publish", "p"),
		flags.bool("publish-all", "P"),
		flags.string("network", "net"),
		flags.stringSlice("network-alias", "net-alias"),
		flags.string("log-driver"),
		flags.string("volume-driver"),
		flags.stringMap("log-opt"),
		flags.stringMap("storage-opt"),
		flags.stringSlice("tmpfs"),
		flags.stringSlice("volumes-from"),
		flags.stringSlice("volume", "v"),
		flags.stringMap("mount"),
		flags.string("health-cmd"),
		flags.string("health-interval"),
		flags.int("health-retries"),
		flags.string("health-timeout"),
		flags.string("health-start-period"),
		flags.bool("no-healthcheck"),
		flags.uint16("blkio-weight"),
		flags.stringSlice("blkio-weight-device"),
		flags.string("cidFile"),
		flags.string("cpuset-cpus"),
		flags.string("cpuset-mems"),
		flags.string("cpu-period"),
		flags.int64("cpu-quota"),
		flags.int64("cpu-rt-period"),
		flags.int64("cpu-rt-runtime"),
		flags.int64("cpu-shares", "c"),
		flags.string("cpus"),
		flags.stringSlice("device-read-bps"),
		flags.stringSlice("device-read-iops"),
		flags.stringSlice("device-write-bps"),
		flags.stringSlice("device-write-iops"),
		flags.string("kernel-memory"),
		flags.string("memory"),
		flags.string("memory-reservation"),
		flags.string("memory-swap"),
		flags.int64("memory-swappiness"),
		flags.bool("oom-kill-disable"),
		flags.int("oom-score-adj"),
		flags.int64("pids-limit"),
		flags.string("cgroup-parent"),
		flags.string("ipc"),
		flags.string("isolation"),
		flags.string("pid"),
		flags.string("shm-size"),
		flags.string("uts"),
		flags.string("runtime"),
		flags.bool("init"),
		"",
		args,
		"",
		command,
	}

	return &schema
}

func newRunContext(c *cli.Context) (*runContext, error) {
	ctx, err := aliases.NewContext(
		c.GlobalString("home"),
		c.GlobalString("config"),
	)
	if err != nil {
		return nil, err
	}
	dir, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	return &runContext{ctx, helper{c}, dir}, nil
}

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
			cli.DurationFlag{Name: "health-interval", Usage: "time between running the check (ms|s|m|h) (default 0s)"},
			cli.IntFlag{Name: "health-retries", Usage: "consecutive failures needed to report unhealthy"},
			cli.DurationFlag{Name: "health-timeout", Usage: "maximum time to allow one check to run (ms|s|m|h) (default 0s)"},
			cli.DurationFlag{Name: "health-start-period", Usage: "start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)"},
			cli.BoolFlag{Name: "no-healthcheck", Usage: "disable any container-specified HEALTHCHECK"},
			// Resource management
			cli.UintFlag{Name: "blkio-weight", Usage: "block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)"},
			cli.BoolFlag{Name: "blkio-weight-device", Usage: "block IO weight (relative device weight)"},
			cli.StringFlag{Name: "cidfile", Usage: "write the container ID to the file"},
			cli.StringFlag{Name: "cpuset-cpus", Usage: "CPUs in which to allow execution (0-3, 0,1)"},
			cli.StringFlag{Name: "cpuset-mems", Usage: "MEMs in which to allow execution (0-3, 0,1)"},
			cli.StringFlag{Name: "cpu-period", Usage: "limit CPU CFS (Completely Fair Scheduler) period"},
			cli.Int64Flag{Name: "cpu-quota", Usage: "limit CPU CFS (Completely Fair Scheduler) quota"},
			cli.Int64Flag{Name: "cpu-rt-period", Usage: "limit CPU real-time period in microseconds"},
			cli.Int64Flag{Name: "cpu-rt-runtime", Usage: "limit CPU real-time runtime in microseconds"},
			cli.Int64Flag{Name: "cpu-shares, c", Usage: "CPU shares (relative weight)"},
			cli.StringFlag{Name: "cpus", Usage: "number of CPUs"},
			cli.StringSliceFlag{Name: "device-read-bps", Usage: "limit read rate (bytes per second) from a device"},
			cli.StringSliceFlag{Name: "device-read-iops", Usage: "limit read rate (IO per second) from a device"},
			cli.StringSliceFlag{Name: "device-write-bps", Usage: "limit write rate (bytes per second) to a device"},
			cli.StringSliceFlag{Name: "device-write-iops", Usage: "limit write rate (IO per second) to a device"},
			cli.StringFlag{Name: "kernel-memory", Usage: "kernel memory limit"},
			cli.StringFlag{Name: "memory", Usage: "memory limit"},
			cli.StringFlag{Name: "memory-reservation", Usage: "memory soft limit"},
			cli.StringFlag{Name: "memory-swap", Usage: "swap limit equal to memory plus swap: '-1' to enable unlimited swap"},
			cli.Int64Flag{Name: "memory-swappiness", Usage: "tune container memory swappiness (0 to 100)"},
			cli.BoolFlag{Name: "oom-kill-disable", Usage: "disable OOM Killer"},
			cli.IntFlag{Name: "oom-score-adj", Usage: "tune host's OOM preferences (-1000 to 1000)"},
			cli.Int64Flag{Name: "pids-limit", Usage: "tune container pids limit (set -1 for unlimited)"},
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
		Action:                 RunAction,
		SkipArgReorder:         true,
		UseShortOptionHandling: true,
	}
}

func RunAction(c *cli.Context) error {
	if c.NArg() == 0 {
		return cli.ShowCommandHelp(c, "run")
	}

	ctx, err := newRunContext(c)
	if err != nil {
		return err
	}

	if err := ctx.MakeExportDir(); err != nil {
		return err
	}

	ledger, err := aliases.NewLedgerFromConfig(ctx.ConfPath())
	if err != nil {
		return err
	}

	index := c.Args()[0]

	src, err := ledger.LookUp(index)
	if err != nil {
		return err
	}
	dst := ctx.GetCommandShema()
	dst.Dependencies = src.Dependencies
	dst.Image = src.Image
	dst.Tag = src.Tag

	if err := ledger.Merge(index, *dst); err != nil {
		return err
	}

	schema, err := ledger.LookUp(index)
	if err != nil {
		return err
	}

	for _, dependency := range schema.Dependencies {
		if dependency.IsSchema() {
			for _, s := range dependency.Schemas() {
				cmd, err := aliases.NewCommand(ctx, s)
				if err != nil {
					return err
				}
				if err := export.Script(path.Join(ctx.ExportPath(), s.FileName), *cmd); err != nil {
					return err
				}
			}
		} else {
			s, err := ledger.LookUp(dependency.String())
			if err != nil {
				return err
			}
			cmd, err := aliases.NewCommand(ctx, *s)
			if err != nil {
				return err
			}
			if err := export.Script(path.Join(ctx.ExportPath(), s.FileName), *cmd); err != nil {
				return err
			}
		}
	}
	cmd, err := aliases.NewCommand(ctx, *schema)
	if err != nil {
		return err
	}

	logger.Debug(cmd)

	if err := posix.Shell(cmd.String()).Run(); err != nil {
		return err
	}

	return nil
}
