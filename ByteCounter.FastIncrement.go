package counters

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// FastIncrement - increment the large counter by one (without locking)
func (c *ByteCounter) FastIncrement() error {
	for pos := 0; pos < len(c.v); pos++ {
		c.v[pos]++
		if c.v[pos] != 0 {
			return nil
		}
	}
	return fmt.Errorf(errors.OverflowError)
}
