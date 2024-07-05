package counters

import "encoding/hex"

// String - Return hex string
func (c *LargeCounter) String() string {
	return hex.EncodeToString(c.Bytes())
}
