package factorize

// Generate the Fibonacci sequence up to n places
func GenerateFibonacci(n int) []int {
	var seq []int = []int{0}
	i := 0
	j := 1
	step := 1
	for step < n {
		seq = append(seq, j)
		t := i + j
		i = j
		j = t
		step++
	}

	return seq
}

func GetNthFibonacci(n int) int {
	i := 0
	j := 1
	step := 1
	for step < n {
		t := i + j
		i = j
		j = t
		step++
	}

	return i
}
