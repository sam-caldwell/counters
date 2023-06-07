package counters

func (counter *ConditionalCounter) Value() int {
	return int(counter.value)
}
