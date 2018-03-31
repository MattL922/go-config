package config

import (
    "testing"
)

func TestString(t *testing.T) {
    json := "{\"a\":\"b\", \"c\":\"d\"}"
    if err := load([]byte(json)); err != nil {
        t.Fatal(err.Error())
    }
    if val := String("a"); val != "b" {
        t.Fatalf("Expected string value %q for key %q, got %q", "b", "a", val)
    }
    if val := String("c"); val != "d" {
        t.Fatalf("Expected string value %q for key %q, got %q", "d", "c", val)
    }
}
