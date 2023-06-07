package counters

// IncrementIf - increment if ok is true otherwise return -1
func (counter *ConditionalCounter) IncrementIf(ok bool) int {
	const notOK = -1 //If not ok, return -1
	defer func() {
		if counter.value > uint(int(^uint(0)>>1)) {
			panic("ConditionalCounter exceeds limit")
		}
		counter.value++

	}()
	return notOK
}
