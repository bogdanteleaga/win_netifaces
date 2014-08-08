package win_netifaces

// These are exactly the types of the classes from WMI
// Reference: http://msdn.microsoft.com/en-us/library/aa394216(v=vs.85).aspx

// The commented properties have not been used for this packages but might be
// needed in the future

type Win32_NetworkAdapter struct {
	AdapterType string
	//AdapterTypeID               uint16
	//AutoSense                   bool
	//Availability                uint16
	//Caption                     string
	//ConfigManagerErrorCode      uint32
	//ConfigManagerUserConfig     bool
	//CreationClassName           string
	Description string
	DeviceID    string
	//ErrorCleared                bool
	//ErrorDescription            string
	GUID  string
	Index uint32
	//Installed                   bool
	InterfaceIndex uint32
	//LastErrorCode               uint32
	MACAddress   string
	Manufacturer string
	//MaxNumberControlled         uint32
	//MaxSpeed                    uint64
	Name string
	//NetConnectionID             string
	NetConnectionStatus uint16
	NetEnabled          bool
	//NetworkAddresses    []string
	PermanentAddress string
	PhysicalAdapter  bool
	PNPDeviceID      string
	//PowerManagementCapabilities []uint16
	//PowerManagementSupported    bool
	//ProductName                 string
	//ServiceName                 string
	//Speed                       uint64
	//Status     string
	//StatusInfo uint16
	//SystemCreationClassName     string
	//SystemName                  string
}

type Win32_NetworkAdapterConfiguration struct {
	//ArpAlwaysSourceRoute         bool
	//ArpUseEtherSNAP              bool
	//Caption                      string
	//DatabasePath                 string
	//DeadGWDetectEnabled          bool
	DefaultIPGateway []string
	//DefaultTOS                   uint8
	//DefaultTTL                   uint8
	Description string
	DHCPEnabled bool
	DHCPServer  string
	//DNSDomain                    string
	//DNSDomainSuffixSearchOrder   []string
	//DNSEnabledForWINSResolution  bool
	DNSHostName string
	//DNSServerSearchOrder         []string
	//DomainDNSRegistrationEnabled bool
	//ForwardBufferMemory          uint32
	//FullDNSRegistrationEnabled   bool
	//GatewayCostMetric            []uint16
	//IGMPLevel                    uint8
	Index          uint32
	InterfaceIndex uint32
	IPAddress      []string
	//IPConnectionMetric           uint32
	IPEnabled bool
	//IPFilterSecurityEnabled      bool
	//IPPortSecurityEnabled        bool
	//IPSecPermitIPProtocols       []string
	//IPSecPermitTCPPorts          []string
	//IPSecPermitUDPPorts          []string
	//IPSubnet                     []string
	//IPUseZeroBroadcast           bool
	//IPXAddress                   string
	//IPXEnabled                   bool
	//IPXFrameType                 []uint32
	//IPXMediaType                 uint32
	//IPXNetworkNumber             []string
	//IPXVirtualNetNumber          string
	//KeepAliveInterval            uint32
	//KeepAliveTime                uint32
	//MACAddress string
	MTU uint32
	//NumForwardPackets            uint32
	//PMTUBHDetectEnabled          bool
	//PMTUDiscoveryEnabled         bool
	//ServiceName                  string
	//SettingID                    string
	//TcpipNetbiosOptions          uint32
	//TcpMaxConnectRetransmissions uint32
	//TcpMaxDataRetransmissions    uint32
	//TcpNumConnections            uint32
	//TcpUseRFC1122UrgentPointer   bool
	TcpWindowSize uint16
	//WINSEnableLMHostsLookup      bool
	//WINSHostLookupFile           string
	//WINSPrimaryServer            string
	//WINSScopeID                  string
	//WINSSecondaryServer          string
}
