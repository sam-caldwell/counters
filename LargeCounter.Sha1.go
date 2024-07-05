package counters

import (
	"crypto/sha1"
	"encoding/hex"
)

// Sha1 - Return a sha1 hash of the LargeCounter value
func (c *LargeCounter) Sha1() string {
	hash := sha1.Sum(c.Bytes())
	return hex.EncodeToString(hash[:])
}

func (c *LargeCounter) Sha1Bytes() [20]byte {
	return sha1.Sum(c.Bytes())
}

func (c *LargeCounter) Sha1Fast() [20]byte {
	return sha1.Sum(c.FastBytes())
}
