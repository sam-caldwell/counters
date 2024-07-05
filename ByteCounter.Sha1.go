package counters

import (
	"crypto/sha1"
	"encoding/hex"
)

// Sha1 - create byte counter hash
func (c *ByteCounter) Sha1() string {
	c.lock.Lock()
	defer c.lock.Unlock()
	hash := sha1.Sum(c.v)
	return hex.EncodeToString(hash[:])
}

// Sha1Bytes - Return the sha1 sum
func (c *ByteCounter) Sha1Bytes() [20]byte {
	c.lock.Lock()
	defer c.lock.Unlock()
	return sha1.Sum(c.v)
}
