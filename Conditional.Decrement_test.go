package counters

import (
	"testing"
)

func TestConditional_Decrement(t *testing.T) {

	var count Conditional

	if count.value != 0 {
		t.Fatalf("initial state")
	}

	for i := 0; i < 4; i++ {
		if v, _ := count.Increment(); v != i {
			t.Fatalf("expect %d Got:%d", i, v)
		}
	}

	for i := 4; i > 0; i-- {
		if v, _ := count.Increment(); v != i {
			t.Fatalf("expect %d Got:%d", i, v)
		}
	}

	if v, _ := count.Decrement(); v != 0 {
		t.Fatalf("expect 0 Got:%d", v)
	}

}
