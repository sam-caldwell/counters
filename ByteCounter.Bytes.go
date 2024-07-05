package counters

// Bytes - return the bytes in the counter.
func (c *ByteCounter) Bytes() []byte {
	c.lock.Lock()
	defer c.lock.Unlock()
	return reverseBytes(c.v)
}
