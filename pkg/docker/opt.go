package docker

import (
	"fmt"
	"strconv"
	"time"
)

type RunOpts struct {
	// see: https://github.com/docker/cli/blob/18.09/cli/command/container/run.go
	Detach     *bool   // Run container in background and print container ID
	SigProxy   *bool   // Proxy received signals to the process
	Name       *string // Assign a name to the container
	DetachKeys *string // Override the key sequence for detaching a container

	Platform            *string // Set platform if server is multi-platform capable
	DisableContentTrust *bool   // Skip image verification

	// see: https://github.com/docker/cli/blob/18.09/cli/command/container/opts.go
	// General purpose flags
	Attach           []string          // Attach to STDIN, STDOUT or STDERR
	DeviceCgroupRule []string          // Add a rule to the cgroup allowed devices list
	Device           []string          // Add a host device to the container
	Env              map[string]string // Set environment variables
	EnvFile          []string          // Read in a file of environment variables
	Entrypoint       *string           // Overwrite the default ENTRYPOINT of the image
	GroupAdd         []string          // Add additional groups to join
	Hostname         *string           // Container host name
	Domainname       *string           // Container NIS domain name
	Interactive      *bool             // Keep STDIN open even if not attached
	Label            map[string]string // Set meta data on a container
	LabelFile        []string          // Read in a line delimited file of labels
	ReadOnly         *bool             // Mount the container's root filesystem as read only
	Restart          *string           // Restart policy to apply when a container exits
	StopSignal       *string           // Signal to stop a container
	StopTimeout      *int              // Timeout (in seconds) to stop a container
	Sysctl           map[string]string // Sysctl options
	Tty              *bool             // Allocate a pseudo-TTY
	Ulimit           map[string]string // Ulimit options
	User             *string           // Username or UID (format: <name|uid>[:<group|gid>])
	Workdir          *string           // Working directory inside the container
	Rm               *bool             // Automatically remove the container when it exits

	// Security
	CapAdd      []string          // Add Linux capabilities
	CapDrop     []string          // Drop Linux capabilities
	Privileged  *bool             // Give extended privileges to this container
	SecurityOpt map[string]string // Security Options
	Userns      *string           // User namespace to use

	// Network and port publishing flag
	AddHost      []string // Add a custom host-to-IP mapping (host:ip)
	Dns          []string // Set custom DNS servers
	DnsOpt       []string // Set DNS options
	DnsOption    []string // Set DNS options
	DnsSearch    []string // Set custom DNS search domains
	Expose       []string // Expose a port or a range of ports
	Ip           *string  // IPv4 address (e.g., 172.30.100.104)
	Ip6          *string  // IPv6 address (e.g., 2001:db8::33)
	Link         []string // Add link to another container
	LinkLocalIp  []string // Container IPv4/IPv6 link-local addresses
	MacAddress   *string  // Container MAC address (e.g., 92:d0:c6:0a:29:33)
	Publish      []string // Publish a container's port(s) to the host
	PublishAll   *bool    // Publish all exposed ports to random ports
	Net          *string  // Connect a container to a network
	Network      *string  // Connect a container to a network
	NetAlias     []string // Add network-scoped alias for the container
	NetworkAlias []string // Add network-scoped alias for the container

	// Logging and storage
	LogDriver    *string           // Logging driver for the container
	VolumeDriver *string           // Optional volume driver for the container
	LogOpt       map[string]string // Log driver options
	StorageOpt   map[string]string // Storage driver options for the container
	Tmpfs        []string          // Mount a tmpfs directory
	VolumesFrom  []string          // Mount volumes from the specified container(s)
	Volume       []string          // Bind mount a volume
	Mount        map[string]string // Attach a filesystem mount to the container

	// Health-checking
	HealthCmd         *string        // Command to run to check health
	HealthInterval    *time.Duration // Time between running the check (ms|s|m|h) (default 0s)
	HealthRetries     *int           // Consecutive failures needed to report unhealthy
	HealthTimeout     *time.Duration // Maximum time to allow one check to run (ms|s|m|h) (default 0s)
	HealthStartPeriod *time.Duration // Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)
	NoHealthcheck     *bool          // Disable any container-specified HEALTHCHECK

	// Resource management
	BlkioWeight       *uint16  // Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)
	BlkioWeightDevice []string // Block IO weight (relative device weight)
	Cidfile           *string  // Write the container ID to the file
	CpusetCpus        *string  // CPUs in which to allow execution (0-3, 0,1)
	CpusetMems        *string  // MEMs in which to allow execution (0-3, 0,1)
	//CpuCount *int64            // CPU count (Windows only)
	//CpuPercent *int64          // CPU percent (Windows only)
	CpuPeriod       *int64   // Limit CPU CFS (Completely Fair Scheduler) period
	CpuQuota        *int64   // Limit CPU CFS (Completely Fair Scheduler) quota
	CpuRtPeriod     *int64   // Limit CPU real-time period in microseconds
	CpuRtRuntime    *int64   // Limit CPU real-time runtime in microseconds
	CpuShares       *int64   // CPU shares (relative weight)
	Cpus            *string  // Number of CPUs
	DeviceReadBps   []string // Limit read rate (bytes per second) from a device
	DeviceReadIops  []string // Limit read rate (IO per second) from a device
	DeviceWriteBps  []string // Limit write rate (bytes per second) to a device
	DeviceWriteIops []string // Limit write rate (IO per second) to a device
	//IoMaxbandwidth *string     // Maximum IO bandwidth limit for the system drive (Windows only)
	//IoMaxiops *uint64          // Maximum IOps limit for the system drive (Windows only)
	KernelMemory      *string // Kernel memory limit
	Memory            *string // Memory limit
	MemoryReservation *string // Memory soft limit
	MemorySwap        *string // Swap limit equal to memory plus swap: '-1' to enable unlimited swap
	MemorySwappiness  *int64  // Tune container memory swappiness (0 to 100)
	OomKillDisable    *bool   // Disable OOM Killer
	OomScoreAdj       *int    // Tune host's OOM preferences (-1000 to 1000)
	PidsLimit         *int64  // Tune container pids limit (set -1 for unlimited)

	// Low-level execution (cgroups, namespaces, ...)
	CgroupParent *string // Optional parent cgroup for the container
	Ipc          *string // IPC mode to use
	Isolation    *string // Container isolation technology
	Pid          *string // PID namespace to use
	ShmSize      *string // Size of /dev/shm
	Uts          *string // UTS namespace to use
	Runtime      *string // Runtime to use for this container

	Init *bool // Run an init inside the container that forwards signals and reaps processes

	Image string
	Args  []string
}

