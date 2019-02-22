package script

import (
	"fmt"
	"strings"

	"github.com/k-kinzal/aliases/pkg/util"

	"github.com/imdario/mergo"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/docker"
	"github.com/k-kinzal/aliases/pkg/posix"
)

// DockerRunAdapter adapts docker run information from spec to a form that can be used in aliases.
type DockerRunAdapter yaml.Option

// Image returns image of docker run.
func (adpt *DockerRunAdapter) Image() string {
	spec := (*yaml.Option)(adpt)

	return fmt.Sprintf("%s:${%s_VERSION:-\"%s\"}", spec.Image, strings.ToUpper(spec.Path.Base()), spec.Tag)
}

// Args returns arguments of docker run.
func (adpt *DockerRunAdapter) Args() []string {
	spec := (*yaml.Option)(adpt)

	return spec.Args
}

// Option returns option of docker run.
func (adpt *DockerRunAdapter) Option() *docker.RunOption {
	spec := (*yaml.Option)(adpt)

	opt := docker.RunOption{}
	opt.AddHost = ExpandColonDelimitedStringListWithEnv(spec.AddHost)
	opt.Attach = spec.Attach
	opt.BlkioWeight = spec.BlkioWeight
	opt.BlkioWeightDevice = ExpandColonDelimitedStringListWithEnv(spec.BlkioWeightDevice)
	opt.CIDFile = spec.CIDFile
	opt.CPUPeriod = spec.CPUPeriod
	opt.CPUQuota = spec.CPUQuota
	opt.CPURtPeriod = spec.CPURtPeriod
	opt.CPURtRuntime = spec.CPURtRuntime
	opt.CPUShares = spec.CPUShares
	opt.CPUs = spec.CPUs
	opt.CPUsetCPUs = spec.CPUsetCPUs
	opt.CPUsetMems = spec.CPUsetMems
	opt.CapAdd = spec.CapAdd
	opt.CapDrop = spec.CapDrop
	opt.CgroupParent = spec.CgroupParent
	opt.DNS = spec.DNS
	opt.DNSOption = spec.DNSOption
	opt.DNSSearch = spec.DNSSearch
	opt.Detach = spec.Detach
	opt.DetachKeys = spec.DetachKeys
	opt.Device = ExpandColonDelimitedStringListWithEnv(spec.Device)
	opt.DeviceCgroupRule = spec.DeviceCgroupRule
	opt.DeviceReadBPS = ExpandColonDelimitedStringListWithEnv(spec.DeviceReadBPS)
	opt.DeviceReadIOPS = ExpandColonDelimitedStringListWithEnv(spec.DeviceReadIOPS)
	opt.DeviceWriteBPS = ExpandColonDelimitedStringListWithEnv(spec.DeviceWriteBPS)
	opt.DeviceWriteIOPS = ExpandColonDelimitedStringListWithEnv(spec.DeviceWriteIOPS)
	opt.DisableContentTrust = spec.DisableContentTrust
	opt.Domainname = spec.Domainname
	if spec.Command != nil {
		opt.Entrypoint = spec.Command
	} else {
		opt.Entrypoint = spec.Entrypoint
	}
	opt.Env = ExpandStringKeyMapWithEnv(spec.Env)
	opt.EnvFile = spec.EnvFile
	opt.Expose = spec.Expose
	opt.GroupAdd = spec.GroupAdd
	opt.HealthCmd = spec.HealthCmd
	opt.HealthInterval = spec.HealthInterval
	opt.HealthRetries = spec.HealthRetries
	opt.HealthStartPeriod = spec.HealthStartPeriod
	opt.HealthTimeout = spec.HealthTimeout
	opt.Hostname = spec.Hostname
	opt.IP = spec.IP
	opt.IP6 = spec.IP6
	opt.IPC = spec.IPC
	opt.Init = spec.Init
	opt.Interactive = spec.Interactive
	opt.Isolation = spec.Isolation
	opt.KernelMemory = spec.KernelMemory
	opt.Label = ExpandStringKeyMapWithEnv(spec.Label)
	opt.LabelFile = spec.LabelFile
	opt.Link = ExpandColonDelimitedStringListWithEnv(spec.Link)
	opt.LinkLocalIP = spec.LinkLocalIP
	opt.LogDriver = spec.LogDriver
	opt.LogOpt = ExpandStringKeyMapWithEnv(spec.LogOpt)
	opt.MacAddress = spec.MacAddress
	opt.Memory = spec.Memory
	opt.MemoryReservation = spec.MemoryReservation
	opt.MemorySwap = spec.MemorySwap
	opt.MemorySwappiness = spec.MemorySwappiness
	opt.Mount = ExpandStringKeyMapWithEnv(spec.Mount)
	opt.Name = spec.Name
	opt.Network = spec.Network
	opt.NetworkAlias = spec.NetworkAlias
	opt.NoHealthcheck = spec.NoHealthcheck
	opt.OOMKillDisable = spec.OOMKillDisable
	opt.OOMScoreAdj = spec.OOMScoreAdj
	opt.PID = spec.PID
	opt.PidsLimit = spec.PidsLimit
	opt.Platform = spec.Platform
	opt.Privileged = spec.Privileged
	opt.Publish = spec.Publish
	opt.PublishAll = spec.PublishAll
	opt.ReadOnly = spec.ReadOnly
	opt.Restart = spec.Restart
	opt.Rm = spec.Rm
	opt.Runtime = spec.Runtime
	opt.SecurityOpt = ExpandStringKeyMapWithEnv(spec.SecurityOpt)
	opt.ShmSize = spec.ShmSize
	opt.SigProxy = spec.SigProxy
	opt.StopSignal = spec.StopSignal
	opt.StopTimeout = spec.StopTimeout
	opt.StorageOpt = ExpandStringKeyMapWithEnv(spec.StorageOpt)
	opt.Sysctl = ExpandStringKeyMapWithEnv(spec.Sysctl)
	opt.TTY = spec.TTY
	opt.Tmpfs = spec.Tmpfs
	opt.UTS = spec.UTS
	opt.Ulimit = ExpandStringKeyMapWithEnv(spec.Ulimit)
	if spec.User != nil {
		user := ExpandColonDelimitedStringWithEnv(*spec.User)
		opt.User = &user
	}
	opt.Userns = spec.Userns
	opt.Volume = ExpandColonDelimitedStringListWithEnv(spec.Volume)
	opt.VolumeDriver = spec.VolumeDriver
	opt.VolumesFrom = spec.VolumesFrom
	opt.Workdir = spec.Workdir

	return &opt
}

