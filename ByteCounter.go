package counters

import "sync"

// ByteCounter - create an array of n bytes and allow the count to be incremented
//
//	ByteCounter will create an array of n bytes and allow the count to be incremented 0...255
//	for each byte, carrying 1 to the n+1 byte.
type ByteCounter struct {
	lock sync.Mutex
	v    []byte // byte array used for the counter.
}
