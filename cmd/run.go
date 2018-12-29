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
	if len(arguments) > 2 {
		command = &arguments[1]
	}
	var args []string
	if len(arguments) > 3 {
		args = arguments[2:]
	}

	schema := aliases.Schema{
		index,
		path.Base(index),
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
			cli.BoolFlag{Name: "detach, d", Usage: "Run container in background and print container ID"},
			cli.BoolFlag{Name: "sig-roxy", Usage: "Proxy received signals to the process"},
			cli.StringFlag{Name: "name", Usage: "Assign a name to the container"},
			cli.StringFlag{Name: "detach-keys", Usage: "Override the key sequence for detaching a container"},
			cli.StringFlag{Name: "platform", Usage: "Set platform if server is multi-platform capable"},
			cli.BoolFlag{Name: "disable-content-trust", Usage: "Skip image verification"},
			// see: https://github.com/docker/cli/blob/18.09/cli/command/container/opts.go
			// General purpose flags
			cli.StringSliceFlag{Name: "attach, a", Usage: "Attach to STDIN, STDOUT or STDERR"},
			cli.StringSliceFlag{Name: "device-cgroup-rule", Usage: "Add a rule to the cgroup allowed devices list"},
			cli.StringSliceFlag{Name: "device", Usage: "Add a host device to the container"},
			cli.StringSliceFlag{Name: "env, e", Usage: "Set environment variables"},
			cli.StringSliceFlag{Name: "env-file", Usage: "Read in a file of environment variables"},
			cli.StringFlag{Name: "entrypoint", Usage: "Overwrite the default ENTRYPOINT of the image"},
			cli.StringSliceFlag{Name: "group-add", Usage: "Add additional groups to join"},
			cli.StringFlag{Name: "hostname, h", Usage: "Container host name"},
			cli.StringFlag{Name: "domainname", Usage: "Container NIS domain name"},
			cli.BoolFlag{Name: "interactive, i", Usage: "Keep STDIN open even if not attached"},
			cli.StringSliceFlag{Name: "label, l", Usage: "Set meta data on a container"},
			cli.StringSliceFlag{Name: "label-file", Usage: "Read in a line delimited file of labels"},
			cli.BoolFlag{Name: "read-only", Usage: "Mount the container's root filesystem as read only"},
			cli.StringFlag{Name: "restart", Usage: "Restart policy to apply when a container exits"},
			cli.StringFlag{Name: "stop-signal", Usage: "Signal to stop a container"},
			cli.BoolFlag{Name: "stop-timeout", Usage: "Timeout (in seconds) to stop a container"},
			cli.StringSliceFlag{Name: "sysctl", Usage: "Sysctl options"},
			cli.BoolFlag{Name: "tty, t", Usage: "Allocate a pseudo-TTY"},
			cli.StringSliceFlag{Name: "ulimit", Usage: "Ulimit options"},
			cli.StringFlag{Name: "user, u", Usage: "Username or UID (format: <name|uid>[:<group|gid>])"},
			cli.StringFlag{Name: "workdir, w", Usage: "Working directory inside the container"},
			cli.BoolFlag{Name: "rm", Usage: "Automatically remove the container when it exits"},
			// Security
			cli.StringSliceFlag{Name: "cap-add", Usage: "Add Linux capabilities"},
			cli.StringSliceFlag{Name: "cap-drop", Usage: "Drop Linux capabilities"},
			cli.BoolFlag{Name: "privileged", Usage: "Give extended privileges to this container"},
			cli.StringSliceFlag{Name: "security-opt", Usage: "Security Options"},
			cli.StringFlag{Name: "userns", Usage: "User namespace to use"},
			// Network and port publishing flag
			cli.StringSliceFlag{Name: "add-host", Usage: "Add a custom host-to-IP mapping (host:ip)"},
			cli.StringSliceFlag{Name: "dns", Usage: "Set custom DNS servers"},
			cli.StringSliceFlag{Name: "dns-option, dns-opt", Usage: "Set DNS options"},
			cli.StringSliceFlag{Name: "dns-search", Usage: "Set custom DNS search domains"},
			cli.StringSliceFlag{Name: "expose", Usage: "Expose a port or a range of ports"},
			cli.StringFlag{Name: "ip", Usage: "IPv4 address (e.g., 172.30.100.104)"},
			cli.StringFlag{Name: "ip6", Usage: "IPv6 address (e.g., 2001:db8::33)"},
			cli.StringSliceFlag{Name: "link", Usage: "Add link to another container"},
			cli.StringSliceFlag{Name: "link-local-ip", Usage: "Container IPv4/IPv6 link-local addresses"},
			cli.StringFlag{Name: "mac-address", Usage: "Container MAC address (e.g., 92:d0:c6:0a:29:33)"},
			cli.StringSliceFlag{Name: "publish, p", Usage: "Publish a container's port(s) to the host"},
			cli.BoolFlag{Name: "publish-all, P", Usage: "Publish all exposed ports to random ports"},
			cli.StringFlag{Name: "network, net", Usage: "Connect a container to a network"},
			cli.StringSliceFlag{Name: "network-alias, net-alias", Usage: "Add network-scoped alias for the container"},
			// Logging and storage
			cli.StringFlag{Name: "log-driver", Usage: "Logging driver for the container"},
			cli.StringFlag{Name: "volume-driver", Usage: "Optional volume driver for the container"},
			cli.StringSliceFlag{Name: "log-opt", Usage: "Log driver options"},
			cli.StringSliceFlag{Name: "storage-opt", Usage: "Storage driver options for the container"},
			cli.StringSliceFlag{Name: "tmpfs", Usage: "Mount a tmpfs directory"},
			cli.StringSliceFlag{Name: "volumes-from", Usage: "Mount volumes from the specified container(s)"},
			cli.StringSliceFlag{Name: "volume, v", Usage: "Bind mount a volume"},
			cli.StringSliceFlag{Name: "mount", Usage: "Attach a filesystem mount to the container"},
			// Health-checking
			cli.StringFlag{Name: "health-cmd", Usage: "Command to run to check health"},
			cli.DurationFlag{Name: "health-interval", Usage: "Time between running the check (ms|s|m|h) (default 0s)"},
			cli.IntFlag{Name: "health-retries", Usage: "Consecutive failures needed to report unhealthy"},
			cli.DurationFlag{Name: "health-timeout", Usage: "Maximum time to allow one check to run (ms|s|m|h) (default 0s)"},
			cli.DurationFlag{Name: "health-start-period", Usage: "Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)"},
			cli.BoolFlag{Name: "no-healthcheck", Usage: "Disable any container-specified HEALTHCHECK"},
			// Resource management
			cli.UintFlag{Name: "blkio-weight", Usage: "Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)"},
			cli.BoolFlag{Name: "blkio-weight-device", Usage: "Block IO weight (relative device weight)"},
			cli.StringFlag{Name: "cidfile", Usage: "Write the container ID to the file"},
			cli.StringFlag{Name: "cpuset-cpus", Usage: "CPUs in which to allow execution (0-3, 0,1)"},
			cli.StringFlag{Name: "cpuset-mems", Usage: "MEMs in which to allow execution (0-3, 0,1)"},
			cli.StringFlag{Name: "cpu-period", Usage: "Limit CPU CFS (Completely Fair Scheduler) period"},
			cli.Int64Flag{Name: "cpu-quota", Usage: "Limit CPU CFS (Completely Fair Scheduler) quota"},
			cli.Int64Flag{Name: "cpu-rt-period", Usage: "Limit CPU real-time period in microseconds"},
			cli.Int64Flag{Name: "cpu-rt-runtime", Usage: "Limit CPU real-time runtime in microseconds"},
			cli.Int64Flag{Name: "cpu-shares, c", Usage: "CPU shares (relative weight)"},
			cli.StringFlag{Name: "cpus", Usage: "Number of CPUs"},
			cli.StringSliceFlag{Name: "device-read-bps", Usage: "Limit read rate (bytes per second) from a device"},
			cli.StringSliceFlag{Name: "device-read-iops", Usage: "Limit read rate (IO per second) from a device"},
			cli.StringSliceFlag{Name: "device-write-bps", Usage: "Limit write rate (bytes per second) to a device"},
			cli.StringSliceFlag{Name: "device-write-iops", Usage: "Limit write rate (IO per second) to a device"},
			cli.StringFlag{Name: "kernel-memory", Usage: "Kernel memory limit"},
			cli.StringFlag{Name: "memory", Usage: "Memory limit"},
			cli.StringFlag{Name: "memory-reservation", Usage: "Memory soft limit"},
			cli.StringFlag{Name: "memory-swap", Usage: "Swap limit equal to memory plus swap: '-1' to enable unlimited swap"},
			cli.Int64Flag{Name: "memory-swappiness", Usage: "Tune container memory swappiness (0 to 100)"},
			cli.BoolFlag{Name: "oom-kill-disable", Usage: "Disable OOM Killer"},
			cli.IntFlag{Name: "oom-score-adj", Usage: "Tune host's OOM preferences (-1000 to 1000)"},
			cli.Int64Flag{Name: "pids-limit", Usage: "Tune container pids limit (set -1 for unlimited)"},
			// Low-level execution (cgroups, namespaces, ...)
			cli.StringFlag{Name: "cgroup-parent", Usage: "Optional parent cgroup for the container"},
			cli.StringFlag{Name: "ipc", Usage: "IPC mode to use"},
			cli.StringFlag{Name: "isolation", Usage: "Container isolation technology"},
			cli.StringFlag{Name: "pid", Usage: "PID namespace to use"},
			cli.StringFlag{Name: "shm-size", Usage: "Size of /dev/shm"},
			cli.StringFlag{Name: "uts", Usage: "UTS namespace to use"},
			cli.StringFlag{Name: "runtime", Usage: "Runtime to use for this container"},
			cli.BoolFlag{Name: "init", Usage: "Run an init inside the container that forwards signals and reaps processes"},
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
	if err := ledger.Merge(index, *ctx.GetCommandShema()); err != nil {
		return err
	}

	schema, err := ledger.LookUp(index)
	if err != nil {
		return err
	}

	for _, dependency := range schema.Dependencies {
		s, err := ledger.LookUp(dependency)
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
