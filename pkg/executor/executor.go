package executor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	pathes "path"
	"strconv"
	"strings"

	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/k-kinzal/aliases/pkg/util"
	"github.com/k-kinzal/aliases/pkg/yaml"
)

type Executor struct {
	schemas map[string]yaml.Schema
}

func (e *Executor) AddSchema(path string, schema yaml.Schema) {
	e.schemas[path] = schema
}

func (e *Executor) Pathes(depth int) []string {
	keys := make([]string, 0)
	for key := range e.schemas {
		keys = append(keys, key)
	}
	return keys
}

func (e *Executor) Command(ctx context.Context, path string) (*exec.Cmd, error) {
	schema, ok := e.schemas[path]
	if !ok {
		return nil, fmt.Errorf("logic error: %s is not defined", path)
	}

	cmd := exec.Command("docker", "run")

	args := make([]string, 0)
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.AddHost) {
		args = append(args, "--add-host", strconv.Quote(v))
	}
	for _, v := range schema.Attach {
		args = append(args, "--attach", strconv.Quote(v))
	}
	if v := schema.BlkioWeight; v != nil {
		args = append(args, "--blkio-weight", strconv.Quote(*v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.BlkioWeightDevice) {
		args = append(args, "--blkio-weight-device", strconv.Quote(v))
	}
	for _, v := range schema.CapAdd {
		args = append(args, fmt.Sprintf("--cap-add=%s", strconv.Quote(v)))
	}
	for _, v := range schema.CapDrop {
		args = append(args, fmt.Sprintf("--cap-drop=%s", strconv.Quote(v)))
	}
	if v := schema.CgroupParent; v != nil {
		args = append(args, fmt.Sprintf("--cgroup-parent==%s", strconv.Quote(*v)))
	}
	if v := schema.CIDFile; v != nil {
		args = append(args, "--cidfile", strconv.Quote(*v))
	}
	if v := schema.CPUPeriod; v != nil {
		args = append(args, "--cpu-period", strconv.Quote(*v))
	}
	if v := schema.CPUQuota; v != nil {
		args = append(args, "--cpu-quota", strconv.Quote(*v))
	}
	if v := schema.CPURtPeriod; v != nil {
		args = append(args, "--cpu-rt-period", strconv.Quote(*v))
	}
	if v := schema.CPURtRuntime; v != nil {
		args = append(args, "--cpu-rt-runtime", strconv.Quote(*v))
	}
	if v := schema.CPUShares; v != nil {
		args = append(args, "--cpu-shares", strconv.Quote(*v))
	}
	if v := schema.CPUs; v != nil {
		args = append(args, "--cpus", strconv.Quote(*v))
	}
	if v := schema.CPUsetCPUs; v != nil {
		args = append(args, "--cpuset-cpus", strconv.Quote(*v))
	}
	if v := schema.CPUsetMems; v != nil {
		args = append(args, "--cpuset-mems", strconv.Quote(*v))
	}
	if v := schema.Detach; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--detach")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--detach\")", strconv.Quote(*v)))
		}
	}
	if v := schema.DetachKeys; v != nil {
		args = append(args, "--detach-keys", strconv.Quote(*v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.Device) {
		args = append(args, "--device", strconv.Quote(v))
	}
	for _, v := range schema.DeviceCgroupRule {
		args = append(args, "--device-cgroup-rule", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceReadBPS) {
		args = append(args, "--device-read-bps", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceReadIOPS) {
		args = append(args, "--device-read-iops", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceWriteBPS) {
		args = append(args, "--device-write-bps", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.DeviceWriteIOPS) {
		args = append(args, "--device-write-iops", strconv.Quote(v))
	}
	if v := schema.DisableContentTrust; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--disable-content-trust")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--disable-content-trust\")", strconv.Quote(*v)))
		}
	}
	for _, v := range schema.DNS {
		args = append(args, "--dns", strconv.Quote(v))
	}
	for _, v := range schema.DNSOpt {
		args = append(args, "--dns-option", strconv.Quote(v))
	}
	for _, v := range schema.DNSOption {
		args = append(args, "--dns-option", strconv.Quote(v))
	}
	for _, v := range schema.DNSSearch {
		args = append(args, "--dns-search", strconv.Quote(v))
	}
	if v := schema.Entrypoint; v != nil {
		args = append(args, "--entrypoint", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Env) {
		args = append(args, "--env", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if (len(schema.Dependencies) > 0) && (ctx.DockerSockType() == context.DockerSockTypeRemote) {
		args = append(args, "--env", fmt.Sprintf("DOCKER_HOST=%s", strconv.Quote(ctx.DockerRemoteHost())))
	}
	if len(schema.Dependencies) > 0 {
		args = append(args, "--env", fmt.Sprintf("ALIASES_PWD=%s", strconv.Quote("${ALIASES_PWD:-$PWD}")))
	}
	for _, v := range schema.EnvFile {
		args = append(args, "--env-file", strconv.Quote(v))
	}
	for _, v := range schema.Expose {
		args = append(args, "--expose", strconv.Quote(v))
	}
	for _, v := range schema.GroupAdd {
		args = append(args, "--group-add", strconv.Quote(v))
	}
	if v := schema.HealthCmd; v != nil {
		args = append(args, "--health-cmd", strconv.Quote(*v))
	}
	if v := schema.HealthInterval; v != nil {
		args = append(args, "--health-interval", strconv.Quote(*v))
	}
	if v := schema.HealthRetries; v != nil {
		args = append(args, "--health-retries", strconv.Quote(*v))
	}
	if v := schema.HealthStartPeriod; v != nil {
		args = append(args, "--health-start-period", strconv.Quote(*v))
	}
	if v := schema.HealthTimeout; v != nil {
		args = append(args, "--health-timeout", strconv.Quote(*v))
	}
	if v := schema.Hostname; v != nil {
		args = append(args, "--hostname", strconv.Quote(*v))
	}
	if v := schema.Init; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--init")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--init\")", strconv.Quote(*v)))
		}
	}
	if v := schema.Interactive; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--interactive")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--interactive\")", strconv.Quote(*v)))
		}
	}
	if v := schema.IP; v != nil {
		args = append(args, "--ip", strconv.Quote(*v))
	}
	if v := schema.IP; v != nil {
		args = append(args, "--ip6", strconv.Quote(*v))
	}
	if v := schema.IPC; v != nil {
		args = append(args, "--ipc", strconv.Quote(*v))
	}
	if v := schema.Isolation; v != nil {
		args = append(args, "--isolation", strconv.Quote(*v))
	}
	if v := schema.KernelMemory; v != nil {
		args = append(args, "--kernel-memory", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Label) {
		args = append(args, "--label", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for _, v := range schema.LabelFile {
		args = append(args, "--label-file", strconv.Quote(v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.Link) {
		args = append(args, "--link", strconv.Quote(v))
	}
	for _, v := range schema.LinkLocalIP {
		args = append(args, "--link-loal-ip", strconv.Quote(v))
	}
	if v := schema.LogDriver; v != nil {
		args = append(args, "--log-driver", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.LogOpt) {
		args = append(args, "--log-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.MacAddress; v != nil {
		args = append(args, "--mac-address", strconv.Quote(*v))
	}
	if v := schema.Memory; v != nil {
		args = append(args, "--memory", strconv.Quote(*v))
	}
	if v := schema.MemoryReservation; v != nil {
		args = append(args, "--memory-reservation", strconv.Quote(*v))
	}
	if v := schema.MemorySwap; v != nil {
		args = append(args, "--memory-swap", strconv.Quote(*v))
	}
	if v := schema.MemorySwappiness; v != nil {
		args = append(args, "--memory-swappiness", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Mount) {
		args = append(args, "--mount", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.Name; v != nil {
		args = append(args, "--name", strconv.Quote(*v))
	}
	if (len(schema.Dependencies) > 0) && (ctx.DockerSockType() == context.DockerSockTypeRemote) {
		args = append(args, "--network", strconv.Quote("host"))
	} else if v := schema.Network; v != nil {
		args = append(args, "--network", strconv.Quote(*v))
	}
	for _, v := range schema.NetworkAlias {
		args = append(args, "--network-alias", strconv.Quote(v))
	}
	if v := schema.NoHealthcheck; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--no-healthcheck")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--no-healthcheck\")", strconv.Quote(*v)))
		}
	}
	if v := schema.OOMKillDisable; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--oom-kill-disable")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--oom-kill-disable\")", strconv.Quote(*v)))
		}
	}
	if v := schema.OOMScoreAdj; v != nil {
		args = append(args, "--oom-secore-adj", strconv.Quote(*v))
	}
	if v := schema.PID; v != nil {
		args = append(args, "--pid", strconv.Quote(*v))
	}
	if v := schema.PidsLimit; v != nil {
		args = append(args, "--pids-limit", strconv.Quote(*v))
	}
	if v := schema.Platform; v != nil {
		args = append(args, "--platform", strconv.Quote(*v))
	}
	if (len(schema.Dependencies) > 0) && (ctx.DockerSockType() == context.DockerSockTypeSock) {
		args = append(args, "--privileged")
	} else if v := schema.Privileged; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--privileged")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--privileged\")", strconv.Quote(*v)))
		}
	}
	for _, v := range schema.Publish {
		args = append(args, "--publish", strconv.Quote(v))
	}
	if v := schema.PublishAll; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--publish-all")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--publish-all\")", strconv.Quote(*v)))
		}
	}
	if v := schema.ReadOnly; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--readonly")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--readonly\")", strconv.Quote(*v)))
		}
	}
	if v := schema.Restart; v != nil {
		args = append(args, "--restart", strconv.Quote(*v))
	}
	if v := schema.Rm; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--rm")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--rm\")", strconv.Quote(*v)))
		}
	}
	if v := schema.Runtime; v != nil {
		args = append(args, "--runtime", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.SecurityOpt) {
		args = append(args, "--security-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.ShmSize; v != nil {
		args = append(args, "--shm-size", strconv.Quote(*v))
	}
	if v := schema.SigProxy; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--sig-proxy")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--sig-proxy\")", strconv.Quote(*v)))
		}
	}
	if v := schema.StopSignal; v != nil {
		args = append(args, "--stop-signal", strconv.Quote(*v))
	}
	if v := schema.StopTimeout; v != nil {
		args = append(args, "--stop-timeout", strconv.Quote(*v))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.StorageOpt) {
		args = append(args, "--storage-opt", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Sysctl) {
		args = append(args, "--sysctl", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	for _, v := range schema.Tmpfs {
		args = append(args, "--tmpfs", strconv.Quote(v))
	}
	if v := schema.TTY; v != nil && *v != "false" {
		if *v == "true" {
			args = append(args, "--tty")
		} else {
			args = append(args, fmt.Sprintf("$(test %s = \"true\" && echo \"--tty\")", strconv.Quote(*v)))
		}
	}
	for k, v := range util.ExpandStringKeyMapWithEnv(schema.Ulimit) {
		args = append(args, "--ulimit", fmt.Sprintf("%s=%s", k, strconv.Quote(v)))
	}
	if v := schema.User; v != nil {
		args = append(args, "--user", strconv.Quote(util.ExpandColonDelimitedStringWithEnv(*v)))
	}
	if v := schema.Userns; v != nil {
		args = append(args, "--userns", strconv.Quote(*v))
	}
	if v := schema.UTS; v != nil {
		args = append(args, "--uts", strconv.Quote(*v))
	}
	for _, v := range util.ExpandColonDelimitedStringListWithEnv(schema.Volume) {
		args = append(args, "--volume", strconv.Quote(v))
	}
	if (len(schema.Dependencies) > 0) && (ctx.DockerSockType() == context.DockerSockTypeSock) {
		args = append(args, "--volume", strconv.Quote(fmt.Sprintf("%s:/usr/local/bin/docker", ctx.DockerBinaryPath())))
		args = append(args, "--volume", strconv.Quote(fmt.Sprintf("%s:/var/run/docker.sock", ctx.DockerSockPath())))
	}
	for _, dep := range schema.Dependencies {
		args = append(args, "--volume", strconv.Quote(fmt.Sprintf("%s/%s:%s", ctx.GetExportPath(), pathes.Base(dep), dep)))
	}
	if v := schema.VolumeDriver; v != nil {
		args = append(args, "--volume-driver", strconv.Quote(*v))
	}
	for _, v := range schema.VolumesFrom {
		args = append(args, "--volumes-from", strconv.Quote(v))
	}
	if v := schema.Workdir; v != nil {
		args = append(args, "--workdir", strconv.Quote(*v))
	}
	for _, v := range schema.VolumesFrom {
		args = append(args, "--volumes-from", strconv.Quote(v))
	}
	args = append(args, fmt.Sprintf("%s:${%s_VERSION:-\"%s\"}", schema.Image, strings.ToUpper(pathes.Base(path)), schema.Tag))
	if schema.Command != nil {
		args = append(args, *schema.Command)
	}
	args = append(args, schema.Args...)

	cmd.Args = append(cmd.Args, args...)

	cmd.Env = os.Environ()

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd, nil
}

func (e *Executor) Commands(ctx context.Context) (map[string]exec.Cmd, error) {
	commands := make(map[string]exec.Cmd, len(e.schemas))
	for _, path := range e.Pathes(0) {
		cmd, err := e.Command(ctx, path)
		if err != nil {
			return nil, err
		}
		commands[path] = *cmd
	}
	return commands, nil
}

func New(ctx context.Context) (*Executor, error) {
	if _, err := os.Stat(ctx.GetConfPath()); os.IsNotExist(err) {
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	buf, err := ioutil.ReadFile(ctx.GetConfPath())
	if err != nil {
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	schemas, err := yaml.UnmarshalConfFile(buf)
	if err != nil {
		return nil, err
	}

	e := Executor{map[string]yaml.Schema{}}
	for path, schema := range schemas {
		e.AddSchema(path, schema)
	}

	return &e, nil
}
