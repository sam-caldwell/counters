package counters

func (counter *SimpleCounter) Decrement() int {
	defer func() {
		counter.value--
	}()
	return counter.value
}
