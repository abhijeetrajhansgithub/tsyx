package uptime

import "fmt"

func Format(uptime DeviceUptime) string {
	return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", uptime.Days, uptime.Hours, uptime.Minutes, uptime.Seconds)
}