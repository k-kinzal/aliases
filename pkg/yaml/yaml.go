package yaml

import (
	"fmt"
	"strings"

	"github.com/creasty/defaults"
	"github.com/k-kinzal/aliases/pkg/validator"
	yaml "gopkg.in/yaml.v2"
)

type Schema struct {
	// aliases configuration
	Dependencies []string `yaml:"dependencies"`
	// docker run options
	Detach              *string           `yaml:"detach" validate:"omitempty,bool|script"`
	SigProxy            *string           `yaml:"sigProxy" validate:"omitempty,bool|script"`
	Name                *string           `yaml:"name"`
	DetachKeys          *string           `yaml:"detachKeys"`
	Platform            *string           `yaml:"platform"`
	DisableContentTrust *string           `yaml:"disableContentTrust" validate:"omitempty,bool|script"`
	Attach              []string          `yaml:"attach"`
	DeviceCgroupRule    []string          `yaml:"deviceCgroupRule"`
	Device              []string          `yaml:"device"`
	Env                 map[string]string `yaml:"env"`
	EnvFile             []string          `yaml:"envFile"`
	Entrypoint          *string           `yaml:"entrypoint"`
	GroupAdd            []string          `yaml:"groupAdd"`
	Hostname            *string           `yaml:"hostname"`
	Domainname          *string           `yaml:"domainname"`
	Interactive         *string           `yaml:"interactive" validate:"omitempty,bool|script" default:"true"`
	Label               map[string]string `yaml:"label"`
	LabelFile           []string          `yaml:"labelFile"`
	ReadOnly            *string           `yaml:"readOnly" validate:"omitempty,bool|script"`
	Restart             *string           `yaml:"restart"`
	StopSignal          *string           `yaml:"stopSignal"`
	StopTimeout         *string           `yaml:"stopTimeout" validate:"omitempty,int|script"`
	Sysctl              map[string]string `yaml:"sysctl"`
	Tty                 *string           `yaml:"tty" validate:"omitempty,bool|script"`
	Ulimit              map[string]string `yaml:"ulimit"`
	User                *string           `yaml:"user"`
	Workdir             *string           `yaml:"workdir"`
	Rm                  *string           `yaml:"rm" validate:"omitempty,bool|script" default:"true"`
	CapAdd              []string          `yaml:"capAdd"`
	CapDrop             []string          `yaml:"capDrop"`
	Privileged          *string           `yaml:"privileged" validate:"omitempty,bool|script"`
	SecurityOpt         map[string]string `yaml:"securityOpt"`
	Userns              *string           `yaml:"userns"`
	AddHost             []string          `yaml:"addHost"`
	Dns                 []string          `yaml:"dns"`
	DnsOpt              []string          `yaml:"dnsOpt"`
	DnsOption           []string          `yaml:"dnsOption"`
	DnsSearch           []string          `yaml:"dnsSearch"`
	Expose              []string          `yaml:"expose"`
	Ip                  *string           `yaml:"ip"`
	Ip6                 *string           `yaml:"ip6"`
	Link                []string          `yaml:"link"`
	LinkLocalIp         []string          `yaml:"linkLocalIp"`
	MacAddress          *string           `yaml:"macAddress"`
	Publish             []string          `yaml:"publish"`
	PublishAll          *string           `yaml:"publishAll" validate:"omitempty,bool|script"`
	Network             *string           `yaml:"network" default:"host"`
	NetworkAlias        []string          `yaml:"networkAlias"`
	LogDriver           *string           `yaml:"logDriver"`
	VolumeDriver        *string           `yaml:"volumeDriver"`
	LogOpt              map[string]string `yaml:"logOpt"`
	StorageOpt          map[string]string `yaml:"storageOpt"`
	Tmpfs               []string          `yaml:"tmpfs"`
	VolumesFrom         []string          `yaml:"volumesFrom"`
	Volume              []string          `yaml:"volume"`
	Mount               map[string]string `yaml:"mount"`
	HealthCmd           *string           `yaml:"healthCmd"`
	HealthInterval      *string           `yaml:"healthInterval" validate:"omitempty,duration|script"`
	HealthRetries       *string           `yaml:"healthRetries" validate:"omitempty,int|script"`
	HealthTimeout       *string           `yaml:"healthTimeout" validate:"omitempty,duration|script"`
	HealthStartPeriod   *string           `yaml:"healthStartPeriod" validate:"omitempty,duration|script"`
	NoHealthcheck       *string           `yaml:"noHealthcheck" validate:"omitempty,bool|script"`
	BlkioWeight         *string           `yaml:"blkioWeight" validate:"omitempty,uint16|script"`
	BlkioWeightDevice   []string          `yaml:"blkioWeightDevice"`
	Cidfile             *string           `yaml:"cidfile"`
	CpusetCpus          *string           `yaml:"cpusetCpus"`
	CpusetMems          *string           `yaml:"cpusetMems"`
	CpuPeriod           *string           `yaml:"cpuPeriod" validate:"omitempty,nanocpus|script"`
	CpuQuota            *string           `yaml:"cpuQuota" validate:"omitempty,int64|script"`
	CpuRtPeriod         *string           `yaml:"cpuRtPeriod" validate:"omitempty,int64|script"`
	CpuRtRuntime        *string           `yaml:"cpuRtRuntime" validate:"omitempty,int64|script"`
	CpuShares           *string           `yaml:"cpuShares" validate:"omitempty,int64|script"`
	Cpus                *string           `yaml:"cpus"`
	DeviceReadBps       []string          `yaml:"deviceReadBps"`
	DeviceReadIops      []string          `yaml:"deviceReadIops"`
	DeviceWriteBps      []string          `yaml:"deviceWriteBps"`
	DeviceWriteIops     []string          `yaml:"deviceWriteIops"`
	KernelMemory        *string           `yaml:"kernelMemory" validate:"omitempty,membytes|script"`
	Memory              *string           `yaml:"memory" validate:"omitempty,membytes|script"`
	MemoryReservation   *string           `yaml:"memoryReservation" validate:"omitempty,membytes|script"`
	MemorySwap          *string           `yaml:"memorySwap" validate:"omitempty,memswapbytes|script"`
	MemorySwappiness    *string           `yaml:"memorySwappiness" validate:"omitempty,int64|script"`
	OomKillDisable      *string           `yaml:"oomKillDisable" validate:"omitempty,bool|script"`
	OomScoreAdj         *string           `yaml:"oomScoreAdj" validate:"omitempty,int|script"`
	PidsLimit           *string           `yaml:"pidsLimit" validate:"omitempty,int64|script"`
	CgroupParent        *string           `yaml:"cgroupParent"`
	Ipc                 *string           `yaml:"ipc"`
	Isolation           *string           `yaml:"isolation"`
	Pid                 *string           `yaml:"pid"`
	ShmSize             *string           `yaml:"shmSize" validate:"omitempty,membytes|script"`
	Uts                 *string           `yaml:"uts"`
	Runtime             *string           `yaml:"runtime"`
	Init                *string           `yaml:"init" validate:"omitempty,bool|script"`
	Image               string            `yaml:"image" validate:"required"`
	Args                []string          `yaml:"args"`
	// extra docker run options
	Tag     string  `yaml:"tag" validate:"required"`
	Command *string `yaml:"command"`
}

func UnmarshalConfFile(buf []byte) (map[string]Schema, error) {
	schemas := make(map[string]Schema)
	if err := yaml.UnmarshalStrict(buf, &schemas); err != nil {
		if e, ok := err.(*yaml.TypeError); ok {
			return nil, fmt.Errorf("yaml error: %s", strings.Replace(e.Errors[0], "in type yaml.Schema", "", 1))
		}
		return nil, fmt.Errorf("runtime error: %s", err)
	}

	validate := validator.New()
	for path, schema := range schemas {
		if err := defaults.Set(&schema); err != nil {
			return nil, err
		}
		if err := validate.Struct(schema); err != nil {
			return nil, fmt.Errorf("yaml error: %s in `%s`", err, path)
		}
		for index, dep := range schema.Dependencies {
			if _, ok := schemas[dep]; !ok {
				return nil, fmt.Errorf("yaml error: invalid parameter `%s` for `dependencies[%d]` is an undefined dependency in `%s`", dep, index, path)
			}

		}
		schemas[path] = schema
	}

	return schemas, nil
}
