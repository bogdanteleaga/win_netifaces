package win_netifaces

import (
	"net"
	"regexp"
	"strings"
)

type InterfaceType int

const (
	Physical = 0
	Virtual  = 1
)

type NetworkAdapter struct {
	InterfaceIndex uint32
	Name           string
	FriendlyName   string
	Protocol       string
	DHCPServer     string
	IP             []string
	MAC            string
	MTU            int
	Type           int
	Up             bool
}

// Returns a slice of NetworkAdapters of the given InterfaceType
func GetAdapters(ifType InterfaceType) (ifs []NetworkAdapter, err error) {
	adapters, err := getAdaptersInfoFromWmi()
	if err != nil {
		return nil, err
	}

	configs, err := getAdaptersConfigInfoFromWmi()
	if err != nil {
		return nil, err
	}

	result := make([]NetworkAdapter, 0, len(adapters))

	for _, adapter := range adapters {
		if !isValidMACAddress(adapter.MACAddress) {
			continue
		}
		for _, config := range configs {
			if adapter.InterfaceIndex == config.InterfaceIndex {
				//Create new adapter
				res := NetworkAdapter{}

				//Check the interface type and see if it matches
				if ifType == Physical && adapter.Manufacturer != "Windows" && !strings.HasPrefix(adapter.PNPDeviceID, "ROOT") {
					res.Type = Physical
				} else if ifType == Virtual && strings.HasPrefix(adapter.PNPDeviceID, "ROOT") {
					res.Type = Virtual
				} else {
					break
				}
				//Other adapter type logic can be filled here

				//Fill easy fields
				res.InterfaceIndex = adapter.InterfaceIndex
				res.Name = adapter.GUID
				res.FriendlyName = adapter.Name
				res.MAC = adapter.MACAddress
				res.Up = adapter.NetEnabled

				//Check for DHCP
				if config.DHCPEnabled {
					res.Protocol = "DHCP"
					res.DHCPServer = config.DHCPServer
				} else {
					res.Protocol = "Static"
					res.DHCPServer = "No DHCP"
				}

				//Get ip addresses and mtu from net package since WMI doesn't
				//return MTU and go-ole doesn't process string arrays(for ips)
				addrs, mtu, err := getIPAddressFromIndex(adapter.InterfaceIndex)
				if err != nil {
					return nil, err
				}
				res.IP = addrs
				res.MTU = mtu

				result = append(result, res)

			}
		}
	}
	return result, nil

}

// Looks into the net package for a given interface index and returns it's IP
// Addresses. Needed for now since we can't get IP Addresses from go-ole
func getIPAddressFromIndex(idx uint32) (addrs []string, mtu int, err error) {
	iface, err := net.InterfaceByIndex(int(idx))
	if err != nil {
		return make([]string, 0), 0, nil
	} else {
		addrs, err := iface.Addrs()
		if err != nil {
			return make([]string, 0), 0, err
		}
		res := make([]string, len(addrs))
		for idx, addr := range addrs {
			res[idx] = addr.String()
		}
		return res, iface.MTU, nil
	}
}

func isValidMACAddress(addr string) bool {
	var validMAC = regexp.MustCompile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")

	return validMAC.MatchString(addr)
}
