package main

import (
	"fmt"
	dbus "github.com/coreos/go-systemd/dbus"
)

func main() {
	target1 := "cups.service"
	target2 := "unexisting.service"

	conn := SetupConn()
	defer conn.Close()

	units, err := conn.ListUnitsByPatterns([]string{}, []string{"cups*", target2})

	if err != nil {
		panic(err)

	}

	unit := getUnitStatus(units, target1)

	if unit == nil {
		panic("unit not found in list")
	} else if unit.ActiveState != "active" {
		panic("Unit should be active")
	}

	fmt.Println("unit found test OK!")
}

func SetupConn() *dbus.Conn {
	conn, err := dbus.New()
	fmt.Println("created Dbus connection")
	if err != nil {
		panic(err)
	}

	return conn
}

func getUnitStatus(units []dbus.UnitStatus, name string) *dbus.UnitStatus {
	for _, u := range units {
		if u.Name == name {
			return &u
		}
	}
	return nil
}
