package util

import (
	"math/rand"
)

func GetRandomUserID() uint64 {
	return GetRandomUnsignedInteger(1, 101)
}

func GetRandomMessage(len int) string {
	return GetRandomString(len)
}

func GetRandomString(len int) string {
	res := make([]byte, len)

	for i := 0; i < len; i++ {
		res[i] = byte(GetRandomUnsignedInteger(65, 91))
	}
	return string(res)
}

// return a unsigned integer in [min, max)
func GetRandomUnsignedInteger(min, max uint64) uint64 {
	return min + rand.Uint64()%(max-min)
}
