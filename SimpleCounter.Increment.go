package counters

func (counter *SimpleCounter) Increment() int {
	defer func() {
		counter.value++
	}()
	return counter.value
}
