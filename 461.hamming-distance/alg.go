package _61_hamming_distance

func hammingDistance(x int, y int) int {
	n := x ^ y
	count := 0
	for ; n != 0; n &= n - 1 {
		count++
	}
	return count
}
