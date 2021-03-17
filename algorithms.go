package factorize

// Return prime factors as primes: powers
func TrialDivision(n int) map[int]int {
	pf := map[int]int{}
	i := 2
	for n > 1 {
		if n%i == 0 {
			pf[i] += 1
			n /= i
		} else {
			i++
		}
	}
	return pf
}
