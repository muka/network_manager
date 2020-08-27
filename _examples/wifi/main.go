package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/muka/network_manager"
)

var nmNs = network_manager.InterfaceNetworkManager

func main() {

	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	nm := network_manager.NewNetworkManager(conn.Object(nmNs, dbus.ObjectPath("/org/freedesktop/NetworkManager")))

	// listActiveConnections(nm, conn)
	devices, err := listWifiDevices(nm, conn)
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {

		connectivity, err := device.GetIp4Connectivity(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		iface, err := device.GetInterface(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		activeConn, err := device.GetActiveConnection(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		if len(activeConn) == 0 {
			log.Printf("WIFI %s is not connected %+v\n", iface)
		} else {

			label := ""
			if network_manager.NM_CONNECTIVITY_NONE == connectivity {
				label = "is not connected"
			}
			if network_manager.NM_CONNECTIVITY_FULL == connectivity {
				label = "is connected"
			}
			if network_manager.NM_CONNECTIVITY_LIMITED == connectivity {
				label = "has limited connectivity"
			}

			log.Printf("WIFI %s %s\n", iface, label)
		}
	}

	os.Exit(0)
}

func listWifiDevices(nm *network_manager.NetworkManager, conn *dbus.Conn) ([]*network_manager.NetworkManager_Device, error) {

	wifi := []*network_manager.NetworkManager_Device{}

	devices, err := nm.GetAllDevices(context.Background())
	if err != nil {
		return wifi, err
	}

	for _, devicePath := range devices {
		device := network_manager.NewNetworkManager_Device(conn.Object(nmNs, devicePath))

		deviceType, err := device.GetDeviceType(context.Background())
		if err != nil {
			log.Printf("Error %s: %s", devicePath, err)
			continue
		}

		if network_manager.NM_DEVICE_TYPE_WIFI == deviceType {
			wifi = append(wifi, device)
		}
	}

	return wifi, nil
}

func listActiveConnections(nm *network_manager.NetworkManager, conn *dbus.Conn) {

	connections, err := nm.GetActiveConnections(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Active connections: %+v\n", connections)

	for _, connectionPath := range connections {

		activeConn := network_manager.NewNetworkManager_Connection_Active(conn.Object(nmNs, connectionPath))

		conntype, err := activeConn.GetType(context.Background())
		if err != nil {
			log.Printf("Error %s: %s", connectionPath, err)
			continue
		}

		ip4config, err := activeConn.GetIp4Config(context.Background())
		if err != nil {
			log.Printf("%s: %s", connectionPath, err)
			continue
		}

		fmt.Printf("type=%s ip4config=%s\n", conntype, ip4config)
	}

}
