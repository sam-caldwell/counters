package counters

import (
	"testing"
)

func TestByteCounter_Struct(t *testing.T) {
	var b ByteCounter
	if b.v != nil {
		t.Fatal("b.v should be nil")
	}
}
