package counters

import (
	"bytes"
	"testing"
)

func TestByteCounter_RawBytes(t *testing.T) {
	t.Run("Setup: Create counter (10 elements)", func(t *testing.T) {
		b, err := NewByteCounter(10)
		if err != nil {
			t.Fatal("expect no error")
		}
		if len(b.v) != 10 {
			t.Fatal("expect 10-element array")
		}
	})
	t.Run("Happy: verify initial state", func(t *testing.T) {
		b, _ := NewByteCounter(10)
		expected := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		if !bytes.Equal(b.v, expected) {
			t.Fatal("expect empty array")
		}
	})
	t.Run("Happy: create state of all 10 elements and verify .Bytes() returns expected output", func(t *testing.T) {
		b, _ := NewByteCounter(10)
		b.v = []byte{0x0A, 0x09, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01}
		actual := b.RawBytes() //should flip byte order
		expected := []byte{0x0A, 0x09, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("output mismatch(1.0)\n"+
				"a:%v\n"+
				"e:%v", actual, expected)
		}
		if (string)(actual) != (string)(expected) {
			t.Fatalf("output mismatch(1.1)")
		}
	})
	t.Run("Happy: Increment counter and confirm .Bytes() returns expected output", func(t *testing.T) {
		b, _ := NewByteCounter(10)
		expected := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		_ = b.Increment()
		actual := b.RawBytes()
		if !bytes.Equal(actual, expected) {
			t.Fatalf("output mismatch(2.0)\na:%v\ne:%v", actual, expected)
		}
		if (string)(actual) != (string)(expected) {
			t.Fatalf("output mismatch (2.1)")
		}
	})

}
