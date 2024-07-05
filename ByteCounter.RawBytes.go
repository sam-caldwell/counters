package counters

// RawBytes - return the raw bytes of the counter.
func (c *ByteCounter) RawBytes() []byte {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.v
}
