package counters

// reverse - reverse the counter bytes
func (c *ByteCounter) reverse() {
	sz := len(c.v)
	rhs := sz
	sz /= 2
	for lhs := 0; lhs < sz; lhs++ {
		rhs--
		c.v[lhs] = c.v[rhs]
	}
}
