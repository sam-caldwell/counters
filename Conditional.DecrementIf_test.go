package counters

import (
	"testing"
)

func TestConditional_DecrementIf(t *testing.T) {
	var count Conditional
	if count.value != 0 {
		t.Fail()
	}
	i := 0

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
	if v, _ := count.DecrementIf(false); v != -1 {
		t.Fatalf("5.expect -1 Got:%d", v)
	}

	i++
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}

	i++
	if v, _ := count.DecrementIf(false); v != -1 {
		t.Fatalf("6.expect -1 Got:%d", v)
	}

	i++
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}

	i++
	if v, _ := count.DecrementIf(true); v != 3 {
		t.Fatalf("7.expect 3 Got:%d", v)
	}

	i++
	if v, _ := count.DecrementIf(true); v != 2 {
		t.Fatalf("8.expect 2 Got:%d", v)
	}

	i++
	if v, _ := count.DecrementIf(true); v != 1 {
		t.Fatalf("9.expect 1 Got:%d", v)
	}

	i++
	if v, _ := count.DecrementIf(true); v != 0 {
		t.Fatalf("10.expect 0 Got:%d", v)
	}

	i++
	if v, _ := count.DecrementIf(true); v != -1 {
		t.Fatalf("10.expect -1 Got:%d", v)
	}

	i++
}
