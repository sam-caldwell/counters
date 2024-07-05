package counters

import (
	"testing"
)

func TestConditional_Value(t *testing.T) {
	var count Conditional
	if count.Value() != 0 {
		t.Fatal("initial state")
	}
	if v, _ := count.Increment(); v != 0 {
		t.Fatalf("step 1")
	}
	if v := count.Value(); v != 1 {
		t.Fatalf("step 2 v=%d", v)
	}
}
