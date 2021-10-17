package d2editor_test

import (
	"encoding/json"
	"github.com/vitalick/go-d2editor"
	"sync"
	"testing"
)

func createEmpty(b *testing.B, wg *sync.WaitGroup) *d2editor.Character {
	c, _ := d2editor.NewEmptyCharacter(97)
	if wg != nil {
		wg.Done()
	}
	return c
}

func openOne(b *testing.B, wg *sync.WaitGroup) *d2editor.Character {
	c, _ := d2editor.Open("LoD-Druid.d2s")
	if wg != nil {
		wg.Done()
	}
	return c
}

func openJSON(b *testing.B, wg *sync.WaitGroup) *d2editor.Character {
	c, _ := d2editor.OpenJSON("LoD-Druid2.json")
	if wg != nil {
		wg.Done()
	}
	return c
}

func saveOne(b *testing.B, c *d2editor.Character, wg *sync.WaitGroup) {
	c.GetBytes()
	if wg != nil {
		wg.Done()
	}
}

func saveJSON(b *testing.B, c *d2editor.Character, wg *sync.WaitGroup) {
	json.Marshal(c)
	if wg != nil {
		wg.Done()
	}
}

func createSaveEmpty(b *testing.B, wg *sync.WaitGroup) {
	saveOne(b, createEmpty(b, nil), nil)
	if wg != nil {
		wg.Done()
	}
}

func openSaveOne(b *testing.B, wg *sync.WaitGroup) {
	saveOne(b, openOne(b, nil), nil)
	if wg != nil {
		wg.Done()
	}
}

func openSaveJSON(b *testing.B, wg *sync.WaitGroup) {
	saveJSON(b, openJSON(b, nil), nil)
	if wg != nil {
		wg.Done()
	}
}

func BenchmarkCreateEmpty(b *testing.B) {
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go createEmpty(b, wg)
	}
	wg.Wait()
}

func BenchmarkSaveEmpty(b *testing.B) {
	wg := &sync.WaitGroup{}
	c := createEmpty(b, nil)
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go saveOne(b, c, wg)
	}
	wg.Wait()
}

func BenchmarkCreateSaveEmpty(b *testing.B) {
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go createSaveEmpty(b, wg)
	}
	wg.Wait()
}

func BenchmarkOpen(b *testing.B) {
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go openOne(b, wg)
	}
	wg.Wait()
}

func BenchmarkSaveFile(b *testing.B) {
	wg := &sync.WaitGroup{}
	c := openOne(b, nil)
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go saveOne(b, c, wg)
	}
	wg.Wait()
}

func BenchmarkOpenSave(b *testing.B) {
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go openSaveOne(b, wg)
	}
	wg.Wait()
}

func BenchmarkOpenJSON(b *testing.B) {
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go openJSON(b, wg)
	}
	wg.Wait()
}

func BenchmarkSaveJSON(b *testing.B) {
	wg := &sync.WaitGroup{}
	c := openJSON(b, nil)
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go saveJSON(b, c, wg)
	}
	wg.Wait()
}

func BenchmarkOpenSaveJSON(b *testing.B) {
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go openSaveJSON(b, wg)
	}
	wg.Wait()
}

func BenchmarkOpenD2sSaveJSON(b *testing.B) {
	c := openOne(b, nil)
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for range make([]bool, b.N) {
		wg.Add(1)
		go saveJSON(b, c, wg)
	}
	wg.Wait()
}
