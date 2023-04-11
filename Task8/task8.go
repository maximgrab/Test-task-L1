package main

import (
	"encoding/binary"
	"fmt"
)

func setBit(val *int64, bit bool, pos int) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(*val))
	b[pos/8] &= ^(1 << (-1 + 8 - (pos % 8)))

	if bit {
		b[pos/8] |= 1 << (-1 + 8 - (pos % 8))
	}
	*val = int64(binary.LittleEndian.Uint64(b))
	fmt.Printf("%08b\n", b)
}
func main() {
	n := int64(5)
	setBit(&n, true, 63)
	setBit(&n, true, 0)
	setBit(&n, false, 5)
}
