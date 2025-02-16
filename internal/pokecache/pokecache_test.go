package pokecache

import (
	"fmt"
	"testing"
	"time"
)


func TestAddGet(t *testing.T) {
	cases := []struct{
		keyInput string
		data []byte
	}{
		{
			keyInput: "https://example.com/1",
			data: []byte("Hello world!"),
		},
		{
			keyInput: "https://example.com/2",
			data: []byte("Foo bar!"),
		},
	}
	
	for i, kase := range cases {
		t.Run(fmt.Sprintf("Running test case %v", i), func(t *testing.T) {
			cache := NewCache(5 * time.Second)
			cache.Add(kase.keyInput, kase.data)

			val, ok := cache.Get(kase.keyInput)
			if !ok {
				t.Errorf("Expected to find key '%v'", kase.keyInput)
				return
			}
			if string(val) != string(kase.data) {
				t.Errorf("Expected data to be '%v', got '%v' instead", string(kase.data), string(val))
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	interval := 3 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("hello world!"))
	data, ok := cache.Get("https://example.com")

	if !ok {
		t.Errorf("Expected data '%v', but not found", string(data))
		return
	}

	time.Sleep(4 * time.Millisecond)
	newData, newOk := cache.Get("https://example.com")
	if newOk {
		t.Errorf("Cached data '%v' should be removed but it still exists", string(newData))
		return
	}

}