// Command returns a command to download the docker run.
func (adpt *DockerRunAdapter) Command(client *docker.Client, overrideArgs []string, overrideOption docker.RunOption) *posix.Cmd {
	image := adpt.Image()
	args := adpt.Args()
	if len(overrideArgs) > 0 {
		args = overrideArgs
	}
	opt := overrideOption
	if err := mergo.Merge(&opt, adpt.Option(), mergo.WithAppendSlice); err != nil {
		panic(err)
	}

	opt.AddHost = util.UniqueStringSlice(opt.AddHost)
	opt.BlkioWeightDevice = util.UniqueStringSlice(opt.BlkioWeightDevice)
	opt.Device = util.UniqueStringSlice(opt.Device)
	opt.DeviceReadBPS = util.UniqueStringSlice(opt.DeviceReadBPS)
	opt.DeviceReadIOPS = util.UniqueStringSlice(opt.DeviceReadIOPS)
	opt.DeviceWriteBPS = util.UniqueStringSlice(opt.DeviceWriteBPS)
	opt.DeviceWriteIOPS = util.UniqueStringSlice(opt.DeviceWriteIOPS)
	opt.Link = util.UniqueStringSlice(opt.Link)
	opt.Volume = util.UniqueStringSlice(opt.Volume)

	return client.Run(image, args, opt)
}

// adaptDockerRun returns DockerRunAdapter.
func adaptDockerRun(spec yaml.Option) *DockerRunAdapter {
	return (*DockerRunAdapter)(&spec)
}
