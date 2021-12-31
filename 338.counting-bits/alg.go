package _38_counting_bits

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		for j := i; j > 0; j &= j - 1 {
			count++
		}
		bits[i] = count
	}

	//for i := 1; i <= n; i++ {
	//	bits[i] = bits[i&(i-1)] + 1
	//	bits[i] = bits[i>>1] + i&1
	//}
	return bits
}
