package randname

import (
	"time"
)

const (
	MAX_UINT64 = ^uint64(0)
	MAX_INT64  = ^int64(0)
)

var seed uint64 = 1234
var init_seed uint64

func init() {
	init_seed = TimeSeed()
	seed = init_seed
}

func InitSeed() uint64 {
	return init_seed
}

func SeedRand(s uint64) uint64 {
	seed = s
	return seed
}

func TimeSeed() uint64 {
	return uint64((time.Now().UnixNano() % int64(MAX_INT64)))
}

func GenRand() uint64 {
	seed ^= (seed << 7)
	seed = (seed >> 1)
	seed ^= (seed << 9)
	return seed
}

func RandIntn(n int) int {
	return int(GenRand()) % n
}

func RandUint64n(n uint64) uint64 {
	return GenRand() % n
}
