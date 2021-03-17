package factorize

import "testing"

func TestIsPrime(t *testing.T) {
	isPrime, err := IsPrime(120)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if isPrime {
		t.Fatal("120 is not prime")
	}

	isPrime, err = IsPrime(221)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if !isPrime {
		t.Fatal("59 is prime")
	}
}

func TestIsPrimeFermat(t *testing.T) {
	isPrime, err := IsPrime(221, "fermat")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if isPrime {
		t.Fatal("221 is not prime")
	}

	isPrime, err = IsPrime(59)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if !isPrime {
		t.Fatal("59 is prime")
	}
}

func TestIsPrimeSelfridge(t *testing.T) {
	isPrime, err := IsPrime(227, "selfridge")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if isPrime {
		t.Fatal("221 is not prime")
	}

	isPrime, err = IsPrime(59)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if !isPrime {
		t.Fatal("59 is prime")
	}
}
