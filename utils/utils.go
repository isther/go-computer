package utils

import "fmt"

func Uint16ToString(x uint16) string {
	if x <= 0x000F {
		return fmt.Sprintf("0x000%X", x)
	} else if x <= 0x00FF {
		return fmt.Sprintf("0x00%X", x)
	} else if x <= 0x0FFF {
		return fmt.Sprintf("0x0%X", x)
	}
	return fmt.Sprintf("0x%X", x)
}

func uint16ToBinary(value uint16) [16]bool {
	var (
		b = [16]bool{}
		x = 0
	)

	for i := 16 - 1; i >= 0; i-- {
		r := (value & (1 << uint16(x)))
		if r != 0 {
			b[i] = true
		} else {
			b[i] = false
		}
		x++
	}
	return b
}

func binaryToUint16(b [16]bool) uint16 {
	var (
		value uint16
		x     int
	)
	for i := 16 - 1; i >= 0; i-- {
		if b[i] {
			value = value | (1 << x)
		} else {
			value = value &^ (1 << x)
		}
		x++
	}
	return value
}
