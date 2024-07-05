package counters

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Increment - increment the large counter by one.
func (c *ByteCounter) Increment() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	for i := 0; i < len(c.v); i++ {
		c.v[i]++
		if c.v[i] != 0 {
			return nil
		}
	}
	return fmt.Errorf(errors.OverflowError)
}
