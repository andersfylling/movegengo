package movegengo

import "math/bits"

func LSB(x uint64) int {
	return bits.TrailingZeros64(x)
}

func NLSB(x *uint64, i int) int {
	*x ^= uint64(1) << uint64(i)
	return LSB(*x)
}
