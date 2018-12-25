package yaml

import (
	"fmt"

	"github.com/creasty/defaults"
	"github.com/k-kinzal/aliases/pkg/validator"
	yaml "gopkg.in/yaml.v2"
)

type Schema struct {
	// aliases configuration
	Dependencies []string `yaml:"dependencies"`
	// docker run options
	Detach              *string           `yaml:"detach" validate:"omitempty,bool|script"`
	SigProxy            *string           `yaml:"sig-proxy" validate:"omitempty,bool|script"`
	Name                *string           `yaml:"name"`
	DetachKeys          *string           `yaml:"detach-keys"`
	Platform            *string           `yaml:"platform"`
	DisableContentTrust *string           `yaml:"disable-content-trust" validate:"omitempty,bool|script"`
	Attach              []string          `yaml:"attach"`
	DeviceCgroupRule    []string          `yaml:"device-cgroup-rule"`
	Device              []string          `yaml:"device"`
	Env                 map[string]string `yaml:"env"`
	EnvFile             []string          `yaml:"env-file"`
	Entrypoint          *string           `yaml:"entrypoint"`
	GroupAdd            []string          `yaml:"group-add"`
	Hostname            *string           `yaml:"hostname"`
	Domainname          *string           `yaml:"domainname"`
	Interactive         *string           `yaml:"interactive" validate:"omitempty,bool|script" default:"true"`
	Label               map[string]string `yaml:"label"`
	LabelFile           []string          `yaml:"label-file"`
	ReadOnly            *string           `yaml:"read-only" validate:"omitempty,bool|script"`
	Restart             *string           `yaml:"restart"`
	StopSignal          *string           `yaml:"stop-signal"`
	StopTimeout         *string           `yaml:"stop-timeout" validate:"omitempty,int|script"`
	Sysctl              map[string]string `yaml:"sysctl"`
	Tty                 *string           `yaml:"tty" validate:"omitempty,bool|script"`
	Ulimit              map[string]string `yaml:"ulimit"`
	User                *string           `yaml:"user"`
	Workdir             *string           `yaml:"workdir"`
	Rm                  *string           `yaml:"rm" validate:"omitempty,bool|script" default:"true"`
	CapAdd              []string          `yaml:"cap-add"`
	CapDrop             []string          `yaml:"cap-drop"`
	Privileged          *string           `yaml:"privileged" validate:"omitempty,bool|script"`
	SecurityOpt         map[string]string `yaml:"security-opt"`
	Userns              *string           `yaml:"userns"`
	AddHost             []string          `yaml:"add-host"`
	Dns                 []string          `yaml:"dns"`
	DnsOpt              []string          `yaml:"dns-opt"`
	DnsOption           []string          `yaml:"dns-option"`
	DnsSearch           []string          `yaml:"dns-search"`
	Expose              []string          `yaml:"expose"`
	Ip                  *string           `yaml:"ip"`
	Ip6                 *string           `yaml:"ip6"`
	Link                []string          `yaml:"link"`
	LinkLocalIp         []string          `yaml:"link-local-ip"`
	MacAddress          *string           `yaml:"mac-address"`
	Publish             []string          `yaml:"publish"`
	PublishAll          *string           `yaml:"publish-all" validate:"omitempty,bool|script"`
	Network             *string           `yaml:"network" default:"host"`
	NetworkAlias        []string          `yaml:"network-alias"`
	LogDriver           *string           `yaml:"log-driver"`
	VolumeDriver        *string           `yaml:"volume-driver"`
	LogOpt              map[string]string `yaml:"log-opt"`
	StorageOpt          map[string]string `yaml:"storage-opt"`
	Tmpfs               []string          `yaml:"tmpfs"`
	VolumesFrom         []string          `yaml:"volumes-from"`
	Volume              []string          `yaml:"volume"`
	Mount               map[string]string `yaml:"mount"`
	HealthCmd           *string           `yaml:"health-cmd"`
	HealthInterval      *string           `yaml:"health-interval" validate:"omitempty,duration|script"`
	HealthRetries       *string           `yaml:"health-retries" validate:"omitempty,int|script"`
	HealthTimeout       *string           `yaml:"health-timeout" validate:"omitempty,duration|script"`
	HealthStartPeriod   *string           `yaml:"health-start-period" validate:"omitempty,duration|script"`
	NoHealthcheck       *string           `yaml:"no-healthcheck" validate:"omitempty,bool|script"`
	BlkioWeight         *string           `yaml:"blkio-weight" validate:"omitempty,uint16|script"`
	BlkioWeightDevice   []string          `yaml:"blkio-weight-device"`
	Cidfile             *string           `yaml:"cidfile"`
	CpusetCpus          *string           `yaml:"cpuset-cpus"`
	CpusetMems          *string           `yaml:"cpuset-mems"`
	CpuPeriod           *string           `yaml:"cpu-period" validate:"omitempty,nanocpus|script"`
	CpuQuota            *string           `yaml:"cpu-quota" validate:"omitempty,int64|script"`
	CpuRtPeriod         *string           `yaml:"cpu-rt-period" validate:"omitempty,int64|script"`
	CpuRtRuntime        *string           `yaml:"cpu-rt-runtime" validate:"omitempty,int64|script"`
	CpuShares           *string           `yaml:"cpu-shares" validate:"omitempty,int64|script"`
	Cpus                *string           `yaml:"cpus"`
	DeviceReadBps       []string          `yaml:"device-read-bps"`
	DeviceReadIops      []string          `yaml:"device-read-iops"`
	DeviceWriteBps      []string          `yaml:"device-write-bps"`
	DeviceWriteIops     []string          `yaml:"device-write-iops"`
	KernelMemory        *string           `yaml:"kernel-memory" validate:"omitempty,membytes|script"`
	Memory              *string           `yaml:"memory" validate:"omitempty,membytes|script"`
	MemoryReservation   *string           `yaml:"memory-reservation" validate:"omitempty,membytes|script"`
	MemorySwap          *string           `yaml:"memory-swap" validate:"omitempty,memswapbytes|script"`
	MemorySwappiness    *string           `yaml:"memory-swappiness" validate:"omitempty,int64|script"`
	OomKillDisable      *string           `yaml:"oom-kill-disable" validate:"omitempty,bool|script"`
	OomScoreAdj         *string           `yaml:"oom-score-adj" validate:"omitempty,int|script"`
	PidsLimit           *string           `yaml:"pids-limit" validate:"omitempty,int64|script"`
	CgroupParent        *string           `yaml:"cgroup-parent"`
	Ipc                 *string           `yaml:"ipc"`
	Isolation           *string           `yaml:"isolation"`
	Pid                 *string           `yaml:"pid"`
	ShmSize             *string           `yaml:"shm-size" validate:"omitempty,membytes|script"`
	Uts                 *string           `yaml:"uts"`
	Runtime             *string           `yaml:"runtime"`
	Init                *string           `yaml:"init" validate:"omitempty,bool|script"`
	Image               string            `yaml:"image" validate:required"`
	Args                []string          `yaml:"args"`
	// extra docker run options
	Tag     string  `yaml:"tag" validate:required"`
	Command *string `yaml:"command"`
}

func UnmarshalConfFile(buf []byte) (map[string]Schema, error) {
	schemas := make(map[string]Schema)
	if err := yaml.UnmarshalStrict(buf, &schemas); err != nil {
		return nil, err
	}

	validate := validator.New()
	for path, schema := range schemas {
		if err := defaults.Set(&schema); err != nil {
			return nil, err
		}
		if err := validate.Struct(schema); err != nil {
			return nil, err // FIXME: add key to error message `xxx in /path/to/cmd.interval`
		}
		for index, dep := range schema.Dependencies {
			if _, ok := schemas[dep]; !ok {
				return nil, fmt.Errorf("`%s` is an undefined dependency in %s.dependencies[%d]", path, dep, index)
			}

		}
		schemas[path] = schema
	}

	return schemas, nil
}
