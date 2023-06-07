package counters

// Increment - Increment value
func (counter *ConditionalCounter) Increment() int {
	defer func() {
		if counter.value > uint(int(^uint(0)>>1)) {
			panic("ConditionalCounter exceeds limit")
		}
		counter.value++

	}()
	return int(counter.value) //if ok, return state
}
