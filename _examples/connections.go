package main

import (
	"context"
	"log"

	"github.com/godbus/dbus/v5"
	"github.com/muka/network_manager"
)

var nmNs = network_manager.InterfaceNetworkManager

func fail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// nm := network_manager.NewNetworkManager(conn.Object(nmNs, dbus.ObjectPath("/org/freedesktop/NetworkManager")))

	settings := network_manager.NewNetworkManager_Settings(conn.Object(nmNs, dbus.ObjectPath("/org/freedesktop/NetworkManager/Settings")))

	hostname, err := settings.GetHostname(context.Background())
	fail(err)
	log.Printf("hostname=%s\n", hostname)

	conns, err := settings.GetConnections(context.Background())
	fail(err)

	for _, connPath := range conns {
		netconn := network_manager.NewNetworkManager_Settings_Connection(conn.Object(nmNs, connPath))
		connSettings, err := netconn.GetSettings(context.Background())
		fail(err)

		log.Printf("%s\n", connPath)
		log.Printf("%++v\n\n", connSettings)
	}

}
