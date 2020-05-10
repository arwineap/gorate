package gorate

import (
	"fmt"
	"time"
)

type Formatter interface {
	String(stats FormatterStats, previous Entry, current Entry) string
}

type FullFormatter struct{}

func (FullFormatter) String(stats FormatterStats, previous Entry, current Entry) string {
	duration := fmtDuration(current.timestamp.Sub(previous.timestamp))
	difference := current.value - previous.value
	plusSign := ""
	if difference > 0 {
		plusSign = "+"
	}

	cumDuration := fmtDuration(current.timestamp.Sub(stats.start))
	cumDifference := current.value - stats.initialValue
	cumPlusSign := ""
	if cumDifference > 0 {
		cumPlusSign = "+"
	}

	return fmt.Sprintf("%d ( %s%d per %s ) ( %s%d per %s )\n", current.value, plusSign, difference, duration, cumPlusSign, cumDifference, cumDuration)
}

type InstantaneousFormatter struct{}

func (InstantaneousFormatter) String(stats FormatterStats, previous Entry, current Entry) string {
	duration := fmtDuration(current.timestamp.Sub(previous.timestamp))
	difference := current.value - previous.value

	plusSign := ""
	if difference > 0 {
		plusSign = "+"
	}

	return fmt.Sprintf("%d ( %s%d per %s )\n", current.value, plusSign, difference, duration)
}

type CumulativeFormatter struct{}

func (CumulativeFormatter) String(stats FormatterStats, previous Entry, current Entry) string {
	cumDuration := fmtDuration(current.timestamp.Sub(stats.start))
	cumDifference := current.value - stats.initialValue
	cumPlusSign := ""
	if cumDifference > 0 {
		cumPlusSign = "+"
	}

	return fmt.Sprintf("%d ( %s%d per %s )\n", current.value, cumPlusSign, cumDifference, cumDuration)
}

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
