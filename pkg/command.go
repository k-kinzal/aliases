package aliases

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var (
	alphaRegexp = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
)

type AliasCommand struct {
	path string
	filename string
	cmd exec.Cmd
}

func (c *AliasCommand) ToDockerRunString() string {
	args := make([]string, 0)
	for i, arg := range c.cmd.Args {
		if i == 0 {
			continue
		}

		if strings.HasPrefix(arg, "--") {
			args = append(args, arg)
		} else if alphaRegexp.MatchString(arg) {
			args = append(args, arg)
		} else {
			args = append(args, fmt.Sprintf("%q", arg))
		}
	}

	return fmt.Sprintf("%s %s", c.cmd.Path, strings.Join(args, " "))
}

func (c *AliasCommand) ToString() string {
	return fmt.Sprintf("alias %s='%s'\n", c.filename, strings.Replace(c.ToDockerRunString(), "'", "\\'", -1))
}

func GenerateCommands(conf *AliasesConf, ctx *Context) []*AliasCommand {
	cmds := make([]*AliasCommand, 0)
	for _, c := range conf.Aliases {
		cmds = append(cmds, GenerateCommand(&c, ctx))
	}

	return cmds
}

func GenerateCommand(conf *AliasConf, ctx *Context) *AliasCommand {
	cmd := exec.Command("docker", "run")

	for _, v := range conf.DockerConf.DockerOpts.AddHost {
		cmd.Args = append(cmd.Args, "--add-host", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.Attach {
		cmd.Args = append(cmd.Args, "--attach", v)
	}
	if v := conf.DockerConf.DockerOpts.BlkioWeight; v != nil {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--blkio-weight=%d", *v))
	}
	for _, v := range conf.DockerConf.DockerOpts.BlkioWeightDevice {
		cmd.Args = append(cmd.Args, "--blkio-weight-device", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.CapAdd {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cap-add=%s", v))
	}
	for _, v := range conf.DockerConf.DockerOpts.CapDrop {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cap-drop=%s", v))
	}
	if v := conf.DockerConf.DockerOpts.CgroupParent; v != nil {
		cmd.Args = append(cmd.Args, fmt.Sprintf("--cgroup-parent==%s", *v))
	}
	if v := conf.DockerConf.DockerOpts.Cidfile; v != nil {
		cmd.Args = append(cmd.Args, "--cidfile", *v)
	}
	if v := conf.DockerConf.DockerOpts.CpuPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-period", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.CpuQuota; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-quota", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.CpuRtPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-rt-period", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.CpuRtRuntime; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-rt-runtime", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.CpuShares; v != nil {
		cmd.Args = append(cmd.Args, "--cpu-shares", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.Cpus; v != nil {
		cmd.Args = append(cmd.Args, "--cpus", strconv.FormatFloat(*v, 'f', 2, 64))
	}
	if v := conf.DockerConf.DockerOpts.CpusetCpus; v != nil {
		cmd.Args = append(cmd.Args, "--cpuset-cpus", *v)
	}
	if v := conf.DockerConf.DockerOpts.CpusetMems; v != nil {
		cmd.Args = append(cmd.Args, "--cpuset-mems", *v)
	}
	if v := conf.DockerConf.DockerOpts.Detach; v != false {
		cmd.Args = append(cmd.Args, "--detach")
	}
	if v := conf.DockerConf.DockerOpts.DetachKeys; v != nil {
		cmd.Args = append(cmd.Args, "--detach-keys", *v)
	}
	for _, v := range conf.DockerConf.DockerOpts.Device {
		cmd.Args = append(cmd.Args, "--device", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.DeviceCgroupRule {
		cmd.Args = append(cmd.Args, "--device-cgroup-rule", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.DeviceReadBps {
		cmd.Args = append(cmd.Args, "--device-read-bps", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.DeviceReadIops {
		cmd.Args = append(cmd.Args, "--device-read-iops", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.DeviceWriteBps {
		cmd.Args = append(cmd.Args, "--device-write-bps", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.DeviceWriteIops {
		cmd.Args = append(cmd.Args, "--device-write-iops", v)
	}
	if v := conf.DockerConf.DockerOpts.DisableContentTrust; v != false {
		cmd.Args = append(cmd.Args, "--disable-content-trust")
	}
	for _, v := range conf.DockerConf.DockerOpts.Dns {
		cmd.Args = append(cmd.Args, "--dns", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.DnsOption {
		cmd.Args = append(cmd.Args, "--dns-option", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.DnsSearch {
		cmd.Args = append(cmd.Args, "--dns-search", v)
	}
	if v := conf.DockerConf.DockerOpts.Entrypoint; v != nil {
		cmd.Args = append(cmd.Args, "--entrypoint", *v)
	}
	for k, v := range conf.DockerConf.DockerOpts.Env {
		cmd.Args = append(cmd.Args, "--env", fmt.Sprintf("%s=%s", k, v))
	}
	for _, v := range conf.DockerConf.DockerOpts.EnvFile {
		cmd.Args = append(cmd.Args, "--env-file", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.Expose {
		cmd.Args = append(cmd.Args, "--expose", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.GroupAdd {
		cmd.Args = append(cmd.Args, "--group-add", v)
	}
	if v := conf.DockerConf.DockerOpts.HealthCmd; v != nil {
		cmd.Args = append(cmd.Args, "--health-cmd", *v)
	}
	if v := conf.DockerConf.DockerOpts.HealthInterval; v != nil {
		cmd.Args = append(cmd.Args, "--health-interval", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.HealthRetries; v != nil {
		cmd.Args = append(cmd.Args, "--health-retries", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.HealthStartPeriod; v != nil {
		cmd.Args = append(cmd.Args, "--health-start-period", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.HealthTimeout; v != nil {
		cmd.Args = append(cmd.Args, "--health-timeout", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.Hostname; v != nil {
		cmd.Args = append(cmd.Args, "--hostname", *v)
	}
	if v := conf.DockerConf.DockerOpts.Init; v != false {
		cmd.Args = append(cmd.Args, "--init")
	}
	if v := conf.DockerConf.DockerOpts.Interactive; v != false {
		cmd.Args = append(cmd.Args, "--interactive")
	}
	if v := conf.DockerConf.DockerOpts.Ip; v != nil {
		cmd.Args = append(cmd.Args, "--ip", *v)
	}
	if v := conf.DockerConf.DockerOpts.Ip6; v != nil {
		cmd.Args = append(cmd.Args, "--ip6", *v)
	}
	if v := conf.DockerConf.DockerOpts.Ipc; v != nil {
		cmd.Args = append(cmd.Args, "--ipc", *v)
	}
	if v := conf.DockerConf.DockerOpts.Isolation; v != nil {
		cmd.Args = append(cmd.Args, "--isolation", *v)
	}
	if v := conf.DockerConf.DockerOpts.KernelMemory; v != nil {
		cmd.Args = append(cmd.Args, "--kernel-memory", strconv.Itoa(*v))
	}
	for k, v := range conf.DockerConf.DockerOpts.Label {
		cmd.Args = append(cmd.Args, "--label", fmt.Sprintf("%s=%s", k, v))
	}
	for _, v := range conf.DockerConf.DockerOpts.LabelFile {
		cmd.Args = append(cmd.Args, "--label-file", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.Link {
		cmd.Args = append(cmd.Args, "--link", v)
	}
	for _, v := range conf.DockerConf.DockerOpts.LinkLocalIp {
		cmd.Args = append(cmd.Args, "--link-loal-ip", v)
	}
	if v := conf.DockerConf.DockerOpts.LogDriver; v != nil {
		cmd.Args = append(cmd.Args, "--log-driver", *v)
	}
	for k, v := range conf.DockerConf.DockerOpts.LogOpt {
		cmd.Args = append(cmd.Args, "--log-opt", fmt.Sprintf("%s=%s", k, v))
	}
	if v := conf.DockerConf.DockerOpts.MacAddress; v != nil {
		cmd.Args = append(cmd.Args, "--mac-address", *v)
	}
	if v := conf.DockerConf.DockerOpts.Memory; v != nil {
		cmd.Args = append(cmd.Args, "--memory", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.MemoryReservation; v != nil {
		cmd.Args = append(cmd.Args, "--memory-reservation", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.MemorySwap; v != nil {
		cmd.Args = append(cmd.Args, "--memory-swap", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.MemorySwappiness; v != nil {
		cmd.Args = append(cmd.Args, "--memory-swappiness", strconv.Itoa(*v))
	}
	for k, v := range conf.DockerConf.DockerOpts.Mount {
		cmd.Args = append(cmd.Args, "--mount", fmt.Sprintf("%s=%s", k, v))
	}
	if v := conf.DockerConf.DockerOpts.Name; v != nil {
		cmd.Args = append(cmd.Args, "--name", *v)
	}
	if v := conf.DockerConf.DockerOpts.Network; v != nil {
		cmd.Args = append(cmd.Args, "--network", *v)
	}
	for _, v := range conf.DockerConf.DockerOpts.NetworkAlias {
		cmd.Args = append(cmd.Args, "--network-alias", v)
	}
	if v := conf.DockerConf.DockerOpts.NoHealthcheck; v != false {
		cmd.Args = append(cmd.Args, "--no-health-check")
	}
	if v := conf.DockerConf.DockerOpts.OomKillDisable; v != false {
		cmd.Args = append(cmd.Args, "--oom-kill-disable")
	}
	if v := conf.DockerConf.DockerOpts.OomScoreAdj; v != nil {
		cmd.Args = append(cmd.Args, "--oom-secore-adj", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.Pid; v != nil {
		cmd.Args = append(cmd.Args, "--pid", *v)
	}
	if v := conf.DockerConf.DockerOpts.PidsLimit; v != nil {
		cmd.Args = append(cmd.Args, "--pids-limit", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.Platform; v != nil {
		cmd.Args = append(cmd.Args, "--platform", *v)
	}
	if v := conf.DockerConf.DockerOpts.Privileged; v != false {
		cmd.Args = append(cmd.Args, "--privileged")
	}
	for _, v := range conf.DockerConf.DockerOpts.Publish {
		cmd.Args = append(cmd.Args, "--publish", v)
	}
	if v := conf.DockerConf.DockerOpts.PublishAll; v != false {
		cmd.Args = append(cmd.Args, "--publish-all")
	}
	if v := conf.DockerConf.DockerOpts.ReadOnly; v != false {
		cmd.Args = append(cmd.Args, "--readonly")
	}
	if v := conf.DockerConf.DockerOpts.Restart; v != nil {
		cmd.Args = append(cmd.Args, "--restart", *v)
	}
	if v := conf.DockerConf.DockerOpts.Rm; v != false {
		cmd.Args = append(cmd.Args, "--rm")
	}
	if v := conf.DockerConf.DockerOpts.Runtime; v != nil {
		cmd.Args = append(cmd.Args, "--runtime", *v)
	}
	for k, v := range conf.DockerConf.DockerOpts.SecurityOpt {
		cmd.Args = append(cmd.Args, "--security-opt", fmt.Sprintf("%s=%s", k, v))
	}
	if v := conf.DockerConf.DockerOpts.ShmSize; v != nil {
		cmd.Args = append(cmd.Args, "--shm-size", strconv.Itoa(*v))
	}
	if v := conf.DockerConf.DockerOpts.SigProxy; v != false {
		cmd.Args = append(cmd.Args, "--sig-proxy")
	}
	if v := conf.DockerConf.DockerOpts.StopSignal; v != nil {
		cmd.Args = append(cmd.Args, "--stop-signal", *v)
	}
	if v := conf.DockerConf.DockerOpts.StopTimeout; v != nil {
		cmd.Args = append(cmd.Args, "--stop-timeout", strconv.Itoa(*v))
	}
	for k, v := range conf.DockerConf.DockerOpts.StorageOpt {
		cmd.Args = append(cmd.Args, "--storage-opt", fmt.Sprintf("%s=%s", k, v))
	}
	for k, v := range conf.DockerConf.DockerOpts.Sysctl {
		cmd.Args = append(cmd.Args, "--sysctl", fmt.Sprintf("%s=%s", k, v))
	}
	for _, v := range conf.DockerConf.DockerOpts.Tmpfs {
		cmd.Args = append(cmd.Args, "--tmpfs", v)
	}
	if v := conf.DockerConf.DockerOpts.Tty; v != false {
		cmd.Args = append(cmd.Args, "--tty")
	}
	for k, v := range conf.DockerConf.DockerOpts.Ulimit {
		cmd.Args = append(cmd.Args, "--ulimit", fmt.Sprintf("%s=%s", k, v))
	}
	if v := conf.DockerConf.DockerOpts.User; v != nil {
		cmd.Args = append(cmd.Args, "--user", *v)
	}
	if v := conf.DockerConf.DockerOpts.Userns; v != nil {
		cmd.Args = append(cmd.Args, "--userns", *v)
	}
	if v := conf.DockerConf.DockerOpts.Uts; v != nil {
		cmd.Args = append(cmd.Args, "--uts", *v)
	}
	for _, v := range conf.DockerConf.DockerOpts.Volume {
		cmd.Args = append(cmd.Args, "--volume", v)
	}
	if v := conf.DockerConf.DockerOpts.VolumeDriver; v != nil {
		cmd.Args = append(cmd.Args, "--volume-driver", *v)
	}
	for _, v := range conf.DockerConf.DockerOpts.VolumesFrom {
		cmd.Args = append(cmd.Args, "--volumes-from", v)
	}
	if v := conf.DockerConf.DockerOpts.Workdir; v != nil {
		cmd.Args = append(cmd.Args, "--workdir", *v)
	}
	cmd.Args = append(cmd.Args, fmt.Sprintf("%s:${%s_VERSION:-%s}", conf.DockerConf.Image, strings.ToUpper(path.Base(conf.Path)), conf.DockerConf.Tag))
	if v := conf.DockerConf.Command; v != nil {
		cmd.Args = append(cmd.Args, *v)
	}
	for _, v := range conf.DockerConf.DockerOpts.VolumesFrom {
		cmd.Args = append(cmd.Args, v)
	}

	cmd.Env = os.Environ()

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return &AliasCommand {
		path: conf.Path,
		filename: path.Base(conf.Path),
		cmd: *cmd,
	}
}
