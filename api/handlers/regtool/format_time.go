package v1

import "time"

func Format(time time.Time) string {
	if "0001-01-01 00:00:00" == time.Format("2006-01-02 15:04:05") {
		return ""
	}
	return time.Format("2006-01-02 15:04:05")
}
