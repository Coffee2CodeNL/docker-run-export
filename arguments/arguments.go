package arguments

type Args struct {
	AddHost             []string
	Attach              []string
	BlkioWeight         int
	BlkioWeightDevice   []string
	CapAdd              []string
	CapDrop             []string
	Cgroupns            string
	CgroupParent        string
	Cidfile             string
	CpuPeriod           int
	CpuQuota            int
	CpuRtPeriod         int
	CpuRtRuntime        int
	Cpus                float32
	CpusetCpus          string
	CpusetMems          string
	CpuShares           int
	Detach              bool
	DetachKeys          string
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
	Domainname          string
	Entrypoint          string
	Env                 []string
	EnvFile             []string
	Expose              []string
	Gpus                string
	GroupAdd            []string
	HealthCmd           string
	HealthInterval      string
	HealthRetries       uint64
	HealthStartPeriod   string
	HealthTimeout       string
	Hostname            string
	Init                bool
	Interactive         bool
	Ip                  string
	Ip6                 string
	Ipc                 string
	Isolation           string
	KernelMemory        int
	Label               []string
	LabelFile           []string
	Link                []string
	LinkLocalIP         []string
	LogDriver           string
	LogOpt              []string
	Mac                 string
	Memory              int64
	MemoryReservation   int64
	MemorySwap          int64
	MemorySwappiness    int64
	Mount               []string
	ContainerName       string
	Network             string
	NetworkAlias        []string
	NoHealthcheck       bool
	OomKillDisable      bool
	OomScore            int
	Pid                 string
	PidsLimit           int
	Platform            string
	Privileged          bool
	Publish             []string
	PublishAll          bool
	Pull                string
	ReadOnly            bool
	Restart             string
	Rm                  bool
	Runtime             string
	SecurityOpt         []string
	ShmSize             int
	SigProxy            bool
	StopSignal          string
	StopTimeout         int
	StorageOpt          []string
	Sysctl              map[string]string
	Tmpfs               []string
	Tty                 bool
	Ulimit              []string
	User                string
	Userns              string
	Uts                 string
	Volume              []string
	VolumeDriver        string
	VolumesFrom         []string
	Workdir             string
}
