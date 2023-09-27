package shadiff

import (
	"crypto/sha256"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(b byte) int {
	return int(pc[b])
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}
	}
	return count
}

func ShaCountBitDiff(a, b []byte) int {
	sha1 := sha256.Sum256(a)
	sha2 := sha256.Sum256(b)
	return bitDiff(sha1[:], sha2[:])
}
