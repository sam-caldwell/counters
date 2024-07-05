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
		t.Fatal("step 1")
	}
	if count.Value() != 1 {
		t.Fatal("step 2")
	}
	t.Fatal("step 2")
}
