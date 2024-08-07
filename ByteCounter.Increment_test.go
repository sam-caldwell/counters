package counters

import (
	"testing"
	"time"
)

func TestByteCounter_Increment(t *testing.T) {
	var err error
	var b *ByteCounter

	t.Run("Happy: Create an initialized ByteCounter (10) and verify initial state", func(t *testing.T) {
		if b, err = NewByteCounter(10); err != nil {
			t.Fatalf("expected ByteCounter with no errors")
		}
		if len(b.v) != 10 {
			t.Fatalf("expected 10-element byte array")
		}
		for i := 0; i < 10; i++ {
			//t.Logf("i: %d", i)
			if b.v[i] != 0 {
				t.Fatalf("expect element %d to be zero", i)
			}
		}
	})

	t.Run("Happy: increment byte 0 to 255 and confirm state", func(t *testing.T) {
		for v := byte(1); v < byte(255); v++ {
			_ = b.Increment()
			if b.v[0] != v {
				t.Fatal("expected zeroth element incremented")
			}
			if b.v[0] != v {
				t.Fatalf("expect element 0 to be %d", v)
			}
			for i := 1; i < 10; i++ {
				if b.v[i] != 0 {
					t.Fatalf("expect element %d to be 0", i)
				}
			}
		}
	})

	t.Run("Happy: Set byte 0 to 255 and test carry to byte 1 if incremented", func(t *testing.T) {
		b.v[0] = 255
		_ = b.Increment()
		if b.v[0] != 0 {
			t.Fatalf("expect element 0 to rollover to 0")
		}
		if b.v[1] != 1 {
			t.Fatalf("expect element 1 to carry over to 1")
		}
	})

	t.Run("Happy: Set bytes 0, 1 to 255 and test carry if byte 0 is incremented.", func(t *testing.T) {
		b.v[0] = 255
		b.v[1] = 255
		_ = b.Increment()
		if b.v[0] != 0 {
			t.Fatalf("expect element 0 to rollover to 0")
		}
		if b.v[1] != 0 {
			t.Fatalf("expect element 1 to carry and rollover to 0")
		}
		if b.v[2] != 1 {
			t.Fatalf("expect element 2 to carry over to 1")
		}
	})

	t.Run("Sad: Set all bytes to 255 and increment, expecting an over-flow error", func(t *testing.T) {
		for i := 0; i < len(b.v); i++ {
			b.v[i] = 255
		}
		_ = b.Increment()

		func() {
			const iterations = 10485760
			for pass := 0; pass < 100; pass++ {
				startTime := time.Now().UnixNano()
				c, err := NewLargeCounter(64 * 10)
				if err != nil {
					t.Fatal(err)
				}
				for i := 0; i < iterations; i++ {
					c.Increment()
				}
				if (*c)[0] != iterations {
					t.Fatal("outcome unexpected")
				}
				stopTime := time.Now().UnixNano()
				elapsedPerIteration := float64(stopTime-startTime) / float64(iterations)
				if elapsedPerIteration > 30 {
					t.Fatalf("baseline performance not expected (%f) ns/iteration", elapsedPerIteration)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}()
	})
}
