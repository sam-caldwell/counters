package counters

import "testing"

func TestLargeCounter(t *testing.T) {
	var o LargeCounter
	if o != nil {
		t.Fatal("expect nil initial state")
	}
}
