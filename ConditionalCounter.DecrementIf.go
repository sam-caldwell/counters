package counters

// DecrementIf - decrement to floor (0) if true otherwise return -1
func (counter *ConditionalCounter) DecrementIf(ok bool) int {
	const notOK = -1 //If not ok, return -1
	if ok {
		//We increment after we return the value (if ok)
		//We will do nothing at the floor (0)
		defer func() {
			if counter.value > 0 {
				counter.value--
			}
		}()
		return int(counter.value) //if ok, return state
	}
	return notOK
}
