package counters

import (
	"math/big"
)

// Int - return the numeric value of our counter state (using big int)
func (c *ByteCounter) Int() *big.Int {

	c.lock.Lock()
	defer c.lock.Unlock()

	var i big.Int
	i.SetBytes(reverseBytes(c.v))
	return &i
}
