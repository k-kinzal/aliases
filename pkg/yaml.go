package aliases

import (
	"github.com/creasty/defaults"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

type YamlDefinition struct {
	// aliases ciguration
	Dependencies []string `yaml:"dependencies"`
	// docker run parameters
	Image   string   `yaml:"image" validate:"required"`
	Tag     string   `yaml:"tag" validate:"required"`
	Command *string  `yaml:"command"`
	Args    []string `yaml:"args"`
	// docker run options
	AddHost             []string          `yaml:"add-host"`
	Attach              []string          `yaml:"attach" validate:"dive,oneof=STDIN STDOUT STDERR"`
	BlkioWeight         *uint16           `yaml:"blkio-weight"`
	BlkioWeightDevice   []string          `yaml:"blkio-weight-device"`
	CapAdd              []string          `yaml:"cap-add"`
	CapDrop             []string          `yaml:"cap-drop"`
	CgroupParent        *string           `yaml:"cgroup-parent"`
	Cidfile             *string           `yaml:"cidfile"`
	CpuPeriod           *int              `yaml:"cpu-period"`
	CpuQuota            *int              `yaml:"cpu-quota"`
	CpuRtPeriod         *int              `yaml:"cpu-rt-period"`
	CpuRtRuntime        *int              `yaml:"cpu-rt-runtime"`
	CpuShares           *int              `yaml:"cpu-shares"`
	Cpus                *float64          `yaml:"cpus"`
	CpusetCpus          *string           `yaml:"cpuset-cpus"`
	CpusetMems          *string           `yaml:"cpuset-mems"`
	Detach              bool              `yaml:"detach"`
	DetachKeys          *string           `yaml:"detach-keys"`
	Device              []string          `yaml:"device"`
	DeviceCgroupRule    []string          `yaml:"device-cgroup-rule"` //FIXME: check parameter
	DeviceReadBps       []string          `yaml:"device-read-bps"`
	DeviceReadIops      []string          `yaml:"device-read-iops"`
	DeviceWriteBps      []string          `yaml:"device-write-bps"`
	DeviceWriteIops     []string          `yaml:"device-write-iops"`
	DisableContentTrust bool              `yaml:"disable-content-trust"`
	Dns                 []string          `yaml:"dns"`
	DnsOption           []string          `yaml:"dns-option"` //FIXME: check parameter
	DnsSearch           []string          `yaml:"dns-search"` //FIXME: check parameter
	Entrypoint          *string           `yaml:"entrypoint"`
	Env                 map[string]string `yaml:"env"`
	EnvFile             []string          `yaml:"env-file"`
	Expose              []string          `yaml:"expose"`
	GroupAdd            []string          `yaml:"group-add"`
	HealthCmd           *string           `yaml:"health-cmd"`
	HealthInterval      *int              `yaml:"health-interval"`
	HealthRetries       *int              `yaml:"health-retries"`
	HealthStartPeriod   *int              `yaml:"health-start-period"`
	HealthTimeout       *int              `yaml:"health-timeout"`
	Hostname            *string           `yaml:"hostname"`
	Init                bool              `yaml:"init"`
	Interactive         bool              `yaml:"interactive" default:"true"`
	Ip                  *string           `yaml:"ip"`
	Ip6                 *string           `yaml:"ip6"`
	Ipc                 *string           `yaml:"ipc"`
	Isolation           *string           `yaml:"isolation"`
	KernelMemory        *int              `yaml:"kernel-memory"`
	Label               map[string]string `yaml:"label"`
	LabelFile           []string          `yaml:"label-file"`
	Link                []string          `yaml:"link"`
	LinkLocalIp         []string          `yaml:"link-local-ip"`
	LogDriver           *string           `yaml:"log-driver"`
	LogOpt              map[string]string `yaml:"log-opt"`
	MacAddress          *string           `yaml:"mac-address"`
	Memory              *int              `yaml:"memory"`
	MemoryReservation   *int              `yaml:"memory-reservation"`
	MemorySwap          *int              `yaml:"memory-swap"`
	MemorySwappiness    *int              `yaml:"memory-swappiness"`
	Mount               map[string]string `yaml:"mount"`
	Name                *string           `yaml:"name"`
	Network             *string           `yaml:"network" default:"host"`
	NetworkAlias        []string          `yaml:"network-alias"`
	NoHealthcheck       bool              `yaml:"no-healthcheck"`
	OomKillDisable      bool              `yaml:"oom-kill-disable"`
	OomScoreAdj         *int              `yaml:"oom-score-adj"`
	Pid                 *string           `yaml:"pid"`
	PidsLimit           *int              `yaml:"pids-limit"`
	Platform            *string           `yaml:"platform"`
	Privileged          bool              `yaml:"privileged"`
	Publish             []string          `yaml:"publish"`
	PublishAll          bool              `yaml:"publish-all"`
	ReadOnly            bool              `yaml:"read-only"`
	Restart             *string           `yaml:"restart"`
	Rm                  bool              `yaml:"rm" default:"true"`
	Runtime             *string           `yaml:"runtime"`
	SecurityOpt         map[string]string `yaml:"security-opt"`
	ShmSize             *int              `yaml:"shm-size"`
	SigProxy            bool              `yaml:"sig-proxy"`
	StopSignal          *string           `yaml:"stop-signal"`
	StopTimeout         *int              `yaml:"stop-timeout"`
	StorageOpt          map[string]string `yaml:"storage-opt"`
	Sysctl              map[string]string `yaml:"sysctl"`
	Tmpfs               []string          `yaml:"tmpfs"`
	Tty                 bool              `yaml:"tty"`
	Ulimit              map[string]string `yaml:"ulimit"`
	User                *string           `yaml:"user"`
	Userns              *string           `yaml:"userns"`
	Uts                 *string           `yaml:"uts"`
	Volume              []string          `yaml:"volume"`
	VolumeDriver        *string           `yaml:"volume-driver"`
	VolumesFrom         []string          `yaml:"volumes-from"`
	Workdir             *string           `yaml:"workdir"`
}

func UnmarshalConfFile(buf []byte) (map[string]YamlDefinition, error) {
	defs := make(map[string]YamlDefinition)
	if err := yaml.Unmarshal(buf, &defs); err != nil {
		return nil, err
	}
	for idx, def := range defs {
		if err := defaults.Set(&def); err != nil {
			return nil, err
		}
		if err := validator.New().Struct(def); err != nil {
			return nil, err
		}
		defs[idx] = def
	}

	return defs, nil
}
