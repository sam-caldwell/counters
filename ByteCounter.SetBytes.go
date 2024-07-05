package counters

import (
	"fmt"
)

// SetBytes - Set the ByteCounter value to a byte array from a given offset (pos).
func (c *ByteCounter) SetBytes(pos int, values []byte) (err error) {
	if c.v == nil {
		//If not initialized, initialize it to avoid bad things.
		c.v = make([]byte, len(values)+pos)
	}
	if pos+len(values) >= len(c.v) {
		//If our values array would overflow c.v, bail!
		err = fmt.Errorf("overflow error")
	} else {
		for i := 0; i < len(values); i++ {
			//copy values[] int c.v[]
			c.v[pos+i] = values[i]
		}
	}
	return err
}

// Reset - clear the counter state
func (c *ByteCounter) Reset() {
	for i := 0; i < len(c.v); i++ {
		c.v[i] = 0x00
	}
}

// Revert - revert the counter to original state
func (c *ByteCounter) Revert(originalState *[]byte) {
	copy(c.v, *originalState)
}
