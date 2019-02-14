package script

import (
	"fmt"
	"path"
	"strings"

	"github.com/k-kinzal/aliases/pkg/aliases/context"

	"github.com/imdario/mergo"

	"github.com/k-kinzal/aliases/pkg/posix"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/docker"
)

// Script is the actual of command alises.
type Script struct {
	path   string
	binary struct {
		image string
		tag   string
	}
	docker   func(args []string, option docker.RunOption) *posix.Cmd
	relative []*Script
}

// Path returns export path.
func (script *Script) Path(p string) string {
	return path.Join(p, script.path)
}

// FileName returns script name.
func (script *Script) FileName() string {
	return path.Base(script.path)
}

// String returns docker run string
func (script *Script) String() string {
	return script.docker(nil, docker.RunOption{}).String()
}

// newDockerRunCommand creates a new docker run command.
func newDockerRunCommand(client *docker.Client, image string, args []string, option docker.RunOption) func(args []string, option docker.RunOption) *posix.Cmd {
	return func(overrideArgs []string, overrideOption docker.RunOption) *posix.Cmd {
		if len(overrideArgs) > 0 {
			args = overrideArgs
		}
		opt := overrideOption
		if err := mergo.Merge(&opt, option, mergo.WithAppendSlice); err != nil {
			panic(err)
		}
		return client.Run(image, args, opt)
	}
}

