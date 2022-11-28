package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegiTiga(t *testing.T) {
	t.Run("Case 1 : [20, 25]", func(t *testing.T) {
		assert.Equal(t, float64(250), SegiTiga(20, 25), "Hasil harus sama")
	})
	t.Run("Case 2 : [21, 24]", func(t *testing.T) {
		assert.Equal(t, float64(252), SegiTiga(21, 24), "Hasil harus sama")
	})
	t.Run("Case 3 : [22, 20]", func(t *testing.T) {
		assert.Equal(t, float64(220), SegiTiga(22, 20), "Hasil harus sama")
	})
	t.Run("Case 4 : [23, 22]", func(t *testing.T) {
		assert.Equal(t, float64(253), SegiTiga(23, 22), "Hasil harus sama")
	})
	t.Run("Case 5 : [24, 21]", func(t *testing.T) {
		assert.Equal(t, float64(252), SegiTiga(24, 21), "Hasil harus sama")
	})
}
