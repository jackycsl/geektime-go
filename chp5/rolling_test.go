package rolling

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {

	n := NewNumbers()
	for _, x := range []float64{5, 11, 7} {
		n.UpdateMax(x)
		time.Sleep(1 * time.Second)
	}
	assert.Equal(t, float64(11), n.Max(time.Now()))
}

func TestAvg(t *testing.T) {

	n := NewNumbers()
	for _, x := range []float64{0.1, 0.2, 0.3, 0.4, 0.5} {
		n.UpdateMax(x)
		time.Sleep(1 * time.Second)
	}
	assert.Equal(t, float64(0.15), n.Avg(time.Now()))
}

func TestMaxSlidingWindow(t *testing.T) {

	n := NewNumbers()
	for _, x := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		n.UpdateMax(x)
		time.Sleep(1 * time.Second)
	}
	assert.Equal(t, 8, len(n.MaxSlidingWindow(time.Now(), 3)))
}
