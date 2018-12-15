package yaml

import (
	"gopkg.in/yaml.v2"
	"time"
)

type YamlDefinition struct {
	// aliases configuration
	Dependencies []string            `yaml:"dependencies"`
	// docker run options
	Detach *bool                     `yaml:"detach"`
	SigProxy *bool                   `yaml:"sig-proxy"`
	Name *string                     `yaml:"name"`
	DetachKeys *string               `yaml:"detach-keys"`
	Platform *string                 `yaml:"platform"`
	DisableContentTrust *bool        `yaml:"disable-content-trust"`
	Attach []string                  `yaml:"attach"`
	DeviceCgroupRule []string        `yaml:"device-cgroup-rule"`
	Device []string                  `yaml:"device"`
	Env map[string]string            `yaml:"env"`
	EnvFile []string                 `yaml:"env-file"`
	Entrypoint *string               `yaml:"entrypoint"`
	GroupAdd []string                `yaml:"group-add"`
	Hostname *string                 `yaml:"hostname"`
	Domainname *string               `yaml:"domainname"`
	Interactive *bool                `yaml:"interactive"`
	Label map[string]string          `yaml:"label"`
	LabelFile []string               `yaml:"label-file"`
	ReadOnly *bool                   `yaml:"read-only"`
	Restart *string                  `yaml:"restart"`
	StopSignal *string               `yaml:"stop-signal"`
	StopTimeout *int                 `yaml:"stop-timeout"`
	Sysctl map[string]string         `yaml:"sysctl"`
	Tty *bool                        `yaml:"tty"`
	Ulimit map[string]string         `yaml:"ulimit"`
	User *string                     `yaml:"user"`
	Workdir *string                  `yaml:"workdir"`
	Rm *bool                         `yaml:"rm"`
	CapAdd []string                  `yaml:"cap-add"`
	CapDrop []string                 `yaml:"cap-drop"`
	Privileged *bool                 `yaml:"privileged"`
	SecurityOpt map[string]string    `yaml:"security-opt"`
	Userns *string                   `yaml:"userns"`
	AddHost []string                 `yaml:"add-host"`
	Dns []string                     `yaml:"dns"`
	DnsOpt []string                  `yaml:"dns-opt"`
	DnsOption []string               `yaml:"dns-option"`
	DnsSearch []string               `yaml:"dns-search"`
	Expose []string                  `yaml:"expose"`
	Ip *string                       `yaml:"ip"`
	Ip6 *string                      `yaml:"ip6"`
	Link []string                    `yaml:"link"`
	LinkLocalIp []string             `yaml:"link-local-ip"`
	MacAddress *string               `yaml:"mac-address"`
	Publish []string                 `yaml:"publish"`
	PublishAll *bool                 `yaml:"publish-all"`
	Net *string                      `yaml:"net"`
	Network *string                  `yaml:"network"`
	NetAlias []string                `yaml:"net-alias"`
	NetworkAlias []string            `yaml:"network-alias"`
	LogDriver *string                `yaml:"log-driver"`
	VolumeDriver *string             `yaml:"volume-driver"`
	LogOpt map[string]string         `yaml:"log-opt"`
	StorageOpt map[string]string     `yaml:"storage-opt"`
	Tmpfs []string                   `yaml:"tmpfs"`
	VolumesFrom []string             `yaml:"volumes-from"`
	Volume []string                  `yaml:"volume"`
	Mount map[string]string          `yaml:"mount"`
	HealthCmd *string                `yaml:"health-cmd"`
	HealthInterval *time.Duration    `yaml:"health-interval"`
	HealthRetries *int               `yaml:"health-retries"`
	HealthTimeout *time.Duration     `yaml:"health-timeout"`
	HealthStartPeriod *time.Duration `yaml:"health-start-period"`
	NoHealthcheck *bool              `yaml:"no-healthcheck"`
	BlkioWeight *uint16              `yaml:"blkio-weight"`
	BlkioWeightDevice []string       `yaml:"blkio-weight-device"`
	Cidfile *string                  `yaml:"cidfile"`
	CpusetCpus *string               `yaml:"cpuset-cpus"`
	CpusetMems *string               `yaml:"cpuset-mems"`
	//CpuCount *int64                  `yaml:"cpu-count"`
	//CpuPercent *int64                `yaml:"cpu-percent"`
	CpuPeriod *int64                 `yaml:"cpu-period"`
	CpuQuota *int64                  `yaml:"cpu-quota"`
	CpuRtPeriod *int64               `yaml:"cpu-rt-period"`
	CpuRtRuntime *int64              `yaml:"cpu-rt-runtime"`
	CpuShares *int64                 `yaml:"cpu-shares"`
	Cpus *string                     `yaml:"cpus"`
	DeviceReadBps []string           `yaml:"device-read-bps"`
	DeviceReadIops []string          `yaml:"device-read-iops"`
	DeviceWriteBps []string          `yaml:"device-write-bps"`
	DeviceWriteIops []string         `yaml:"device-write-iops"`
	//IoMaxbandwidth *string           `yaml:"io-maxbandwidth"`
	//IoMaxiops *uint64                `yaml:"io-maxiops"`
	KernelMemory *string             `yaml:"kernel-memory"`
	Memory *string                   `yaml:"memory"`
	MemoryReservation *string        `yaml:"memory-reservation"`
	MemorySwap *string               `yaml:"memory-swap"`
	MemorySwappiness *int64          `yaml:"memory-swappiness"`
	OomKillDisable *bool             `yaml:"oom-kill-disable"`
	OomScoreAdj *int                 `yaml:"oom-score-adj"`
	PidsLimit *int64                 `yaml:"pids-limit"`
	CgroupParent *string             `yaml:"cgroup-parent"`
	Ipc *string                      `yaml:"ipc"`
	Isolation *string                `yaml:"isolation"`
	Pid *string                      `yaml:"pid"`
	ShmSize *string                  `yaml:"shm-size"`
	Uts *string                      `yaml:"uts"`
	Runtime *string                  `yaml:"runtime"`
	Init *bool                       `yaml:"init"`
	Image string                     `yaml:"image"`
	Args  []string                   `yaml:"args"`
	// extra docker run options
	Tag     string                   `yaml:"tag"`
	Command *string                  `yaml:"command"`
}

func UnmarshalConfFile(buf []byte) (map[string]YamlDefinition, error) {
	defs := make(map[string]YamlDefinition)
	if err := yaml.UnmarshalStrict(buf, &defs); err != nil {
		return nil, err
	}

	return defs, nil
}
