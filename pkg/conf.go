package aliases

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
)


type AliasConf struct {
	Path         string
	Dependencies []AliasConf

	DockerConf struct {
		Image         string
		Tag           string
		Command       *string
		Args          []string
		DockerOpts struct {
			AddHost             []string
			Attach              []string
			BlkioWeight         *uint16
			BlkioWeightDevice   []string
			CapAdd              []string
			CapDrop             []string
			CgroupParent        *string
			Cidfile             *string
			CpuPeriod           *int
			CpuQuota            *int
			CpuRtPeriod         *int
			CpuRtRuntime        *int
			CpuShares           *int
			Cpus                *float64
			CpusetCpus          *string
			CpusetMems          *string
			Detach              bool
			DetachKeys          *string
			Device              []string
			DeviceCgroupRule    []string
			DeviceReadBps       []string
			DeviceReadIops      []string
			DeviceWriteBps      []string
			DeviceWriteIops     []string
			DisableContentTrust bool
			Dns                 []string
			DnsOption           []string
			DnsSearch           []string
			Entrypoint          *string
			Env                 map[string]string
			EnvFile             []string
			Expose              []string
			GroupAdd            []string
			HealthCmd           *string
			HealthInterval      *int
			HealthRetries       *int
			HealthStartPeriod   *int
			HealthTimeout       *int
			Hostname            *string
			Init                bool
			Interactive         bool
			Ip                  *string
			Ip6                 *string
			Ipc                 *string
			Isolation           *string
			KernelMemory        *int
			Label               map[string]string
			LabelFile           []string
			Link                []string
			LinkLocalIp         []string
			LogDriver           *string
			LogOpt              map[string]string
			MacAddress          *string
			Memory              *int
			MemoryReservation   *int
			MemorySwap          *int
			MemorySwappiness    *int
			Mount               map[string]string
			Name                *string
			Network             *string
			NetworkAlias        []string
			NoHealthcheck       bool
			OomKillDisable      bool
			OomScoreAdj         *int
			Pid                 *string
			PidsLimit           *int
			Platform            *string
			Privileged          bool
			Publish             []string
			PublishAll          bool
			ReadOnly            bool
			Restart             *string
			Rm                  bool
			Runtime             *string
			SecurityOpt         map[string]string
			ShmSize             *int
			SigProxy            bool
			StopSignal          *string
			StopTimeout         *int
			StorageOpt          map[string]string
			Sysctl              map[string]string
			Tmpfs               []string
			Tty                 bool
			Ulimit              map[string]string
			User                *string
			Userns              *string
			Uts                 *string
			Volume              []string
			VolumeDriver        *string
			VolumesFrom         []string
			Workdir             *string
		}
	}

}

type AliasesConf struct {
	PathMap map[string]*AliasConf
	Hash string
	Aliases []AliasConf
}

