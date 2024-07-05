package counters

import (
	"testing"
)

func TestByteCounter_String(t *testing.T) {
	t.Run("Happy: Create an initialized ByteCounter (10) and verify initial state", func(t *testing.T) {
		b, err := NewByteCounter(10)
		if err != nil {
			t.Fatal("expect no error")
		}
		if len(b.v) != 10 {
			t.Fatal("expect 10-element array")
		}
		expected := "00000000000000000000"
		if actual := b.String(); actual != expected {
			t.Fatal("expect zero string")
		}
	})
	t.Run("Happy: Set the internal counter state and call .String() and verify the result.", func(t *testing.T) {
		b, err := NewByteCounter(10)
		if err != nil {
			t.Fatal("expect no error")
		}
		b.v = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A}
		t.Logf("b.v: %v", b.v)
		expected := "0A090807060504030201"
		actual := b.String() //should flip byte order
		if actual != expected {
			t.Fatalf("output mismatch(1)\na:'%s'\ne:'%s'", actual, expected)
		}
	})
	t.Run("Happy: Increment the counter and call .String() then verify the result.", func(t *testing.T) {
		b, err := NewByteCounter(10)
		if err != nil {
			t.Fatal("expect no error")
		}
		b.v = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A}
		expected := "0A090807060504030202" //increment element(0)
		_ = b.Increment()
		actual := b.String()
		if actual != expected {
			t.Fatalf("output mismatch(2)\na:%v\ne:%v", actual, expected)
		}
	})
}
