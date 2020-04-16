package cloudinit

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateMetadata(t *testing.T) {

	t.Run("WhenDHCP", func(tt *testing.T) {

		expectedMetadata := `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: dhcp
`

		values := &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       true,
					IPAddress:   "",
					Netmask:     "",
					Gateway:     "",
					NameServers: nil,
					DNSSearch:   nil,
				},
			},
		}

		metadataBytes, err := GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

	})

	t.Run("WhenStatic", func(tt *testing.T) {

		expectedMetadata := `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: static
          address: ipaddress
          netmask: netmask
          gateway: gateway
          dns_nameservers:
            - dns1
            - dns2
          dns_search:
            - search1
            - search2
`

		values := &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       false,
					IPAddress:   "ipaddress",
					Netmask:     "netmask",
					Gateway:     "gateway",
					NameServers: []string{"dns1", "dns2"},
					DNSSearch:   []string{"search1", "search2"},
				},
			},
		}

		metadataBytes, err := GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

	})

	t.Run("WhenStaticWithMissingValues", func(tt *testing.T) {

		expectedMetadata := `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: static
          address: ipaddress
          gateway: gateway
          dns_nameservers:
            - dns1
            - dns2
          dns_search:
            - search1
            - search2
`

		values := &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       false,
					IPAddress:   "ipaddress",
					Netmask:     "",
					Gateway:     "gateway",
					NameServers: []string{"dns1", "dns2"},
					DNSSearch:   []string{"search1", "search2"},
				},
			},
		}

		metadataBytes, err := GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

		expectedMetadata = `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: static
          address: ipaddress
          netmask: netmask
          dns_nameservers:
            - dns1
            - dns2
          dns_search:
            - search1
            - search2
`

		values = &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       false,
					IPAddress:   "ipaddress",
					Netmask:     "netmask",
					Gateway:     "",
					NameServers: []string{"dns1", "dns2"},
					DNSSearch:   []string{"search1", "search2"},
				},
			},
		}

		metadataBytes, err = GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

		expectedMetadata = `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: static
          address: ipaddress
          netmask: netmask
          gateway: gateway
          dns_search:
            - search1
            - search2
`

		values = &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       false,
					IPAddress:   "ipaddress",
					Netmask:     "netmask",
					Gateway:     "gateway",
					NameServers: []string{},
					DNSSearch:   []string{"search1", "search2"},
				},
			},
		}

		metadataBytes, err = GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

		expectedMetadata = `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: static
          address: ipaddress
          netmask: netmask
          gateway: gateway
          dns_nameservers:
            - dns1
            - dns2
`

		values = &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       false,
					IPAddress:   "ipaddress",
					Netmask:     "netmask",
					Gateway:     "gateway",
					NameServers: []string{"dns1", "dns2"},
					DNSSearch:   []string{},
				},
			},
		}

		metadataBytes, err = GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

	})

	t.Run("WhenMultipleNICsDHCP", func(tt *testing.T) {

		expectedMetadata := `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: dhcp
    - type: physical
      name: id1
      mac_address: mac2
      subnets:
        - type: dhcp
`

		values := &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       true,
					IPAddress:   "",
					Netmask:     "",
					Gateway:     "",
					NameServers: nil,
					DNSSearch:   nil,
				},
				{
					MACAddress:  "mac2",
					DHCP4:       true,
					IPAddress:   "",
					Netmask:     "",
					Gateway:     "",
					NameServers: nil,
					DNSSearch:   nil,
				},
			},
		}

		metadataBytes, err := GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

	})

	t.Run("WhenMultipleNICsStatic", func(tt *testing.T) {

		expectedMetadata := `
instance-id: "host"
local-hostname: "host"
network:
  version: 1
  config:
    - type: physical
      name: id0
      mac_address: mac
      subnets:
        - type: static
          address: ipaddress
          netmask: netmask
          gateway: gateway
          dns_nameservers:
            - dns1
            - dns2
          dns_search:
            - search1
            - search2
    - type: physical
      name: id1
      mac_address: mac2
      subnets:
        - type: static
          address: ipaddress2
          netmask: netmask2
          gateway: gateway2
          dns_nameservers:
            - dns12
            - dns22
          dns_search:
            - search12
            - search22
`

		values := &MetadataValues{
			Hostname: "host",
			Networks: []NetworkConfig{
				{
					MACAddress:  "mac",
					DHCP4:       false,
					IPAddress:   "ipaddress",
					Netmask:     "netmask",
					Gateway:     "gateway",
					NameServers: []string{"dns1", "dns2"},
					DNSSearch:   []string{"search1", "search2"},
				},
				{
					MACAddress:  "mac2",
					DHCP4:       false,
					IPAddress:   "ipaddress2",
					Netmask:     "netmask2",
					Gateway:     "gateway2",
					NameServers: []string{"dns12", "dns22"},
					DNSSearch:   []string{"search12", "search22"},
				},
			},
		}

		metadataBytes, err := GetMetadata(values)
		require.NoError(tt, err)
		assert.Equal(tt, expectedMetadata, string(metadataBytes))

	})

}
