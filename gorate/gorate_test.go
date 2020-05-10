package gorate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	t.Parallel()
	t.Run("testing client", func(t *testing.T) {

		client := NewClient()
		result := client.NewEntry(1000)
		assert.Equal(t, "1000 ( 0 per 0s ) ( 0 per 0s )\n", result)

		time.Sleep(2 * time.Second)
		result = client.NewEntry(1002)
		assert.Equal(t, "1002 ( +2 per 2s ) ( +2 per 2s )\n", result)

		time.Sleep(1 * time.Second)
		result = client.NewEntry(1000)
		assert.Equal(t, "1000 ( -2 per 1s ) ( 0 per 3s )\n", result)
	})
}

func TestCumClient(t *testing.T) {
	t.Parallel()
	t.Run("testing client", func(t *testing.T) {
		client := NewClient()
		result := client.NewEntry(1000)
		assert.Equal(t, "1000 ( 0 per 0s ) ( 0 per 0s )\n", result)

		time.Sleep(2 * time.Second)
		result = client.NewEntry(1002)
		assert.Equal(t, "1002 ( +2 per 2s ) ( +2 per 2s )\n", result)

		time.Sleep(1 * time.Second)
		result = client.NewEntry(1000)
		assert.Equal(t, "1000 ( -2 per 1s ) ( 0 per 3s )\n", result)
	})
}
