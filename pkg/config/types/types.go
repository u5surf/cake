package types

// ConfigSpec holds information needed to register HCI with NKS
type ConfigSpec struct {
	Provider              string `yaml:"Provider" json:"provider"`
	VCenterURL            string `yaml:"VCenterURL" json:"vcenterurl"`
	VCenterUser           string `yaml:"VCenterUser" json:"vcenteruser"`
	VCenterPassword       string `yaml:"VCenterPassword" json:"vcenterpassword"`
	DatacenterID          string `yaml:"DatacenterID" json:"datacenterid"`
	ResourcePoolID        string `yaml:"ResourcePoolID" json:"resourcepoolid"`
	DatastoreID           string `yaml:"DatastoreID" json:"datastoreid"`
	ManagementNetworkID   string `yaml:"ManagementNetworkID" json:"managementnetworkid"`
	ManagementNetworkName string `yaml:"ManagementNetworkName" json:"managementnetworkname"`
	WorkloadNetworkID     string `yaml:"WorkloadNetworkID" json:"workloadnetworkid"`
	WorkloadNetworkName   string `yaml:"WorkloadNetworkName" json:"workloadnetworkname"`
	StorageNetworkID      string `yaml:"StorageNetworkID" json:"storagenetworkid"`
	StorageNetworkName    string `yaml:"StorageNetworkName" json:"storagenetworkname"`
	RegionName            string `yaml:"RegionName" json:"regionname"`
	OrganizationID        string `yaml:"OrganizationID" json:"organizationid"`
	CloudCentralKey       string `yaml:"CloudCentralKey" json:"cloudcentralkey"`
	//Solidfire             Solidfire     `yaml:"Solidfire,omitempty" json:"solidfire,omitempty"`
	//IPAM                  IPAMConfig    `yaml:"IPAM,omitempty" json:"ipam,omitempty"`
	//ProxySettings         ProxySettings `yaml:"ProxySettings,omitempty" json:"proxysettings,omitempty"`
	//OptionalConfiguration Configuration `yaml:"Configuration" json:"configuration,omitempty"`
}
