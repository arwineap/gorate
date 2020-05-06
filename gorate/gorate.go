package gorate

import (
	"sync"
	"time"
)

type GoRateClient interface {
	NewEntry(v float64)
}

type rateClient struct {
	previousValue     int64
	previousTimestamp time.Time
	startTime         time.Time
	initialEntry      bool
	initialValue      int64
	handler           func(currentValue int64, difference int64, duration string, cumDifference int64, cumDuration string)
	sync.Mutex
}

func NewClient(fn func(currentValue int64, difference int64, duration string, cumDifference int64, cumDuration string)) rateClient {
	now := time.Now()
	return rateClient{
		startTime:         now,
		initialEntry:      true,
		initialValue:      0,
		previousTimestamp: now,
		previousValue:     0,
		handler:           fn,
	}
}

func (r *rateClient) NewEntry(currentValue int64) {
	if r.initialEntry {
		r.setInitialEntry(false)
		r.setInitialValue(currentValue)
		r.setPreviousValue(currentValue)
	}
	now := time.Now()
	duration := fmtDuration(now.Sub(r.previousTimestamp))
	difference := currentValue - r.previousValue
	r.setPreviousTime(now)
	r.setPreviousValue(currentValue)
	cumDuration := fmtDuration(now.Sub(r.startTime))
	cumDifference := currentValue - r.initialValue
	r.handler(currentValue, difference, duration, cumDifference, cumDuration)
}

func (r *rateClient) setInitialEntry(b bool) {
	r.Lock()
	defer r.Unlock()
	r.initialEntry = b
}

func (r *rateClient) setInitialValue(i int64) {
	r.Lock()
	defer r.Unlock()
	r.initialValue = i
}

func (r *rateClient) setPreviousValue(i int64) {
	r.Lock()
	defer r.Unlock()
	r.previousValue = i
}

func (r *rateClient) setPreviousTime(t time.Time) {
	r.Lock()
	defer r.Unlock()
	r.previousTimestamp = t
}
