package counters

import (
	"encoding/hex"
	"strings"
)

// String - return the string value of our counter state.
func (c *ByteCounter) String() string {
	c.lock.Lock()
	defer c.lock.Unlock()
	return strings.ToUpper(hex.EncodeToString(reverseBytes(c.v)))
}
