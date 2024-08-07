package counters

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

func TestByteCounter_Sha1(t *testing.T) {
	hashFunc := func(i []byte) string {
		hash := sha1.Sum(i)
		return hex.EncodeToString(hash[:])
	}
	t.Run("Happy path: empty hash", func(t *testing.T) {
		var b ByteCounter
		if b.Sha1() != hashFunc([]byte{}) {
			t.Fatal("sha1 hash of empty state fails")
		}
	})
	t.Run("Happy path: test hash of non-nil set", func(t *testing.T) {
		input := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A}
		var b ByteCounter
		b.v = input
		if b.Sha1() != hashFunc(input) {
			t.Fatal("mismatch")
		}
	})
}
