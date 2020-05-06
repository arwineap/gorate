package gorate

import (
	"fmt"
	"time"
)

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	result := ""
	if h > 0 {
		result = fmt.Sprintf("%s%dh", result, h)
	}
	if m > 0 {
		result = fmt.Sprintf("%s%dm", result, m)
	}
	if s > 0 {
		result = fmt.Sprintf("%s%ds", result, s)
	}
	if result == "" {
		result = "0s"
	}
	return result
}
