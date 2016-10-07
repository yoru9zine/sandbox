package hashmap

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	h := NewHashMap(2)

	// add new key
	h.Set("key", "val")
	got := h.Get("key").(string)
	expected := "val"
	if got != expected {
		t.Errorf(`failed to get value for "key": got=%s expected=%s`, got, expected)
	}

	// add other new key
	h.Set("key2", "val2")
	got = h.Get("key2").(string)
	expected = "val2"
	if got != expected {
		t.Errorf(`failed to get value for "key": got=%s expected=%s`, got, expected)
	}

	// overwrite key
	h.Set("key", "val2")
	got = h.Get("key").(string)
	expected = "val2"
	if got != expected {
		t.Errorf(`failed to get value for "key": got=%s expected=%s`, got, expected)
	}

	for i := 0; i < 10000; i++ {
		h.Set(fmt.Sprint(i), "value")
	}
}
