package counters

import (
	"testing"
)

func TestConditional_IncrementIf(t *testing.T) {
	var count Conditional

	if count.value != 0 {
		t.Fatal("initial state")
	}

	i := 1
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}

	i++
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}

	i++
	if v, _ := count.IncrementIf(true); v != 0 {
		t.Fatal("expect 0")
	}

	i++
	if v, _ := count.IncrementIf(true); v != 1 {
		t.Fatal("expect 1")
	}

	i++
	if v, _ := count.IncrementIf(true); v != 2 {
		t.Fatal("expect 2")
	}

	i++
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}

	i++
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}
	i++
}
