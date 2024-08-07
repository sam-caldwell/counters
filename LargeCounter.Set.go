package counters

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Set - Set element at pos to value and initialize the c.v if nil
func (c *LargeCounter) Set(pos int, value uint64) (err error) {
	if *c == nil {
		//if not initialized, initialize the data structure
		*c = make([]uint64, pos+1)
	}
	if pos >= len(*c) {
		return fmt.Errorf(errors.OverflowError)
	}
	(*c)[pos] = value
	return err
}
