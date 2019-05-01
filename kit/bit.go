package kit

func BitAt(n int, i uint) int {
	if (n & (1 << i)) == 0 {
		return 0
	} else {
		return 1
	}
}

func BitSet(num int, i uint, bit uint) int {
	if (bit == 0) {
		return num & (^(1 << i))
	} else {
		return num | (1 << i)
	}
}
