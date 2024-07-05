package counters

import (
	"fmt"
)

func NewByteCounter(sz int) (*ByteCounter, error) {
	var b ByteCounter
	if sz <= 0 {
		return &b, fmt.Errorf("ByteCounter size must be > 0")
	}
	b.v = make([]byte, sz)
	return &b, nil
}
