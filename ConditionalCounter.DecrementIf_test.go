package counters

import "testing"

func TestConditionalCounter_DecrementIf(t *testing.T) {
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
	if v := count.DecrementIf(false); v != -1 {
		t.Fatalf("5.expect -1 Got:%d", v)
	}
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}
	if v := count.DecrementIf(false); v != -1 {
		t.Fatalf("6.expect -1 Got:%d", v)
	}
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}
	if v := count.DecrementIf(true); v != 3 {
		t.Fatalf("7.expect 3 Got:%d", v)
	}
	if v := count.DecrementIf(true); v != 2 {
		t.Fatalf("8.expect 2 Got:%d", v)
	}
	if v := count.DecrementIf(true); v != 1 {
		t.Fatalf("9.expect 1 Got:%d", v)
	}
	if v := count.DecrementIf(true); v != 0 {
		t.Fatalf("10.expect 0 Got:%d", v)
	}
	if v := count.DecrementIf(true); v != 0 {
		t.Fatalf("10.expect 0 Got:%d", v)
	}
}
