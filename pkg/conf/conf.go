package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/k-kinzal/aliases/pkg/docker"
	"github.com/k-kinzal/aliases/pkg/util"
	"github.com/k-kinzal/aliases/pkg/yaml"
)

type CommandConf struct {
	Path          string
	Dependencies  []*CommandConf
	DockerRunOpts docker.RunOpts
}

type AliasesConf struct {
	Commands []CommandConf
}

func LoadConfFile(ctx *context.Context) (*AliasesConf, error) {
	if _, err := os.Stat(ctx.GetConfPath()); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file is not exists `%s`", ctx.GetConfPath())
	}

	buf, err := ioutil.ReadFile(ctx.GetConfPath())
	if err != nil {
		return nil, fmt.Errorf("configuration file cannot read `%q`", err)
	}

	defs, err := yaml.UnmarshalConfFile(buf)
	if err != nil {
		return nil, err
	}

	conf := new(AliasesConf)

	pathMap := make(map[string]*CommandConf)
	for key := range defs {
		pathMap[key] = &CommandConf{Path: key}
	}

	for key, def := range defs {
		c := pathMap[key]

		c.Path = key
		if def.Image == "" {
			return nil, fmt.Errorf("image is required in %s.image", key)
		}
		if def.Tag == "" {
			return nil, fmt.Errorf("tag is required in %s.tag", key)
		}
		c.DockerRunOpts.Image = fmt.Sprintf("%s:${%s_VERSION:-\"%s\"}", def.Image, strings.ToUpper(path.Base(key)), def.Tag)
		if def.Command != nil {
			c.DockerRunOpts.Args = []string{*def.Command}
		}
		c.DockerRunOpts.Args = append(c.DockerRunOpts.Args, def.Args...)
		c.DockerRunOpts.AddHost = util.ExpandColonDelimitedStringListWithEnv(def.AddHost)
		c.DockerRunOpts.Attach = def.Attach
		c.DockerRunOpts.BlkioWeight = def.BlkioWeight
		c.DockerRunOpts.BlkioWeightDevice = util.ExpandColonDelimitedStringListWithEnv(def.BlkioWeightDevice)
		c.DockerRunOpts.CapAdd = def.CapAdd
		c.DockerRunOpts.CapDrop = def.CapDrop
		c.DockerRunOpts.CgroupParent = def.CgroupParent
		c.DockerRunOpts.Cidfile = def.Cidfile
		c.DockerRunOpts.CpuPeriod = def.CpuPeriod
		c.DockerRunOpts.CpuQuota = def.CpuQuota
		c.DockerRunOpts.CpuRtPeriod = def.CpuRtPeriod
		c.DockerRunOpts.CpuRtRuntime = def.CpuRtRuntime
		c.DockerRunOpts.CpuShares = def.CpuShares
		c.DockerRunOpts.Cpus = def.Cpus
		c.DockerRunOpts.CpusetCpus = def.CpusetCpus
		c.DockerRunOpts.CpusetMems = def.CpusetMems
		c.DockerRunOpts.Detach = def.Detach
		c.DockerRunOpts.DetachKeys = def.DetachKeys
		c.DockerRunOpts.Device = util.ExpandColonDelimitedStringListWithEnv(def.Device)
		c.DockerRunOpts.DeviceCgroupRule = def.DeviceCgroupRule
		c.DockerRunOpts.DeviceReadBps = util.ExpandColonDelimitedStringListWithEnv(def.DeviceReadBps)
		c.DockerRunOpts.DeviceReadIops = util.ExpandColonDelimitedStringListWithEnv(def.DeviceReadIops)
		c.DockerRunOpts.DeviceWriteBps = util.ExpandColonDelimitedStringListWithEnv(def.DeviceWriteBps)
		c.DockerRunOpts.DeviceWriteIops = util.ExpandColonDelimitedStringListWithEnv(def.DeviceWriteIops)
		c.DockerRunOpts.DisableContentTrust = def.DisableContentTrust
		c.DockerRunOpts.Dns = def.Dns
		c.DockerRunOpts.DnsOption = def.DnsOption
		c.DockerRunOpts.DnsSearch = def.DnsSearch
		c.DockerRunOpts.Entrypoint = def.Entrypoint
		c.DockerRunOpts.Env = util.ExpandStringKeyMapWithEnv(def.Env)
		c.DockerRunOpts.EnvFile = def.EnvFile
		c.DockerRunOpts.Expose = def.Expose
		c.DockerRunOpts.GroupAdd = def.GroupAdd
		c.DockerRunOpts.HealthCmd = def.HealthCmd
		c.DockerRunOpts.HealthInterval = def.HealthInterval
		c.DockerRunOpts.HealthRetries = def.HealthRetries
		c.DockerRunOpts.HealthStartPeriod = def.HealthStartPeriod
		c.DockerRunOpts.HealthTimeout = def.HealthTimeout
		c.DockerRunOpts.Hostname = def.Hostname
		c.DockerRunOpts.Init = def.Init
		c.DockerRunOpts.Interactive = def.Interactive
		c.DockerRunOpts.Ip = def.Ip
		c.DockerRunOpts.Ip6 = def.Ip6
		c.DockerRunOpts.Ipc = def.Ipc
		c.DockerRunOpts.Isolation = def.Isolation
		c.DockerRunOpts.KernelMemory = def.KernelMemory
		c.DockerRunOpts.Label = util.ExpandStringKeyMapWithEnv(def.Label)
		c.DockerRunOpts.LabelFile = def.LabelFile
		c.DockerRunOpts.Link = util.ExpandColonDelimitedStringListWithEnv(def.Link)
		c.DockerRunOpts.LinkLocalIp = def.LinkLocalIp
		c.DockerRunOpts.LogDriver = def.LogDriver
		c.DockerRunOpts.LogOpt = util.ExpandStringKeyMapWithEnv(def.LogOpt)
		c.DockerRunOpts.MacAddress = def.MacAddress
		c.DockerRunOpts.Memory = def.Memory
		c.DockerRunOpts.MemoryReservation = def.MemoryReservation
		c.DockerRunOpts.MemorySwap = def.MemorySwap
		c.DockerRunOpts.MemorySwappiness = def.MemorySwappiness
		c.DockerRunOpts.Mount = util.ExpandStringKeyMapWithEnv(def.Mount)
		c.DockerRunOpts.Name = def.Name
		c.DockerRunOpts.Network = def.Network
		c.DockerRunOpts.NetworkAlias = def.NetworkAlias
		c.DockerRunOpts.NoHealthcheck = def.NoHealthcheck
		c.DockerRunOpts.OomKillDisable = def.OomKillDisable
		c.DockerRunOpts.OomScoreAdj = def.OomScoreAdj
		c.DockerRunOpts.Pid = def.Pid
		c.DockerRunOpts.PidsLimit = def.PidsLimit
		c.DockerRunOpts.Platform = def.Platform
		c.DockerRunOpts.Privileged = def.Privileged
		c.DockerRunOpts.Publish = def.Publish
		c.DockerRunOpts.PublishAll = def.PublishAll
		c.DockerRunOpts.ReadOnly = def.ReadOnly
		c.DockerRunOpts.Restart = def.Restart
		c.DockerRunOpts.Rm = def.Rm
		c.DockerRunOpts.Runtime = def.Runtime
		c.DockerRunOpts.SecurityOpt = util.ExpandStringKeyMapWithEnv(def.SecurityOpt)
		c.DockerRunOpts.ShmSize = def.ShmSize
		c.DockerRunOpts.SigProxy = def.SigProxy
		c.DockerRunOpts.StopSignal = def.StopSignal
		c.DockerRunOpts.StopTimeout = def.StopTimeout
		c.DockerRunOpts.StorageOpt = util.ExpandStringKeyMapWithEnv(def.StorageOpt)
		c.DockerRunOpts.Sysctl = util.ExpandStringKeyMapWithEnv(def.Sysctl)
		c.DockerRunOpts.Tmpfs = def.Tmpfs
		c.DockerRunOpts.Tty = def.Tty
		c.DockerRunOpts.Ulimit = util.ExpandStringKeyMapWithEnv(def.Ulimit)
		if def.User != nil {
			c.DockerRunOpts.User = util.Pstr(util.ExpandColonDelimitedStringWithEnv(*def.User))
		}
		c.DockerRunOpts.Userns = def.Userns
		c.DockerRunOpts.Uts = def.Uts
		c.DockerRunOpts.Volume = util.ExpandColonDelimitedStringListWithEnv(def.Volume)
		c.DockerRunOpts.VolumeDriver = def.VolumeDriver
		c.DockerRunOpts.VolumesFrom = def.VolumesFrom
		c.DockerRunOpts.Workdir = def.Workdir

		if len(def.Dependencies) > 0 {
			for _, target := range def.Dependencies {
				if dep, ok := pathMap[target]; ok {
					c.Dependencies = append(c.Dependencies, dep)
				} else {
					return nil, errors.New(fmt.Sprintf("undefined dependency: `%s` in `%s.dependencies[]`", target, key))
				}
			}

			cmd := exec.Command("docker")
			if cmd.Path == "docker" {
				return nil, errors.New("docker is not installed. see https://docs.docker.com/install/")
			}
			// see: https://github.com/moby/moby/blob/bb1914b19572524b9f7d2b3415f146c545c1bb8b/client/client.go#L384
			host := os.Getenv("DOCKER_HOST")
			if host == "" {
				sock := "/var/run/docker.sock"
				if _, err := os.Stat(sock); err != nil {
					return nil, fmt.Errorf("%s: no such file. please set DOCKER_HOST", sock)
				}
				host = fmt.Sprintf("unix://%s", sock)
			}
			if strings.HasPrefix(host, "unix://") {
				sock := strings.TrimPrefix(host, "unix://")
				c.DockerRunOpts.Volume = append(c.DockerRunOpts.Volume, fmt.Sprintf("%s:/usr/local/bin/docker", cmd.Path))
				c.DockerRunOpts.Volume = append(c.DockerRunOpts.Volume, fmt.Sprintf("%s:/var/run/docker.sock", sock))
				c.DockerRunOpts.Privileged = util.Pstr("true")
			} else {
				if c.DockerRunOpts.Env == nil {
					c.DockerRunOpts.Env = make(map[string]string)
				}
				c.DockerRunOpts.Env["DOCKER_HOST"] = host
				c.DockerRunOpts.Network = util.Pstr("host")
			}

			c.DockerRunOpts.Env["ALIASES_PWD"] = "${ALIASES_PWD:-$PWD}"

			for _, dep := range c.Dependencies {
				from := fmt.Sprintf("%s/%s", ctx.GetExportPath(), path.Base(dep.Path))
				volume := fmt.Sprintf("%s:%s", from, dep.Path)
				c.DockerRunOpts.Volume = append(c.DockerRunOpts.Volume, volume)
			}
		}

		conf.Commands = append(conf.Commands, *c)
	}

	return conf, nil
}
