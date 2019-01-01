package aliases

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/k-kinzal/aliases/pkg/posix"
	"github.com/k-kinzal/aliases/pkg/util"
)

func NewCommand(ctx Context, schema Schema) (*posix.Cmd, error) {
	cmd := posix.Command("docker", "run")

	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.AddHost) {
		cmd.Args = append(cmd.Args, "--add-host", strconv.Quote(v))
	}
	for _, v := range schema.Attach {
		cmd.Args = append(cmd.Args, "--attach", strconv.Quote(v))
	}
	if v := schema.BlkioWeight; v != nil {
		cmd.Args = append(cmd.Args, "--blkio-weight", strconv.Quote(*v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.BlkioWeightDevice) {
		cmd.Args = append(cmd.Args, "--blkio-weight-device", strconv.Quote(v))
	}
	for _, v := range schema.CapAdd {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cap-add=%s", strconv.Quote(v)))
	}
	for _, v := range schema.CapDrop {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cap-drop=%s", strconv.Quote(v)))
	}
	if v := schema.CgroupParent; v != nil {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cgroup-parent==%s", strconv.Quote(*v)))
	}
	if v := schema.CIDFile; v != nil {
		cmd.Args = append(cmd.Args, "--cidfile", strconv.Quote(*v))
	}
	if v := schema.CPUPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-period", strconv.Quote(*v))
	}
	if v := schema.CPUQuota; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-quota", strconv.Quote(*v))
	}
	if v := schema.CPURtPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-rt-period", strconv.Quote(*v))
	}
	if v := schema.CPURtRuntime; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-rt-runtime", strconv.Quote(*v))
	}
	if v := schema.CPUShares; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-shares", strconv.Quote(*v))
	}
	if v := schema.CPUs; v != nil {
		cmd.Args = append(cmd.Args, "--cpus", strconv.Quote(*v))
	}
	if v := schema.CPUsetCPUs; v != nil {
		cmd.Args = append(cmd.Args, "--cpuset-cpus", strconv.Quote(*v))
	}
	if v := schema.CPUsetMems; v != nil {
		cmd.Args = append(cmd.Args, "--cpuset-mems", strconv.Quote(*v))
	}
	if v := schema.Detach; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--detach")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--detach\")", strconv.Quote(*v)))
		}
	}
	if v := schema.DetachKeys; v != nil {
		cmd.Args = append(cmd.Args, "--detach-keys", strconv.Quote(*v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.Device) {
		cmd.Args = append(cmd.Args, "--device", strconv.Quote(v))
	}
	for _, v := range schema.DeviceCgroupRule {
		cmd.Args = append(cmd.Args, "--device-cgroup-rule", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceReadBPS) {
		cmd.Args = append(cmd.Args, "--device-read-bps", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceReadIOPS) {
		cmd.Args = append(cmd.Args, "--device-read-iops", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceWriteBPS) {
		cmd.Args = append(cmd.Args, "--device-write-bps", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceWriteIOPS) {
		cmd.Args = append(cmd.Args, "--device-write-iops", strconv.Quote(v))
	}
	if v := schema.DisableContentTrust; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--disable-content-trust")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--disable-content-trust\")", strconv.Quote(*v)))
		}
	}
	for _, v := range schema.DNS {
		cmd.Args = append(cmd.Args, "--dns", strconv.Quote(v))
	}
	for _, v := range schema.DNSOption {
		cmd.Args = append(cmd.Args, "--dns-option", strconv.Quote(v))
	}
	for _, v := range schema.DNSSearch {
		cmd.Args = append(cmd.Args, "--dns-search", strconv.Quote(v))
	}
	if v := schema.Entrypoint; v != nil {
		cmd.Args = append(cmd.Args, "--entrypoint", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Env) {
		cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if (len(schema.Dependencies) > 0) && !ctx.HasDockerSocket() {
		cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("DOCKER_HOST=%s", strconv.Quote(ctx.DockerHost())))
	}
	if len(schema.Dependencies) > 0 {
		cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("ALIASES_PWD=%s", strconv.Quote("${ALIASES_PWD:-$PWD}")))
	}
	for _, v := range schema.EnvFile {
		cmd.Args = append(cmd.Args, "--env-file", strconv.Quote(v))
	}
	for _, v := range schema.Expose {
		cmd.Args = append(cmd.Args, "--expose", strconv.Quote(v))
	}
	for _, v := range schema.GroupAdd {
		cmd.Args = append(cmd.Args, "--group-add", strconv.Quote(v))
	}
	if v := schema.HealthCmd; v != nil {
		cmd.Args = append(cmd.Args, "--health-cmd", strconv.Quote(*v))
	}
	if v := schema.HealthInterval; v != nil {
		cmd.Args = append(cmd.Args, "--health-interval", strconv.Quote(*v))
	}
	if v := schema.HealthRetries; v != nil {
		cmd.Args = append(cmd.Args, "--health-retries", strconv.Quote(*v))
	}
	if v := schema.HealthStartPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--health-start-period", strconv.Quote(*v))
	}
	if v := schema.HealthTimeout; v != nil {
		cmd.Args = append(cmd.Args, "--health-timeout", strconv.Quote(*v))
	}
	if v := schema.Hostname; v != nil {
		cmd.Args = append(cmd.Args, "--hostname", strconv.Quote(*v))
	}
	if v := schema.Init; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--init")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--init\")", strconv.Quote(*v)))
		}
	}
	if v := schema.Interactive; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--interactive")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--interactive\")", strconv.Quote(*v)))
		}
	}
	if v := schema.IP; v != nil {
		cmd.Args = append(cmd.Args, "--ip", strconv.Quote(*v))
	}
	if v := schema.IP; v != nil {
		cmd.Args = append(cmd.Args, "--ip6", strconv.Quote(*v))
	}
	if v := schema.IPC; v != nil {
		cmd.Args = append(cmd.Args, "--ipc", strconv.Quote(*v))
	}
	if v := schema.Isolation; v != nil {
		cmd.Args = append(cmd.Args, "--isolation", strconv.Quote(*v))
	}
	if v := schema.KernelMemory; v != nil {
		cmd.Args = append(cmd.Args, "--kernel-memory", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Label) {
		cmd.Args = append(cmd.Args, "--label", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for _, v := range schema.LabelFile {
		cmd.Args = append(cmd.Args, "--label-file", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.Link) {
		cmd.Args = append(cmd.Args, "--link", strconv.Quote(v))
	}
	for _, v := range schema.LinkLocalIP {
		cmd.Args = append(cmd.Args, "--link-loal-ip", strconv.Quote(v))
	}
	if v := schema.LogDriver; v != nil {
		cmd.Args = append(cmd.Args, "--log-driver", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.LogOpt) {
		cmd.Args = append(cmd.Args, "--log-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.MacAddress; v != nil {
		cmd.Args = append(cmd.Args, "--mac-address", strconv.Quote(*v))
	}
	if v := schema.Memory; v != nil {
		cmd.Args = append(cmd.Args, "--memory", strconv.Quote(*v))
	}
	if v := schema.MemoryReservation; v != nil {
		cmd.Args = append(cmd.Args, "--memory-reservation", strconv.Quote(*v))
	}
	if v := schema.MemorySwap; v != nil {
		cmd.Args = append(cmd.Args, "--memory-swap", strconv.Quote(*v))
	}
	if v := schema.MemorySwappiness; v != nil {
		cmd.Args = append(cmd.Args, "--memory-swappiness", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Mount) {
		cmd.Args = append(cmd.Args, "--mount", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.Name; v != nil {
		cmd.Args = append(cmd.Args, "--name", strconv.Quote(*v))
	}
	if (len(schema.Dependencies) > 0) && !ctx.HasDockerSocket() {
		cmd.Args = append(cmd.Args, "--network", strconv.Quote("host"))
	} else if v := schema.Network; v != nil {
		cmd.Args = append(cmd.Args, "--network", strconv.Quote(*v))
	}
	for _, v := range schema.NetworkAlias {
		cmd.Args = append(cmd.Args, "--network-alias", strconv.Quote(v))
	}
	if v := schema.NoHealthcheck; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--no-healthcheck")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--no-healthcheck\")", strconv.Quote(*v)))
		}
	}
	if v := schema.OOMKillDisable; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--oom-kill-disable")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--oom-kill-disable\")", strconv.Quote(*v)))
		}
	}
	if v := schema.OOMScoreAdj; v != nil {
		cmd.Args = append(cmd.Args, "--oom-secore-adj", strconv.Quote(*v))
	}
	if v := schema.PID; v != nil {
		cmd.Args = append(cmd.Args, "--pid", strconv.Quote(*v))
	}
	if v := schema.PidsLimit; v != nil {
		cmd.Args = append(cmd.Args, "--pids-limit", strconv.Quote(*v))
	}
	if v := schema.Platform; v != nil {
		cmd.Args = append(cmd.Args, "--platform", strconv.Quote(*v))
	}
	if (len(schema.Dependencies) > 0) && ctx.HasDockerSocket() {
		cmd.Args = append(cmd.Args, "--privileged")
	} else if v := schema.Privileged; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--privileged")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--privileged\")", strconv.Quote(*v)))
		}
	}
	for _, v := range schema.Publish {
		cmd.Args = append(cmd.Args, "--publish", strconv.Quote(v))
	}
	if v := schema.PublishAll; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--publish-all")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--publish-all\")", strconv.Quote(*v)))
		}
	}
	if v := schema.ReadOnly; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--readonly")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--readonly\")", strconv.Quote(*v)))
		}
	}
	if v := schema.Restart; v != nil {
		cmd.Args = append(cmd.Args, "--restart", strconv.Quote(*v))
	}
	if v := schema.Rm; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--rm")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--rm\")", strconv.Quote(*v)))
		}
	}
	if v := schema.Runtime; v != nil {
		cmd.Args = append(cmd.Args, "--runtime", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.SecurityOpt) {
		cmd.Args = append(cmd.Args, "--security-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.ShmSize; v != nil {
		cmd.Args = append(cmd.Args, "--shm-size", strconv.Quote(*v))
	}
	if v := schema.SigProxy; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--sig-proxy")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--sig-proxy\")", strconv.Quote(*v)))
		}
	}
	if v := schema.StopSignal; v != nil {
		cmd.Args = append(cmd.Args, "--stop-signal", strconv.Quote(*v))
	}
	if v := schema.StopTimeout; v != nil {
		cmd.Args = append(cmd.Args, "--stop-timeout", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.StorageOpt) {
		cmd.Args = append(cmd.Args, "--storage-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Sysctl) {
		cmd.Args = append(cmd.Args, "--sysctl", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for _, v := range schema.Tmpfs {
		cmd.Args = append(cmd.Args, "--tmpfs", strconv.Quote(v))
	}
	if v := schema.TTY; v != nil && *v != "false" {
		if *v == "true" {
			cmd.Args = append(cmd.Args, "--tty")
		} else {
			cmd.Args = append(cmd.Args, fmt.Sprintf("$(test %s = \"true\" && echo \"--tty\")", strconv.Quote(*v)))
		}
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Ulimit) {
		cmd.Args = append(cmd.Args, "--ulimit", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.User; v != nil {
		cmd.Args = append(cmd.Args, "--user", strconv.Quote(util.ExpandColonDelimitedStringWithEnv(*v)))
	}
	if v := schema.Userns; v != nil {
		cmd.Args = append(cmd.Args, "--userns", strconv.Quote(*v))
	}
	if v := schema.UTS; v != nil {
		cmd.Args = append(cmd.Args, "--uts", strconv.Quote(*v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.Volume) {
		cmd.Args = append(cmd.Args, "--volume", strconv.Quote(v))
	}
	if (len(schema.Dependencies) > 0) && ctx.HasDockerSocket() {
		binary := BinaryManager{path.Join(ctx.HomePath(), "docker")}
		binarypath, err := binary.Get(schema.Docker.Image, schema.Docker.Tag)
		if err != nil {
			return nil, err
		}
		cmd.Args = append(cmd.Args, "--volume", strconv.Quote(fmt.Sprintf("%s:/usr/local/bin/docker", *binarypath)))
		if sock := ctx.DockerSocketPath(); sock != nil {
			cmd.Args = append(cmd.Args, "--volume", strconv.Quote(fmt.Sprintf("%s:/var/run/docker.sock", *sock)))
		}
	}
	for _, dep := range schema.Dependencies {
		cmd.Args = append(cmd.Args, "--volume", strconv.Quote(fmt.Sprintf("%s/%s:%s", ctx.ExportPath(), path.Base(dep), dep)))
	}
	if v := schema.VolumeDriver; v != nil {
		cmd.Args = append(cmd.Args, "--volume-driver", strconv.Quote(*v))
	}
	for _, v := range schema.VolumesFrom {
		cmd.Args = append(cmd.Args, "--volumes-from", strconv.Quote(v))
	}
	if v := schema.Workdir; v != nil {
		cmd.Args = append(cmd.Args, "--workdir", strconv.Quote(*v))
	}
	for _, v := range schema.VolumesFrom {
		cmd.Args = append(cmd.Args, "--volumes-from", strconv.Quote(v))
	}
	cmd.Args = append(cmd.Args, fmt.Sprintf("%s:${%s_VERSION:-\"%s\"}", schema.Image, strings.ToUpper(schema.FileName), schema.Tag))
	if schema.Command != nil {
		cmd.Args = append(cmd.Args, *schema.Command)
	}
	for _, arg := range schema.Args {
		if strings.Contains(arg, " ") {
			cmd.Args = append(cmd.Args, strconv.Quote(arg))
		} else {
			cmd.Args = append(cmd.Args, arg)
		}
	}

	return cmd, nil
}
