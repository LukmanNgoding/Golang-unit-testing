package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamaOrang(t *testing.T) {
	t.Run("Case 1 = Kobar", func(t *testing.T) {
		assert.Equal(t, "Hello Kobar! Saya Golang bahasa yang sangat Menyenangkan", NamaOrang("Kobar"), "Hasil Harus sama")
	})
	t.Run("Case 2 = lukman", func(t *testing.T) {
		assert.Equal(t, "Hello lukman! Saya Golang bahasa yang sangat Menyenangkan", NamaOrang("lukman"), "Hasil Harus sama")
	})
	t.Run("Case 3 = alif", func(t *testing.T) {
		assert.Equal(t, "Hello alif! Saya Golang bahasa yang sangat Menyenangkan", NamaOrang("alif"), "Hasil Harus sama")
	})
	t.Run("Case 4 = al", func(t *testing.T) {
		assert.Equal(t, "Hello al! Saya Golang bahasa yang sangat Menyenangkan", NamaOrang("al"), "Hasil Harus sama")
	})
	t.Run("Case 5 = moh", func(t *testing.T) {
		assert.Equal(t, "Hello moh! Saya Golang bahasa yang sangat Menyenangkan", NamaOrang("moh"), "Hasil Harus sama")
	})
}
