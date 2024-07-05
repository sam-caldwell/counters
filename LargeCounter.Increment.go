package counters

// Increment - increment the large counter by one.
func (c *LargeCounter) Increment() {
	for i := 0; i < len(*c); i++ {
		(*c)[i]++
		if (*c)[i] != 0 {
			break
		}
	}
	return
}