func (opt *RunOpts) toArguments() []string {
	args := make([]string, 0)
	for _, v := range opt.AddHost {
		args = append(args, "--add-host", strconv.Quote(v))
	}
	for _, v := range opt.Attach {
		args = append(args, "--attach", strconv.Quote(v))
	}
	if v := opt.BlkioWeight; v != nil {
		args = append(args, fmt.Sprintf("--blkio-weight=%d", *v))
	}
	for _, v := range opt.BlkioWeightDevice {
		args = append(args, "--blkio-weight-device", strconv.Quote(v))
	}
	for _, v := range opt.CapAdd {
		args = append(args, fmt.Sprintf("--cap-add=%s", strconv.Quote(v)))
	}
	for _, v := range opt.CapDrop {
		args = append(args, fmt.Sprintf("--cap-drop=%s", strconv.Quote(v)))
	}
	if v := opt.CgroupParent; v != nil {
		args = append(args, fmt.Sprintf("--cgroup-parent==%s", strconv.Quote(*v)))
	}
	if v := opt.Cidfile; v != nil {
		args = append(args, "--cidfile", strconv.Quote(*v))
	}
	if v := opt.CpuPeriod; v != nil {
		args = append(args, "--cpu-period", strconv.FormatInt(*v, 10))
	}
	if v := opt.CpuQuota; v != nil {
		args = append(args, "--cpu-quota", strconv.FormatInt(*v, 10))
	}
	if v := opt.CpuRtPeriod; v != nil {
		args = append(args, "--cpu-rt-period", strconv.FormatInt(*v, 10))
	}
	if v := opt.CpuRtRuntime; v != nil {
		args = append(args, "--cpu-rt-runtime", strconv.FormatInt(*v, 10))
	}
	if v := opt.CpuShares; v != nil {
		args = append(args, "--cpu-shares", strconv.FormatInt(*v, 10))
	}
	if v := opt.Cpus; v != nil {
		args = append(args, "--cpus", strconv.Quote(*v))
	}
	if v := opt.CpusetCpus; v != nil {
		args = append(args, "--cpuset-cpus", strconv.Quote(*v))
	}
	if v := opt.CpusetMems; v != nil {
		args = append(args, "--cpuset-mems", strconv.Quote(*v))
	}
	if v := opt.Detach; v != nil && *v != false {
		args = append(args, "--detach")
	}
	if v := opt.DetachKeys; v != nil {
		args = append(args, "--detach-keys", strconv.Quote(*v))
	}
	for _, v := range opt.Device {
		args = append(args, "--device", strconv.Quote(v))
	}
	for _, v := range opt.DeviceCgroupRule {
		args = append(args, "--device-cgroup-rule", strconv.Quote(v))
	}
	for _, v := range opt.DeviceReadBps {
		args = append(args, "--device-read-bps", strconv.Quote(v))
	}
	for _, v := range opt.DeviceReadIops {
		args = append(args, "--device-read-iops", strconv.Quote(v))
	}
	for _, v := range opt.DeviceWriteBps {
		args = append(args, "--device-write-bps", strconv.Quote(v))
	}
	for _, v := range opt.DeviceWriteIops {
		args = append(args, "--device-write-iops", strconv.Quote(v))
	}
	if v := opt.DisableContentTrust; v != nil && *v != false {
		args = append(args, "--disable-content-trust")
	}
	for _, v := range opt.Dns {
		args = append(args, "--dns", strconv.Quote(v))
	}
	for _, v := range opt.DnsOpt {
		args = append(args, "--dns-option", strconv.Quote(v))
	}
	for _, v := range opt.DnsOption {
		args = append(args, "--dns-option", strconv.Quote(v))
	}
	for _, v := range opt.DnsSearch {
		args = append(args, "--dns-search", strconv.Quote(v))
	}
	if v := opt.Entrypoint; v != nil {
		args = append(args, "--entrypoint", strconv.Quote(*v))
	}
	for k, v := range opt.Env {
		args = append(args, "--env", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for _, v := range opt.EnvFile {
		args = append(args, "--env-file", strconv.Quote(v))
	}
	for _, v := range opt.Expose {
		args = append(args, "--expose", strconv.Quote(v))
	}
	for _, v := range opt.GroupAdd {
		args = append(args, "--group-add", strconv.Quote(v))
	}
	if v := opt.HealthCmd; v != nil {
		args = append(args, "--health-cmd", strconv.Quote(*v))
	}
	if v := opt.HealthInterval; v != nil {
		args = append(args, "--health-interval", (*v).String())
	}
	if v := opt.HealthRetries; v != nil {
		args = append(args, "--health-retries", strconv.Itoa(*v))
	}
	if v := opt.HealthStartPeriod; v != nil {
		args = append(args, "--health-start-period", (*v).String())
	}
	if v := opt.HealthTimeout; v != nil {
		args = append(args, "--health-timeout", (*v).String())
	}
	if v := opt.Hostname; v != nil {
		args = append(args, "--hostname", strconv.Quote(*v))
	}
	if v := opt.Init; v != nil && *v != false {
		args = append(args, "--init")
	}
	if v := opt.Interactive; v != nil && *v != false {
		args = append(args, "--interactive")
	}
	if v := opt.Ip; v != nil {
		args = append(args, "--ip", strconv.Quote(*v))
	}
	if v := opt.Ip6; v != nil {
		args = append(args, "--ip6", strconv.Quote(*v))
	}
	if v := opt.Ipc; v != nil {
		args = append(args, "--ipc", strconv.Quote(*v))
	}
	if v := opt.Isolation; v != nil {
		args = append(args, "--isolation", strconv.Quote(*v))
	}
	if v := opt.KernelMemory; v != nil {
		args = append(args, "--kernel-memory", strconv.Quote(*v))
	}
	for k, v := range opt.Label {
		args = append(args, "--label", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for _, v := range opt.LabelFile {
		args = append(args, "--label-file", strconv.Quote(v))
	}
	for _, v := range opt.Link {
		args = append(args, "--link", strconv.Quote(v))
	}
	for _, v := range opt.LinkLocalIp {
		args = append(args, "--link-loal-ip", strconv.Quote(v))
	}
	if v := opt.LogDriver; v != nil {
		args = append(args, "--log-driver", strconv.Quote(*v))
	}
	for k, v := range opt.LogOpt {
		args = append(args, "--log-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := opt.MacAddress; v != nil {
		args = append(args, "--mac-address", strconv.Quote(*v))
	}
	if v := opt.Memory; v != nil {
		args = append(args, "--memory", strconv.Quote(*v))
	}
	if v := opt.MemoryReservation; v != nil {
		args = append(args, "--memory-reservation", strconv.Quote(*v))
	}
	if v := opt.MemorySwap; v != nil {
		args = append(args, "--memory-swap", strconv.Quote(*v))
	}
	if v := opt.MemorySwappiness; v != nil {
		args = append(args, "--memory-swappiness", strconv.FormatInt(*v, 10))
	}
	for k, v := range opt.Mount {
		args = append(args, "--mount", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := opt.Name; v != nil {
		args = append(args, "--name", strconv.Quote(*v))
	}
	if v := opt.Net; v != nil {
		args = append(args, "--network", strconv.Quote(*v))
	}
	if v := opt.Network; v != nil {
		args = append(args, "--network", strconv.Quote(*v))
	}
	for _, v := range opt.NetworkAlias {
		args = append(args, "--network-alias", strconv.Quote(v))
	}
	if v := opt.NoHealthcheck; v != nil && *v != false {
		args = append(args, "--no-healthcheck")
	}
	if v := opt.OomKillDisable; v != nil && *v != false {
		args = append(args, "--oom-kill-disable")
	}
	if v := opt.OomScoreAdj; v != nil {
		args = append(args, "--oom-secore-adj", strconv.Itoa(*v))
	}
	if v := opt.Pid; v != nil {
		args = append(args, "--pid", strconv.Quote(*v))
	}
	if v := opt.PidsLimit; v != nil {
		args = append(args, "--pids-limit", strconv.FormatInt(*v, 10))
	}
	if v := opt.Platform; v != nil {
		args = append(args, "--platform", strconv.Quote(*v))
	}
	if v := opt.Privileged; v != nil && *v != false {
		args = append(args, "--privileged")
	}
	for _, v := range opt.Publish {
		args = append(args, "--publish", strconv.Quote(v))
	}
	if v := opt.PublishAll; v != nil && *v != false {
		args = append(args, "--publish-all")
	}
	if v := opt.ReadOnly; v != nil && *v != false {
		args = append(args, "--readonly")
	}
	if v := opt.Restart; v != nil {
		args = append(args, "--restart", strconv.Quote(*v))
	}
	if v := opt.Rm; v != nil && *v != false {
		args = append(args, "--rm")
	}
	if v := opt.Runtime; v != nil {
		args = append(args, "--runtime", strconv.Quote(*v))
	}
	for k, v := range opt.SecurityOpt {
		args = append(args, "--security-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := opt.ShmSize; v != nil {
		args = append(args, "--shm-size", strconv.Quote(*v))
	}
	if v := opt.SigProxy; v != nil && *v != false {
		args = append(args, "--sig-proxy")
	}
	if v := opt.StopSignal; v != nil {
		args = append(args, "--stop-signal", strconv.Quote(*v))
	}
	if v := opt.StopTimeout; v != nil {
		args = append(args, "--stop-timeout", strconv.Itoa(*v))
	}
	for k, v := range opt.StorageOpt {
		args = append(args, "--storage-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for k, v := range opt.Sysctl {
		args = append(args, "--sysctl", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for _, v := range opt.Tmpfs {
		args = append(args, "--tmpfs", strconv.Quote(v))
	}
	if v := opt.Tty; v != nil && *v != false {
		args = append(args, "--tty")
	}
	for k, v := range opt.Ulimit {
		args = append(args, "--ulimit", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := opt.User; v != nil {
		args = append(args, "--user", strconv.Quote(*v))
	}
	if v := opt.Userns; v != nil {
		args = append(args, "--userns", strconv.Quote(*v))
	}
	if v := opt.Uts; v != nil {
		args = append(args, "--uts", strconv.Quote(*v))
	}
	for _, v := range opt.Volume {
		args = append(args, "--volume", strconv.Quote(v))
	}
	if v := opt.VolumeDriver; v != nil {
		args = append(args, "--volume-driver", strconv.Quote(*v))
	}
	for _, v := range opt.VolumesFrom {
		args = append(args, "--volumes-from", strconv.Quote(v))
	}
	if v := opt.Workdir; v != nil {
		args = append(args, "--workdir", strconv.Quote(*v))
	}
	for _, v := range opt.VolumesFrom {
		args = append(args, "--volumes-from", strconv.Quote(v))
	}
	args = append(args, opt.Image)
	args = append(args, opt.Args...)

	return args
}
