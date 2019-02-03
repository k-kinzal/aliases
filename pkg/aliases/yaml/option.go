package yaml

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
	BlkioWeight         *string           `yaml:"blkioWeight" validate:"omitempty,uint16|env,_,env|min=10,_,env|max=1000"`
	BlkioWeightDevice   []string          `yaml:"blkioWeightDevice"`
	CIDFile             *string           `yaml:"cidfile"`
	CPUPeriod           *string           `yaml:"cpuPeriod" validate:"omitempty,nanocpus|env"`
	CPUQuota            *string           `yaml:"cpuQuota" validate:"omitempty,int64|env"`
	CPURtPeriod         *string           `yaml:"cpuRtPeriod" validate:"omitempty,int64|env"`
	CPURtRuntime        *string           `yaml:"cpuRtRuntime" validate:"omitempty,int64|env"`
	CPUShares           *string           `yaml:"cpuShares" validate:"omitempty,int64|env"`
	CPUs                *string           `yaml:"cpus"`
	CPUsetCPUs          *string           `yaml:"cpusetCpus"`
	CPUsetMems          *string           `yaml:"cpusetMems"`
	CapAdd              []string          `yaml:"capAdd"`
	CapDrop             []string          `yaml:"capDrop"`
	CgroupParent        *string           `yaml:"cgroupParent"`
	DNS                 []string          `yaml:"dns"`
	DNSOption           []string          `yaml:"dnsOption"`
	DNSSearch           []string          `yaml:"dnsSearch"`
	Detach              *string           `yaml:"detach" validate:"omitempty,bool|env"`
	DetachKeys          *string           `yaml:"detachKeys"`
	Device              []string          `yaml:"device"`
	DeviceCgroupRule    []string          `yaml:"deviceCgroupRule"`
	DeviceReadBPS       []string          `yaml:"deviceReadBps"`
	DeviceReadIOPS      []string          `yaml:"deviceReadIops"`
	DeviceWriteBPS      []string          `yaml:"deviceWriteBps"`
	DeviceWriteIOPS     []string          `yaml:"deviceWriteIops"`
	DisableContentTrust *string           `yaml:"disableContentTrust" validate:"omitempty,bool|env"`
	Domainname          *string           `yaml:"domainname"`
	Entrypoint          *string           `yaml:"entrypoint"`
	Env                 map[string]string `yaml:"env"`
	EnvFile             []string          `yaml:"envFile"`
	Expose              []string          `yaml:"expose"`
	GroupAdd            []string          `yaml:"groupAdd"`
	HealthCmd           *string           `yaml:"healthCmd"`
	HealthInterval      *string           `yaml:"healthInterval" validate:"omitempty,duration|env"`
	HealthRetries       *string           `yaml:"healthRetries" validate:"omitempty,int|env"`
	HealthStartPeriod   *string           `yaml:"healthStartPeriod" validate:"omitempty,duration|env"`
	HealthTimeout       *string           `yaml:"healthTimeout" validate:"omitempty,duration|env"`
	Hostname            *string           `yaml:"hostname"`
	IP                  *string           `yaml:"ip" validate:"omitempty,ipv4|env"`
	IP6                 *string           `yaml:"ip6" validate:"omitempty,ipv6|env"`
	IPC                 *string           `yaml:"ipc"`
	Init                *string           `yaml:"init" validate:"omitempty,bool|env"`
	Interactive         *string           `yaml:"interactive" validate:"omitempty,bool|env" default:"true"`
	Isolation           *string           `yaml:"isolation"`
	KernelMemory        *string           `yaml:"kernelMemory" validate:"omitempty,membytes|env"`
	Label               map[string]string `yaml:"label"`
	LabelFile           []string          `yaml:"labelFile"`
	Link                []string          `yaml:"link"`
	LinkLocalIP         []string          `yaml:"linkLocalIp"`
	LogDriver           *string           `yaml:"logDriver"`
	LogOpt              map[string]string `yaml:"logOpt"`
	MacAddress          *string           `yaml:"macAddress" validate:"omitempty,mac|env"`
	Memory              *string           `yaml:"memory" validate:"omitempty,membytes|env"`
	MemoryReservation   *string           `yaml:"memoryReservation" validate:"omitempty,membytes|env"`
	MemorySwap          *string           `yaml:"memorySwap" validate:"omitempty,memswapbytes|env"`
	MemorySwappiness    *string           `yaml:"memorySwappiness" validate:"omitempty,int64|env,_,env|min=-1,_,env|max=100"`
	Mount               map[string]string `yaml:"mount"`
	Name                *string           `yaml:"name"`
	Network             *string           `yaml:"network" default:"host"`
	NetworkAlias        []string          `yaml:"networkAlias"`
	NoHealthcheck       *string           `yaml:"noHealthcheck" validate:"omitempty,bool|env"`
	OOMKillDisable      *string           `yaml:"oomKillDisable" validate:"omitempty,bool|env"`
	OOMScoreAdj         *string           `yaml:"oomScoreAdj" validate:"omitempty,int|env,_,env|min=-1000,_,env|max=1000"`
	PID                 *string           `yaml:"pid"`
	PidsLimit           *string           `yaml:"pidsLimit" validate:"omitempty,int64|env"`
	Platform            *string           `yaml:"platform"`
	Privileged          *string           `yaml:"privileged" validate:"omitempty,bool|env"`
	Publish             []string          `yaml:"publish"`
	PublishAll          *string           `yaml:"publishAll" validate:"omitempty,bool|env"`
	ReadOnly            *string           `yaml:"readOnly" validate:"omitempty,bool|env"`
	Restart             *string           `yaml:"restart"`
	Rm                  *string           `yaml:"rm" validate:"omitempty,bool|env" default:"true"`
	Runtime             *string           `yaml:"runtime"`
	SecurityOpt         map[string]string `yaml:"securityOpt"`
	ShmSize             *string           `yaml:"shmSize" validate:"omitempty,membytes|env"`
	SigProxy            *string           `yaml:"sigProxy" validate:"omitempty,bool|env"`
	StopSignal          *string           `yaml:"stopSignal"`
	StopTimeout         *string           `yaml:"stopTimeout" validate:"omitempty,int|env"`
	StorageOpt          map[string]string `yaml:"storageOpt"`
	Sysctl              map[string]string `yaml:"sysctl"`
	TTY                 *string           `yaml:"tty" validate:"omitempty,bool|env" default:"$(if tty >/dev/null; then echo true; else echo false; fi)"`
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
