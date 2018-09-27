package utils

func PrimeFactors(n int) []int {

	pfs := make([]int, 0)

	for i := 2; n > 1; i++ {
		if n % i == 0 {
			pfs = append(pfs, i)
			n /= i
		}
	}

	return pfs
}