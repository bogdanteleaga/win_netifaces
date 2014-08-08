package win_netifaces

import (
	"fmt"
	"reflect"

	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

// Function for pretty printing an interface
// More types can be added if needed
func PrettyPrintInterface(adapter NetworkAdapter) {
	fmt.Printf("Interface Index %d\n", adapter.InterfaceIndex)
	fmt.Printf("Name: %s\n", adapter.Name)
	fmt.Printf("FriendlyName: %s\n", adapter.FriendlyName)
	fmt.Printf("Protocol: %s\n", adapter.Protocol)
	fmt.Printf("DHCPServer: %s\n", adapter.DHCPServer)
	fmt.Printf("IP: %s\n", adapter.IP)
	fmt.Printf("MAC: %s\n", adapter.MAC)
	fmt.Printf("MTU: %d\n", adapter.MTU)

	var stype string
	switch adapter.Type {
	case Physical:
		stype = "Physical"
	case Virtual:
		stype = "Virtual"
	}
	fmt.Printf("Type: %s\n", stype)

	fmt.Printf("Up: %t\n", adapter.Up)
}

// The following 2 functions connect to WMI make a query and then process it
// accordingly
func getAdaptersInfoFromWmi() (adapters []Win32_NetworkAdapter, err error) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	obj, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		return nil, err
	}
	defer obj.Release()

	wmi, err := obj.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	defer wmi.Release()

	// service is a SWbemServices
	serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer")
	if err != nil {
		return nil, err
	}
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	// result is a SWBemObjectSet
	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_NetworkAdapter")
	if err != nil {
		return nil, err
	}
	result := resultRaw.ToIDispatch()
	defer result.Release()

	countVar, err := oleutil.GetProperty(result, "Count")
	if err != nil {
		return nil, err
	}
	count := int(countVar.Val)

	adapters = make([]Win32_NetworkAdapter, count)
	parseAdapters(adapters, result, count)

	return adapters, nil

}

func getAdaptersConfigInfoFromWmi() (configs []Win32_NetworkAdapterConfiguration, err error) {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	obj, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		return nil, err
	}
	defer obj.Release()

	wmi, err := obj.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	defer wmi.Release()

	// service is a SWbemServices
	serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer")
	if err != nil {
		return nil, err
	}
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	// result is a SWBemObjectSet
	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", "SELECT * FROM Win32_NetworkAdapterConfiguration")
	if err != nil {
		return nil, err
	}
	result := resultRaw.ToIDispatch()
	defer result.Release()

	countVar, err := oleutil.GetProperty(result, "Count")
	if err != nil {
		return nil, err
	}
	count := int(countVar.Val)

	adaptersConfig := make([]Win32_NetworkAdapterConfiguration, count)
	parseAdaptersConfiguration(adaptersConfig, result, count)

	return adaptersConfig, nil

}

