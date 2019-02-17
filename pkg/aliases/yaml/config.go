package yaml

import (
	"github.com/k-kinzal/aliases/pkg/types"
)

// OptionSpec is the specification of option in aliases config
//
// ```
// image: alpine
// tag: latest
// ...
// ```
type OptionSpec struct {
	// binary settings
	Docker struct {
		Image string `yaml:"image" default:"docker"`
		Tag   string `yaml:"tag" default:"18.09.0"`
	} `yaml:"docker"`
	// docker run settings
	Image   string   `yaml:"image" validate:"required"`
	Tag     string   `yaml:"tag" validate:"required"`
	Command *string  `yaml:"command"`
	Args    []string `yaml:"args"`
	// docker run option settings
	AddHost             []string          `yaml:"addHost"` // host:ip
	Attach              []string          `yaml:"attach" validate:"omitempty,dive,oneof=STDIN STDOUT STDERR"`
	BlkioWeight         *string           `yaml:"blkioWeight" validate:"omitempty,uint16|shell,_,shell|min=10,_,shell|max=1000"`
	BlkioWeightDevice   []string          `yaml:"blkioWeightDevice"`
	CIDFile             *string           `yaml:"cidfile"`
	CPUPeriod           *string           `yaml:"cpuPeriod" validate:"omitempty,nanocpus|shell"`
	CPUQuota            *string           `yaml:"cpuQuota" validate:"omitempty,int64|shell"`
	CPURtPeriod         *string           `yaml:"cpuRtPeriod" validate:"omitempty,int64|shell"`
	CPURtRuntime        *string           `yaml:"cpuRtRuntime" validate:"omitempty,int64|shell"`
	CPUShares           *string           `yaml:"cpuShares" validate:"omitempty,int64|shell"`
	CPUs                *string           `yaml:"cpus"`
	CPUsetCPUs          *string           `yaml:"cpusetCpus"`
	CPUsetMems          *string           `yaml:"cpusetMems"`
	CapAdd              []string          `yaml:"capAdd"`
	CapDrop             []string          `yaml:"capDrop"`
	CgroupParent        *string           `yaml:"cgroupParent"`
	DNS                 []string          `yaml:"dns"`
	DNSOption           []string          `yaml:"dnsOption"`
	DNSSearch           []string          `yaml:"dnsSearch"`
	Detach              *string           `yaml:"detach" validate:"omitempty,bool|shell"`
	DetachKeys          *string           `yaml:"detachKeys"`
	Device              []string          `yaml:"device"`
	DeviceCgroupRule    []string          `yaml:"deviceCgroupRule"`
	DeviceReadBPS       []string          `yaml:"deviceReadBps"`
	DeviceReadIOPS      []string          `yaml:"deviceReadIops"`
	DeviceWriteBPS      []string          `yaml:"deviceWriteBps"`
	DeviceWriteIOPS     []string          `yaml:"deviceWriteIops"`
	DisableContentTrust *string           `yaml:"disableContentTrust" validate:"omitempty,bool|shell"`
	Domainname          *string           `yaml:"domainname"`
	Entrypoint          *string           `yaml:"entrypoint"`
	Env                 map[string]string `yaml:"env"`
	EnvFile             []string          `yaml:"shellFile"`
	Expose              []string          `yaml:"expose"`
	GroupAdd            []string          `yaml:"groupAdd"`
	HealthCmd           *string           `yaml:"healthCmd"`
	HealthInterval      *string           `yaml:"healthInterval" validate:"omitempty,duration|shell"`
	HealthRetries       *string           `yaml:"healthRetries" validate:"omitempty,int|shell"`
	HealthStartPeriod   *string           `yaml:"healthStartPeriod" validate:"omitempty,duration|shell"`
	HealthTimeout       *string           `yaml:"healthTimeout" validate:"omitempty,duration|shell"`
	Hostname            *string           `yaml:"hostname"`
	IP                  *string           `yaml:"ip" validate:"omitempty,ipv4|shell"`
	IP6                 *string           `yaml:"ip6" validate:"omitempty,ipv6|shell"`
	IPC                 *string           `yaml:"ipc"`
	Init                *string           `yaml:"init" validate:"omitempty,bool|shell"`
	Interactive         *string           `yaml:"interactive" validate:"omitempty,bool|shell" default:"true"`
	Isolation           *string           `yaml:"isolation"`
	KernelMemory        *string           `yaml:"kernelMemory" validate:"omitempty,membytes|shell"`
	Label               map[string]string `yaml:"label"`
	LabelFile           []string          `yaml:"labelFile"`
	Link                []string          `yaml:"link"`
	LinkLocalIP         []string          `yaml:"linkLocalIp"`
	LogDriver           *string           `yaml:"logDriver"`
	LogOpt              map[string]string `yaml:"logOpt"`
	MacAddress          *string           `yaml:"macAddress" validate:"omitempty,mac|shell"`
	Memory              *string           `yaml:"memory" validate:"omitempty,membytes|shell"`
	MemoryReservation   *string           `yaml:"memoryReservation" validate:"omitempty,membytes|shell"`
	MemorySwap          *string           `yaml:"memorySwap" validate:"omitempty,memswapbytes|shell"`
	MemorySwappiness    *string           `yaml:"memorySwappiness" validate:"omitempty,int64|shell,_,shell|min=-1,_,shell|max=100"`
	Mount               map[string]string `yaml:"mount"`
	Name                *string           `yaml:"name"`
	Network             *string           `yaml:"network" default:"host"`
	NetworkAlias        []string          `yaml:"networkAlias"`
	NoHealthcheck       *string           `yaml:"noHealthcheck" validate:"omitempty,bool|shell"`
	OOMKillDisable      *string           `yaml:"oomKillDisable" validate:"omitempty,bool|shell"`
	OOMScoreAdj         *string           `yaml:"oomScoreAdj" validate:"omitempty,int|shell,_,shell|min=-1000,_,shell|max=1000"`
	PID                 *string           `yaml:"pid"`
	PidsLimit           *string           `yaml:"pidsLimit" validate:"omitempty,int64|shell"`
	Platform            *string           `yaml:"platform"`
	Privileged          *string           `yaml:"privileged" validate:"omitempty,bool|shell"`
	Publish             []string          `yaml:"publish"`
	PublishAll          *string           `yaml:"publishAll" validate:"omitempty,bool|shell"`
	ReadOnly            *string           `yaml:"readOnly" validate:"omitempty,bool|shell"`
	Restart             *string           `yaml:"restart"`
	Rm                  *string           `yaml:"rm" validate:"omitempty,bool|shell" default:"true"`
	Runtime             *string           `yaml:"runtime"`
	SecurityOpt         map[string]string `yaml:"securityOpt"`
	ShmSize             *string           `yaml:"shmSize" validate:"omitempty,membytes|shell"`
	SigProxy            *string           `yaml:"sigProxy" validate:"omitempty,bool|shell"`
	StopSignal          *string           `yaml:"stopSignal"`
	StopTimeout         *string           `yaml:"stopTimeout" validate:"omitempty,int|shell"`
	StorageOpt          map[string]string `yaml:"storageOpt"`
	Sysctl              map[string]string `yaml:"sysctl"`
	TTY                 *string           `yaml:"tty" validate:"omitempty,bool|shell" default:"tty >/dev/null"`
	Tmpfs               []string          `yaml:"tmpfs"`
	UTS                 *string           `yaml:"uts"`
	Ulimit              map[string]string `yaml:"ulimit"`
	User                *string           `yaml:"user"` // <name|uid>[:<group|gid>]
	Userns              *string           `yaml:"userns"`
	Volume              []string          `yaml:"volume"`
	VolumeDriver        *string           `yaml:"volumeDriver"`
	VolumesFrom         []string          `yaml:"volumesFrom"`
	Workdir             *string           `yaml:"workdir"`
	// dependency
	Dependencies []DependencySpec `yaml:"dependencies"`
}

