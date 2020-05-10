package gorate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Run("testing client", func(t *testing.T) {
		client := NewClient()
		result := client.NewEntry(1000)
		assert.Equal(t, "1000 ( 0 per 0s ) ( 0 per 0s )\n", result)

		time.Sleep(1 * time.Second)
		result = client.NewEntry(1001)
		assert.Equal(t, "1001 ( +1 per 1s ) ( +1 per 1s )\n", result)

		time.Sleep(1 * time.Second)
		result = client.NewEntry(1002)
		assert.Equal(t, "1002 ( +1 per 1s ) ( +2 per 2s )\n", result)

		time.Sleep(1 * time.Second)
		result = client.NewEntry(1003)
		assert.Equal(t, "1003 ( +1 per 1s ) ( +3 per 3s )\n", result)

		time.Sleep(7 * time.Second)
		result = client.NewEntry(1010)
		assert.Equal(t, "1010 ( +7 per 7s ) ( +10 per 10s )\n", result)

		time.Sleep(1 * time.Second)
		result = client.NewEntry(1002)
		assert.Equal(t, "1002 ( -8 per 1s ) ( +2 per 11s )\n", result)
	})
}
