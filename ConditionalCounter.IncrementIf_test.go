package counters

import "testing"

func TestConditionalCounter_IncrementIf(t *testing.T) {
	var count ConditionalCounter
	if count.value != 0 {
		t.Fail()
	}
	if count.IncrementIf(false) != -1 {
		t.Fatal("expect -1")
	}
	if count.IncrementIf(false) != -1 {
		t.Fatal("expect -1")
	}
	if count.IncrementIf(true) != 0 {
		t.Fatal("expect 0")
	}
	if count.IncrementIf(true) != 1 {
		t.Fatal("expect 1")
	}
	if count.IncrementIf(true) != 2 {
		t.Fatal("expect 2")
	}
	if count.IncrementIf(false) != -1 {
		t.Fatal("expect -1")
	}
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}
}
