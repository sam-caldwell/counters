package counters

import (
	"testing"
)

func TestIntegerCounter_incrementing(t *testing.T) {

	count := IntegerCounter(0, 1)

	testFunction := func(n int) {
		if count() != n {
			t.Fatalf("testing %d\n", n)
		}
	}

	for i := 1; i <= 100; i++ {
		testFunction(i)
	}
}

func TestIntegerCounter_decrementing(t *testing.T) {

	count := IntegerCounter(10, -1)

	testFunction := func(n int) {
		if v := count(); v != n-1 {
			t.Fatalf("testing %d Got: %d\n", n, v)
		}
	}

	for i := 10; i > 0; i-- {
		testFunction(i)
	}

}
