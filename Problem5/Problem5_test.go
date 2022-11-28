package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCariLuas(t *testing.T) {
	t.Run("Case 1 : [20, 4]", func(t *testing.T) {
		assert.Equal(t, float64(602.88), CariLuas(20, 4), "Hasil harus sama")
	})
	t.Run("Case 2 : [19, 3]", func(t *testing.T) {
		assert.Equal(t, float64(414.48), CariLuas(19, 3), "Hasil harus sama")
	})
	t.Run("Case 3 : [18, 1]", func(t *testing.T) {
		assert.Equal(t, float64(119.32000000000001), CariLuas(18, 1), "Hasil harus sama")
	})
	t.Run("Case 4 : [17, 3]", func(t *testing.T) {
		assert.Equal(t, float64(376.8), CariLuas(17, 3), "Hasil harus sama")
	})
	t.Run("Case 5 : [20, 1]", func(t *testing.T) {
		assert.Equal(t, float64(131.88), CariLuas(20, 1), "Hasil harus sama")
	})
}
