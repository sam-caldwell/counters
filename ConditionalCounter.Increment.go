package counters

// Increment - Increment value
func (counter *ConditionalCounter) Increment() int {
	defer func() {
		counter.value++
	}()
	return int(counter.value) //if ok, return state
}
