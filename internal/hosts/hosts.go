package hosts

import "runtime"

func GetHostsPath() string {
	switch runtime.GOOS {
	case "windows":
		return `C:\Windows\System32\drivers\etc\hosts`
	default:
		return "/etc/hosts"
	}
}