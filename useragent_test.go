package ua

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("Random Type", func(test *testing.T) {
		var times = 100000
		for key := range uaCache {
			for i := 0; i < times; i++ {
				RandomType(key)
			}
		}
	})
}

func TestRandom(t *testing.T) {
	t.Run("Random All", func(t *testing.T) {
		var times = 1000000
		for i := 0; i < times; i++ {
			Random()
		}
	})
}
