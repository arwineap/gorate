package gorate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFullFormatter(t *testing.T) {
	now := time.Now()

	t.Run("testing full formatter", func(t *testing.T) {
		formatter := FullFormatter{}
		result := formatter.String(
			FormatterStats{
				start:        now.Add(-time.Second),
				entryCount:   1,
				initialValue: 0,
			},
			Entry{
				timestamp: now.Add(-time.Second),
				value:     0,
			},
			Entry{
				timestamp: now,
				value:     1,
			},
		)

		assert.Equal(t, "1 ( +1 per 1s ) ( +1 per 1s )\n", result)
	})

	t.Run("testing sequence", func(t *testing.T) {
		formatter := FullFormatter{}
		result := formatter.String(
			FormatterStats{
				start:        now.Add(-time.Second * 11),
				entryCount:   6,
				initialValue: 1000,
			},
			Entry{
				timestamp: now.Add(-time.Second),
				value:     1010,
			},
			Entry{
				timestamp: now,
				value:     1002,
			},
		)

		assert.Equal(t, "1002 ( -8 per 1s ) ( +2 per 11s )\n", result)
	})

}
