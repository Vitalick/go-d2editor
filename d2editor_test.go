package d2editor_test

import (
	"github.com/vitalick/go-d2editor"
	"testing"
)

func benchOne(b *testing.B, ch chan []byte) {
	c, _ := d2editor.NewEmptyCharacter(97)
	bts, _ := c.GetCorrectBytes()
	ch <- bts
}

func BenchmarkCreateEmpty(b *testing.B) {
	var c = make(chan []byte)
	count := 0
	for range make([]bool, b.N) {
		go benchOne(b, c)
	}
	for range c {
		count++
		if count == b.N {
			break
		}
	}
	close(c)
}
