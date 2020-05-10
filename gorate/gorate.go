package gorate

import (
	"sync"
	"time"
)

type Entry struct {
	timestamp time.Time
	value     int64
}

type FormatterStats struct {
	start        time.Time
	entryCount   int64
	initialValue int64
}

type GoRateClient interface {
	NewEntry(v float64)
}

type rateClient struct {
	previousEntry Entry
	startTime     time.Time
	initialValue  int64
	entryCount    int64
	Formatter     Formatter
	sync.Mutex
}

func NewClient() rateClient {
	now := time.Now()
	return rateClient{
		startTime:    now,
		initialValue: 0,
		entryCount:   0,
		previousEntry: Entry{
			timestamp: now,
			value:     0,
		},
	}
}

func (r *rateClient) NewEntry(currentValue int64) string {
	now := time.Now()
	if r.entryCount == 0 {
		r.setInitialValue(currentValue)
		r.setPreviousEntry(Entry{timestamp: now, value: currentValue})
	}
	previousEntry := r.previousEntry
	currentEntry := Entry{timestamp: now, value: currentValue}
	stats := r.getFormatterStats()
	r.incrementEntry()
	r.setPreviousEntry(currentEntry)
	return r.getFormatter().String(stats, previousEntry, currentEntry)
}

func (r *rateClient) setPreviousEntry(e Entry) {
	r.Lock()
	defer r.Unlock()
	r.previousEntry = e
}

func (r *rateClient) incrementEntry() {
	r.Lock()
	defer r.Unlock()
	r.entryCount++
}

func (r *rateClient) setInitialValue(i int64) {
	r.Lock()
	defer r.Unlock()
	r.initialValue = i
}

func (r *rateClient) SetFormatter(f Formatter) {
	r.Formatter = f
}

func (r *rateClient) getFormatter() Formatter {
	if r.Formatter == nil {
		return FullFormatter{}
	}
	return r.Formatter
}

func (r *rateClient) getFormatterStats() FormatterStats {
	return FormatterStats{start: r.startTime, entryCount: r.entryCount, initialValue: r.initialValue}
}