// NewScript creates a new Script.
func NewScript(client *docker.Client, opt config.Option) *Script {
	script := &Script{}
	// global
	script.path = strings.Replace(fmt.Sprintf("%s/%s", opt.Namespace, opt.FileName), "//", "/", -1)
	// docker binary
	script.binary.image = opt.Docker.Image
	script.binary.tag = opt.Docker.Tag
	// docker run command
	// image
	image := fmt.Sprintf("%s:${%s_VERSION:-\"%s\"}", opt.Image, strings.ToUpper(opt.FileName), opt.Tag)
	// args
	args := make([]string, 0)
	if opt.Command != nil {
		args = append(args, *opt.Command)
	}
	args = append(args, opt.Args...)
	// options
	o := docker.RunOption{}
	o.AddHost = ExpandColonDelimitedStringListWithEnv(opt.AddHost)
	o.Attach = opt.Attach
	o.BlkioWeight = opt.BlkioWeight
	o.BlkioWeightDevice = ExpandColonDelimitedStringListWithEnv(opt.BlkioWeightDevice)
	o.CIDFile = opt.CIDFile
	o.CPUPeriod = opt.CPUPeriod
	o.CPUQuota = opt.CPUQuota
	o.CPURtPeriod = opt.CPURtPeriod
	o.CPURtRuntime = opt.CPURtRuntime
	o.CPUShares = opt.CPUShares
	o.CPUs = opt.CPUs
	o.CPUsetCPUs = opt.CPUsetCPUs
	o.CPUsetMems = opt.CPUsetMems
	o.CapAdd = opt.CapAdd
	o.CapDrop = opt.CapDrop
	o.CgroupParent = opt.CgroupParent
	o.DNS = opt.DNS
	o.DNSOption = opt.DNSOption
	o.DNSSearch = opt.DNSSearch
	o.Detach = opt.Detach
	o.DetachKeys = opt.DetachKeys
	o.Device = ExpandColonDelimitedStringListWithEnv(opt.Device)
	o.DeviceCgroupRule = opt.DeviceCgroupRule
	o.DeviceReadBPS = ExpandColonDelimitedStringListWithEnv(opt.DeviceReadBPS)
	o.DeviceReadIOPS = ExpandColonDelimitedStringListWithEnv(opt.DeviceReadIOPS)
	o.DeviceWriteBPS = ExpandColonDelimitedStringListWithEnv(opt.DeviceWriteBPS)
	o.DeviceWriteIOPS = ExpandColonDelimitedStringListWithEnv(opt.DeviceWriteIOPS)
	o.DisableContentTrust = opt.DisableContentTrust
	o.Domainname = opt.Domainname
	o.Entrypoint = opt.Entrypoint
	o.Env = ExpandStringKeyMapWithEnv(opt.Env)
	o.EnvFile = opt.EnvFile
	o.Expose = opt.Expose
	o.GroupAdd = opt.GroupAdd
	o.HealthCmd = opt.HealthCmd
	o.HealthInterval = opt.HealthInterval
	o.HealthRetries = opt.HealthRetries
	o.HealthStartPeriod = opt.HealthStartPeriod
	o.HealthTimeout = opt.HealthTimeout
	o.Hostname = opt.Hostname
	o.IP = opt.IP
	o.IP6 = opt.IP6
	o.IPC = opt.IPC
	o.Init = opt.Init
	o.Interactive = opt.Interactive
	o.Isolation = opt.Isolation
	o.KernelMemory = opt.KernelMemory
	o.Label = ExpandStringKeyMapWithEnv(opt.Label)
	o.LabelFile = opt.LabelFile
	o.Link = ExpandColonDelimitedStringListWithEnv(opt.Link)
	o.LinkLocalIP = opt.LinkLocalIP
	o.LogDriver = opt.LogDriver
	o.LogOpt = ExpandStringKeyMapWithEnv(opt.LogOpt)
	o.MacAddress = opt.MacAddress
	o.Memory = opt.Memory
	o.MemoryReservation = opt.MemoryReservation
	o.MemorySwap = opt.MemorySwap
	o.MemorySwappiness = opt.MemorySwappiness
	o.Mount = ExpandStringKeyMapWithEnv(opt.Mount)
	o.Name = opt.Name
	o.Network = opt.Network
	o.NetworkAlias = opt.NetworkAlias
	o.NoHealthcheck = opt.NoHealthcheck
	o.OOMKillDisable = opt.OOMKillDisable
	o.OOMScoreAdj = opt.OOMScoreAdj
	o.PID = opt.PID
	o.PidsLimit = opt.PidsLimit
	o.Platform = opt.Platform
	o.Privileged = opt.Privileged
	o.Publish = opt.Publish
	o.PublishAll = opt.PublishAll
	o.ReadOnly = opt.ReadOnly
	o.Restart = opt.Restart
	o.Rm = opt.Rm
	o.Runtime = opt.Runtime
	o.SecurityOpt = ExpandStringKeyMapWithEnv(opt.SecurityOpt)
	o.ShmSize = opt.ShmSize
	o.SigProxy = opt.SigProxy
	o.StopSignal = opt.StopSignal
	o.StopTimeout = opt.StopTimeout
	o.StorageOpt = ExpandStringKeyMapWithEnv(opt.StorageOpt)
	o.Sysctl = ExpandStringKeyMapWithEnv(opt.Sysctl)
	o.TTY = opt.TTY
	o.Tmpfs = opt.Tmpfs
	o.UTS = opt.UTS
	o.Ulimit = ExpandStringKeyMapWithEnv(opt.Ulimit)
	if opt.User != nil {
		user := ExpandColonDelimitedStringWithEnv(*opt.User)
		o.User = &user
	}
	o.Userns = opt.Userns
	o.Volume = ExpandColonDelimitedStringListWithEnv(opt.Volume)
	o.VolumeDriver = opt.VolumeDriver
	o.VolumesFrom = opt.VolumesFrom
	o.Workdir = opt.Workdir
	// dependencies
	if len(opt.Dependencies) > 0 {
		if opt.Env == nil {
			o.Env = make(map[string]string, 0)
		}
		if opt.Volume == nil {
			o.Volume = make([]string, 0)
		}
		o.Env["ALIASES_PWD"] = "${ALIASES_PWD:-$PWD}"
		if sock := client.Sock(); sock != nil {
			// unix socket
			literal := "true"
			o.Privileged = &literal
			o.Volume = append(o.Volume, fmt.Sprintf("%s:/usr/local/bin/docker", opt.Binary(context.BinaryPath()).Path))
			o.Volume = append(o.Volume, fmt.Sprintf("%s:/var/run/docker.sock", *sock))
		} else {
			// tcp, http...
			literal := "host"
			o.Network = &literal
			o.Env["DOCKER_HOST"] = client.Host()
		}
		for _, dep := range opt.Dependencies {
			o.Volume = append(o.Volume, fmt.Sprintf("%s:%s", path.Join(context.ExportPath(), dep.Namespace, dep.FileName), dep.Path))
		}
	}
	script.docker = newDockerRunCommand(client, image, args, o)
	// relative
	relative := make([]*Script, len(opt.Dependencies))
	for i, dep := range opt.Dependencies {
		relative[i] = NewScript(client, *dep)
	}
	script.relative = relative

	return script
}
