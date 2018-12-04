package main_test

import (
	"fmt"
	dbus "github.com/coreos/go-systemd/dbus"
	. "github.com/onsi/ginkgo"
)

func setupDbusConn() *dbus.Conn {
	conn, err := dbus.New()
	if err != nil {
		panic(err)
	}
	fmt.Println("Dbus connection created!")
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

var _ = Describe("Check kubic-init.service", func() {
	It("kubic-init.service should be up and running", func() {
		target1 := "cups.service"
		service_regex := "cups*"

		dbusConn := setupDbusConn()
		defer dbusConn.Close()
		units, err := dbusConn.ListUnitsByPatterns([]string{}, []string{service_regex})

		if err != nil {
			Fail("No systemd-units found!")

		}

		unit := getUnitStatus(units, target1)

		if unit == nil {
			Fail("unit not found in list")
		} else if unit.ActiveState != "active" {
			Fail("Unit should be active")
		}

		fmt.Println("unit found test OK!")

	})
})
