package Modules

import (
	"runtime"
)


func OsInfo() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	}
	return "MacOS"
}