// The following 2 functions parse the output from the respective queries
func parseAdapters(adapters []Win32_NetworkAdapter, result *ole.IDispatch, count int) (err error) {
	for i := 0; i < count; i++ {
		// item is a SWbemObject, but really a Win32_Process
		itemRaw, _ := oleutil.CallMethod(result, "ItemIndex", i)
		item := itemRaw.ToIDispatch()
		defer item.Release()

		adapters[i].AdapterType = oleutil.MustGetProperty(item, "AdapterType").ToString()

		adapters[i].Description = oleutil.MustGetProperty(item, "Description").ToString()

		adapters[i].DeviceID = oleutil.MustGetProperty(item, "DeviceID").ToString()

		adapters[i].GUID = oleutil.MustGetProperty(item, "GUID").ToString()

		idx := oleutil.MustGetProperty(item, "Index").Value().(int64)
		adapters[i].Index = uint32(reflect.ValueOf(idx).Int())

		if_idx := oleutil.MustGetProperty(item, "InterfaceIndex").Value().(int64)
		adapters[i].InterfaceIndex = uint32(reflect.ValueOf(if_idx).Int())

		adapters[i].MACAddress = oleutil.MustGetProperty(item, "MACAddress").ToString()

		adapters[i].Manufacturer = oleutil.MustGetProperty(item, "Manufacturer").ToString()

		adapters[i].Name = oleutil.MustGetProperty(item, "Name").ToString()

		net_conn_status := oleutil.MustGetProperty(item, "NetConnectionStatus").Value()
		if reflect.TypeOf(net_conn_status) == reflect.TypeOf(int64(0)) {
			adapters[i].NetConnectionStatus = uint16(reflect.ValueOf(net_conn_status).Int())
		} else {
			adapters[i].NetConnectionStatus = 0
		}

		net_enabled := oleutil.MustGetProperty(item, "NetEnabled").Value()
		if reflect.TypeOf(net_enabled) == reflect.TypeOf(true) {
			adapters[i].NetEnabled = net_enabled.(bool)
		} else {
			adapters[i].NetEnabled = false
		}

		physical_adapter := oleutil.MustGetProperty(item, "PhysicalAdapter").Value()
		if reflect.TypeOf(physical_adapter) == reflect.TypeOf(true) {
			adapters[i].PhysicalAdapter = physical_adapter.(bool)
		} else {
			adapters[i].PhysicalAdapter = false
		}

		adapters[i].PNPDeviceID = oleutil.MustGetProperty(item, "PNPDeviceID").ToString()
	}
	return nil
}

func parseAdaptersConfiguration(adapters []Win32_NetworkAdapterConfiguration, result *ole.IDispatch, count int) (err error) {
	for i := 0; i < count; i++ {
		// item is a SWbemObject, but really a Win32_Process
		itemRaw, _ := oleutil.CallMethod(result, "ItemIndex", i)
		item := itemRaw.ToIDispatch()
		defer item.Release()

		adapters[i].Description = oleutil.MustGetProperty(item, "Description").ToString()

		adapters[i].DNSHostName = oleutil.MustGetProperty(item, "DNSHostName").ToString()

		idx := oleutil.MustGetProperty(item, "Index").Value().(int64)
		adapters[i].Index = uint32(reflect.ValueOf(idx).Int())

		if_idx := oleutil.MustGetProperty(item, "InterfaceIndex").Value().(int64)
		adapters[i].InterfaceIndex = uint32(reflect.ValueOf(if_idx).Int())

		// MTU information does not currently work properly for whatever reason
		/*
		 *mtu := oleutil.MustGetProperty(item, "MTU").Value()
		 *if reflect.TypeOf(mtu) == reflect.TypeOf(int64(0)) {
		 *    fmt.Println("HERE")
		 *    adapters[i].MTU = uint32(reflect.ValueOf(mtu).Int())
		 *} else {
		 *    adapters[i].MTU = 0
		 *}
		 */

		ip_enabled := oleutil.MustGetProperty(item, "IPEnabled").Value()
		if reflect.TypeOf(ip_enabled) == reflect.TypeOf(true) {
			adapters[i].IPEnabled = ip_enabled.(bool)
		} else {
			adapters[i].IPEnabled = false
		}

		dhcp_enabled := oleutil.MustGetProperty(item, "DHCPEnabled").Value()
		if reflect.TypeOf(dhcp_enabled) == reflect.TypeOf(true) {
			adapters[i].DHCPEnabled = dhcp_enabled.(bool)
		} else {
			adapters[i].DHCPEnabled = false
		}

		adapters[i].DHCPServer = oleutil.MustGetProperty(item, "DHCPServer").ToString()

		// Unpacking string arrays doesn't work yet in go-ole
		/*
			default_ip_gateway := oleutil.MustGetProperty(item, "IPAddress").ToArray()
			_, err = default_ip_gateway.GetType()
			if err != nil {
				adapters[i].DefaultIPGateway = nil
			} else {
			    adapters[i].DefaultIPGateway = default_ip_gateway.ToStringArray()
			}
		*/

	}
	return nil
}