func LoadConfFile(ctx Context) (*AliasesConf, error) {
	if _, err := os.Stat(ctx.GetConfPath()); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file is not exists `%s`", ctx.GetConfPath())
	}

	buf, err := ioutil.ReadFile(ctx.GetConfPath())
	if err != nil {
		return nil, fmt.Errorf("configuration file cannot read `%q`", err)
	}

	defs, err:= UnmarshalConfFile(buf)
	if err != nil {
		return nil, err
	}

	conf := new(AliasesConf)
	conf.Hash = uuid.NewMD5(uuid.UUID{}, buf).String()
	conf.PathMap = make(map[string]*AliasConf)
	for path := range defs {
		conf.PathMap[path] = &AliasConf{}
	}

	for path, def := range defs {
		c := conf.PathMap[path]

		c.Path = path
		for _, dep := range def.Dependencies {
			if co, ok := conf.PathMap[dep]; ok {
				c.Dependencies = append(c.Dependencies, *co)
			} else {
				return nil, errors.New(fmt.Sprintf("undefined dependency: `%s` in `%s.Dependencies[]`", dep, path))
			}
		}

		c.DockerConf.Image = def.Image
		c.DockerConf.Tag = def.Tag
		c.DockerConf.Command = def.Command
		c.DockerConf.Args = def.Args

		c.DockerConf.DockerOpts.AddHost = expandColonDelimitedStringListWithEnv(def.AddHost)
		c.DockerConf.DockerOpts.Attach = expandColonDelimitedStringListWithEnv(def.Attach)
		c.DockerConf.DockerOpts.BlkioWeight = def.BlkioWeight
		c.DockerConf.DockerOpts.BlkioWeightDevice = expandColonDelimitedStringListWithEnv(def.BlkioWeightDevice)
		c.DockerConf.DockerOpts.CapAdd = def.CapAdd
		c.DockerConf.DockerOpts.CapDrop = def.CapDrop
		c.DockerConf.DockerOpts.CgroupParent = def.CgroupParent
		c.DockerConf.DockerOpts.Cidfile = def.Cidfile
		c.DockerConf.DockerOpts.CpuPeriod = def.CpuPeriod
		c.DockerConf.DockerOpts.CpuQuota = def.CpuQuota
		c.DockerConf.DockerOpts.CpuRtPeriod = def.CpuRtPeriod
		c.DockerConf.DockerOpts.CpuRtRuntime = def.CpuRtRuntime
		c.DockerConf.DockerOpts.CpuShares = def.CpuShares
		c.DockerConf.DockerOpts.Cpus = def.Cpus
		c.DockerConf.DockerOpts.CpusetCpus = def.CpusetCpus
		c.DockerConf.DockerOpts.CpusetMems = def.CpusetMems
		c.DockerConf.DockerOpts.Detach = def.Detach
		c.DockerConf.DockerOpts.DetachKeys = def.DetachKeys
		c.DockerConf.DockerOpts.Device = expandColonDelimitedStringListWithEnv(def.Device)
		c.DockerConf.DockerOpts.DeviceCgroupRule = def.DeviceCgroupRule
		c.DockerConf.DockerOpts.DeviceReadBps = expandColonDelimitedStringListWithEnv(def.DeviceReadBps)
		c.DockerConf.DockerOpts.DeviceReadIops = expandColonDelimitedStringListWithEnv(def.DeviceReadIops)
		c.DockerConf.DockerOpts.DeviceWriteBps = expandColonDelimitedStringListWithEnv(def.DeviceWriteBps)
		c.DockerConf.DockerOpts.DeviceWriteIops = expandColonDelimitedStringListWithEnv(def.DeviceWriteIops)
		c.DockerConf.DockerOpts.DisableContentTrust = def.DisableContentTrust
		c.DockerConf.DockerOpts.Dns = def.Dns
		c.DockerConf.DockerOpts.DnsOption = def.DnsOption
		c.DockerConf.DockerOpts.DnsSearch = def.DnsSearch
		c.DockerConf.DockerOpts.Entrypoint = def.Entrypoint
		c.DockerConf.DockerOpts.Env = expandStringKeyMapWithEnv(def.Env)
		c.DockerConf.DockerOpts.EnvFile = def.EnvFile
		c.DockerConf.DockerOpts.Expose = def.Expose
		c.DockerConf.DockerOpts.GroupAdd = def.GroupAdd
		c.DockerConf.DockerOpts.HealthCmd = def.HealthCmd
		c.DockerConf.DockerOpts.HealthInterval = def.HealthInterval
		c.DockerConf.DockerOpts.HealthRetries = def.HealthRetries
		c.DockerConf.DockerOpts.HealthStartPeriod = def.HealthStartPeriod
		c.DockerConf.DockerOpts.HealthTimeout = def.HealthTimeout
		c.DockerConf.DockerOpts.Hostname = def.Hostname
		c.DockerConf.DockerOpts.Init = def.Init
		c.DockerConf.DockerOpts.Interactive = def.Interactive
		c.DockerConf.DockerOpts.Ip = def.Ip
		c.DockerConf.DockerOpts.Ip6 = def.Ip6
		c.DockerConf.DockerOpts.Ipc = def.Ipc
		c.DockerConf.DockerOpts.Isolation = def.Isolation
		c.DockerConf.DockerOpts.KernelMemory = def.KernelMemory
		c.DockerConf.DockerOpts.Label = expandStringKeyMapWithEnv(def.Label)
		c.DockerConf.DockerOpts.LabelFile = def.LabelFile
		c.DockerConf.DockerOpts.Link = expandColonDelimitedStringListWithEnv(def.Link)
		c.DockerConf.DockerOpts.LinkLocalIp = def.LinkLocalIp
		c.DockerConf.DockerOpts.LogDriver = def.LogDriver
		c.DockerConf.DockerOpts.LogOpt = expandStringKeyMapWithEnv(def.LogOpt)
		c.DockerConf.DockerOpts.MacAddress = def.MacAddress
		c.DockerConf.DockerOpts.Memory = def.Memory
		c.DockerConf.DockerOpts.MemoryReservation = def.MemoryReservation
		c.DockerConf.DockerOpts.MemorySwap = def.MemorySwap
		c.DockerConf.DockerOpts.MemorySwappiness = def.MemorySwappiness
		c.DockerConf.DockerOpts.Mount = expandStringKeyMapWithEnv(def.Mount)
		c.DockerConf.DockerOpts.Name = def.Name
		c.DockerConf.DockerOpts.Network = def.Network
		c.DockerConf.DockerOpts.NetworkAlias = def.NetworkAlias
		c.DockerConf.DockerOpts.NoHealthcheck = def.NoHealthcheck
		c.DockerConf.DockerOpts.OomKillDisable = def.OomKillDisable
		c.DockerConf.DockerOpts.OomScoreAdj = def.OomScoreAdj
		c.DockerConf.DockerOpts.Pid = def.Pid
		c.DockerConf.DockerOpts.PidsLimit = def.PidsLimit
		c.DockerConf.DockerOpts.Platform = def.Platform
		c.DockerConf.DockerOpts.Privileged = def.Privileged
		c.DockerConf.DockerOpts.Publish = def.Publish
		c.DockerConf.DockerOpts.PublishAll = def.PublishAll
		c.DockerConf.DockerOpts.ReadOnly = def.ReadOnly
		c.DockerConf.DockerOpts.Restart = def.Restart
		c.DockerConf.DockerOpts.Rm = def.Rm
		c.DockerConf.DockerOpts.Runtime = def.Runtime
		c.DockerConf.DockerOpts.SecurityOpt = expandStringKeyMapWithEnv(def.SecurityOpt)
		c.DockerConf.DockerOpts.ShmSize = def.ShmSize
		c.DockerConf.DockerOpts.SigProxy = def.SigProxy
		c.DockerConf.DockerOpts.StopSignal = def.StopSignal
		c.DockerConf.DockerOpts.StopTimeout = def.StopTimeout
		c.DockerConf.DockerOpts.StorageOpt = expandStringKeyMapWithEnv(def.StorageOpt)
		c.DockerConf.DockerOpts.Sysctl = expandStringKeyMapWithEnv(def.Sysctl)
		c.DockerConf.DockerOpts.Tmpfs = def.Tmpfs
		c.DockerConf.DockerOpts.Tty = def.Tty
		c.DockerConf.DockerOpts.Ulimit = expandStringKeyMapWithEnv(def.Ulimit)
		c.DockerConf.DockerOpts.User = def.User
		c.DockerConf.DockerOpts.Userns = def.Userns
		c.DockerConf.DockerOpts.Uts = def.Uts
		c.DockerConf.DockerOpts.Volume = expandColonDelimitedStringListWithEnv(def.Volume)
		c.DockerConf.DockerOpts.VolumeDriver = def.VolumeDriver
		c.DockerConf.DockerOpts.VolumesFrom = def.VolumesFrom
		c.DockerConf.DockerOpts.Workdir = def.Workdir

		conf.Aliases = append(conf.Aliases, *c)
	}

	return conf, nil
}