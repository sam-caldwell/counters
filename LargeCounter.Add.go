package counters

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Add - Add value to the LargeCounter.
func (c *LargeCounter) Add(value uint64) error {
	return c.carryValue(0, value)
}

// carryValue - Add the value to an element at pos and carry the overflow.
func (c *LargeCounter) carryValue(pos uint, v uint64) error {
	for pos < uint(len(*c)) && v > 0 {
		sum := (*c)[pos] + v
		if sum < (*c)[pos] { // Check for overflow
			v = 1
		} else {
			v = 0
		}
		(*c)[pos] = sum
		pos++
	}
	if v > 0 {
		return fmt.Errorf(errors.OverflowError)
	}
	return nil
}
