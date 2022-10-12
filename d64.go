package msd

import "math"

type D64 = uint64

func D64Value(d D64) float64 {
	n := int64(d >> 8)
	dec := int((d & 0xF0) >> 4)
	if d&0x01 == 1 {
		n = -n
	}
	if dec == 0 {
		return float64(n)
	} else {
		return float64(n) / math.Pow10(dec)
	}
}
