package factorize

import (
	"errors"
	"math"
	"math/rand"
)

// Simple trial division with 6k±1 optimization detailed here
// https://en.wikipedia.org/wiki/Primality_test#Simple_methods
func simple(n int) (bool, error) {
	if n <= 3 {
		return n > 1, nil
	}

	if (n%2 == 0) || (n%3) == 0 {
		return false, nil
	}

	i := 5
	for i*i <= n {
		if (n%i == 0) || ((i + 2) == 0) {
			return false, nil
		}
		i += 6
	}
	return true, nil
}

// Selfridge conjecture, simplified.
// n must be odd. n must be congruent to ±2 (mod 5).
// Note that this is an unproven conjecture.
func selfridge(n int) (bool, error) {
	if (n%2 == 0) || (math.Abs(float64(n%5)) != 2.0) {
		return false, errors.New("selfridge: n must be odd and congruent to ±2 (mod 5).")
	}

	cond1 := int(math.Pow(2, float64(n-1)))%n == 1 // Fermat primality test
	cond2 := GetNthFibonacci(n+1)%n == 0

	if cond1 && cond2 {
		return true, nil
	}

	return false, nil
}

// Fermat Primality Test
// Note this is a probabilistic test, where n is *probably* prime
func fermat(n int) (bool, error) {
	a := rand.Intn(n-1) + 1
	if (int(math.Pow(float64(a), float64(n-1))) % n) == 1 {
		// Check another random a value to be sure that a is not a "Fermat Liar"
		a2 := rand.Intn(n-1) + 1
		for a2 == a {
			a2 = rand.Intn(n-1) + 1
		}

		// If the primality test passes again for a2 then it is probably prime, otherwise composite
		if (int(math.Pow(float64(a2), float64(n-1))) % n) == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, nil
}

var primalityMethods map[string]func(int) (bool, error) = map[string]func(int) (bool, error){
	"simple":    simple,
	"selfridge": selfridge,
	"fermat":    fermat,
}

// Test if an integer is prime. The user may
func IsPrime(n int, method ...string) (bool, error) {
	if len(method) > 0 {
		if len(method) != 1 {
			return false, errors.New("IsPrime: more than one method specified. Choose from [simple, ...]")
		}

		return primalityMethods[method[0]](n)
	}

	return simple(n)
}
