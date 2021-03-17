package factorize

import (
	"reflect"
	"testing"
)

func TestTrialDivision(t *testing.T) {
	pf := TrialDivision(12)
	expected := map[int]int{2: 2, 3: 1}
	if eq := reflect.DeepEqual(pf, expected); !eq {
		t.Fatalf("Got: %v, Expected: %v", pf, expected)
	}
}

func BenchmarkTrialDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrialDivision(12)
	}
}