// ConfigSpec is the specification of aliases config
//
// ```
// /path/to/command:
//   image: alpine
//   tag: latest
//   ...
// ```
type ConfigSpec map[string]OptionSpec

// BreadthWalk is executes the function when entry traverse OptionSpec.
func (spec *ConfigSpec) BreadthWalk(fn func(path SpecPath, current OptionSpec) (*OptionSpec, error)) error {
	type value struct {
		path    SpecPath
		current *OptionSpec
		parent  *OptionSpec
	}
	queue := make([]value, 0)
	for key, opt := range *spec {
		o := opt
		queue = append(queue, value{SpecPath(key), &o, nil})
	}
	for i := 0; i < len(queue); i++ {
		v := queue[i]

		current, err := fn(v.path, *v.current)
		if err != nil {
			return err
		}
		if v.parent == nil {
			(*spec)[v.path.String()] = *current
		} else {
			conf := v.parent.Dependencies[v.path.Index()].Config()
			conf[v.path.Name()] = *current
			v.parent.Dependencies[v.path.Index()] = *NewDependencySpec(conf)
		}

		for index, dependency := range v.current.Dependencies {
			if dependency.IsString() {
				continue
			}
			for key, opt := range dependency.Config() {
				o := opt
				queue = append(queue, value{*v.path.Dependencies(index, key), &o, v.current})
			}
		}
	}

	return nil
}

// Walk is executes the function when leaving traverse OptionSpec.
func (spec *ConfigSpec) DepthWalk(fn func(path SpecPath, current OptionSpec) (*OptionSpec, error)) error {
	type value struct {
		path    SpecPath
		current *OptionSpec
		parent  *OptionSpec
	}

	for key, opt := range *spec {
		stack := types.NewStack(nil)
		callstack := types.NewStack(nil)
		callstack.Push(value{SpecPath(key), &opt, nil})
		for val := callstack.Pop(); val != nil; val = callstack.Pop() {
			v := val.(value)
			for i, dependency := range v.current.Dependencies {
				if dependency.IsString() {
					continue
				}
				for k, c := range dependency.Config() {
					cmd := c
					callstack.Push(value{*v.path.Dependencies(i, k), &cmd, v.current})
				}
			}
			stack.Push(val)
		}
		for _, val := range stack.Slice() {
			v := val.(value)
			current, err := fn(v.path, *v.current)
			if err != nil {
				return err
			}
			if v.parent == nil {
				(*spec)[v.path.String()] = *current
			} else {
				conf := v.parent.Dependencies[v.path.Index()].Config()
				conf[v.path.Name()] = *current
				v.parent.Dependencies[v.path.Index()] = *NewDependencySpec(conf)
			}
		}
	}

	return nil
}
