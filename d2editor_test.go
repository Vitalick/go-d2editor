package d2editor_test

import (
	"encoding/json"
	"github.com/vitalick/go-d2editor"
	"testing"
)

func createEmpty(b *testing.B) *d2editor.Character {
	c, _ := d2editor.NewEmptyCharacter(97)
	return c
}

func openOne(b *testing.B) *d2editor.Character {
	c, _ := d2editor.Open("LoD-Druid.d2s")
	return c
}

func openJSON(b *testing.B) *d2editor.Character {
	c, _ := d2editor.OpenJSON("LoD-Druid2.json")
	return c
}

func saveOne(b *testing.B, c *d2editor.Character) {
	c.GetBytes()
}

func saveJSON(b *testing.B, c *d2editor.Character) {
	json.Marshal(c)
}

func createSaveEmpty(b *testing.B) {
	saveOne(b, createEmpty(b))
}

func openSaveOne(b *testing.B) {
	saveOne(b, openOne(b))
}

func openSaveJSON(b *testing.B) {
	saveJSON(b, openJSON(b))
}

func BenchmarkCreateEmptyConcurrently(b *testing.B) {
	for range make([]bool, b.N) {
		go createEmpty(b)
	}
}

func BenchmarkSaveEmptyConcurrently(b *testing.B) {
	c := createEmpty(b)
	b.ResetTimer()
	for range make([]bool, b.N) {
		go saveOne(b, c)
	}
}

func BenchmarkCreateSaveEmptyConcurrently(b *testing.B) {
	for range make([]bool, b.N) {
		go createSaveEmpty(b)
	}
}

func BenchmarkOpenConcurrently(b *testing.B) {
	for range make([]bool, b.N) {
		go openOne(b)
	}
}

func BenchmarkSaveFileConcurrently(b *testing.B) {
	c := openOne(b)
	b.ResetTimer()
	for range make([]bool, b.N) {
		go saveOne(b, c)
	}
}

func BenchmarkOpenSaveConcurrently(b *testing.B) {
	for range make([]bool, b.N) {
		go openSaveOne(b)
	}
}

func BenchmarkOpenJSONConcurrently(b *testing.B) {
	for range make([]bool, b.N) {
		go openJSON(b)
	}
}

func BenchmarkSaveJSONConcurrently(b *testing.B) {
	c := openJSON(b)
	b.ResetTimer()
	for range make([]bool, b.N) {
		go saveJSON(b, c)
	}
}

func BenchmarkOpenSaveJSONConcurrently(b *testing.B) {
	for range make([]bool, b.N) {
		go openSaveJSON(b)
	}
}
