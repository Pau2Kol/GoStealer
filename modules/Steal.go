package Modules

import (
	"net"
	"os"
	"runtime"
)

// OsInfo return the name of the operating system where the program is running
// it's a simple switch case that return the name of the os
// it's not the best way to do it but it's the simplest one
func OsInfo() string {
	//j'ai pas tout compris mais en gros GOOS c un truc qui peut trouver l'os grace a une variable au démarrage
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	}
	return "MacOS"
}

// HostnameInfo returns the hostname of the system where the program is running.
// It utilizes the os.Hostname() function to retrieve this information.
func HostnameInfo() string {
	hostname, _ := os.Hostname()
	return hostname
}

// HomeDirInfo returns the home directory of the user where the program is running.
// It utilizes the os.UserHomeDir() function to retrieve this information.
func HomeDirInfo() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir
}

func MacAdressInfo() string {
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {

		if i.Flags&net.FlagUp != 0 { // net.flagsup Vérifie l'état de l'interface 0 veut dire qu'elle est down

			return i.HardwareAddr.String() //HardwareAddr return une slice ducoup faut la convertir en string
		}
	}
	return "ce message n'est pas normal il y'a eu un prblm au niveau de l'adresse mac"
}

func IpAdressInfo() string {
	address, _ := net.InterfaceAddrs()

	for _, address := range address {
		ipNet, _ := address.(*net.IPNet)                     //on convertit toute les addresse vers IPnet pour avoir ip/masque
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil { //on verifie que l'adresse n'est pas une loopback et qu'elle est une adresse ipv4

			return ipNet.IP.String()
		}

	}
	return "prblm récup ip"
}

func NetworkInfo() string {
	return "j'ai eu la flemme TODO"
}
