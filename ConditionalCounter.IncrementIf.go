package counters

// IncrementIf - increment if ok is true otherwise return -1
func (counter *ConditionalCounter) IncrementIf(ok bool) int {
	const notOK = -1 //If not ok, return -1
	if ok {
		//We increment after we return the value (if ok)
		defer func() {
			counter.value++
		}()
		return int(counter.value) //if ok, return state
	}
	return notOK
}
