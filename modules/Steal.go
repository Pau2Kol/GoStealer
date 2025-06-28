package Modules

import (
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
