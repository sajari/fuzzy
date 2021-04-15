package fuzzy

import (
	"testing"
)

func TestLevshtein(t *testing.T) {
	s1 := "hello"
	s2 := "hollaaaa"
	lev := Levenshtein(&s1, &s2)

	if lev != 5 {
		t.Errorf("Lev %v", lev)
	}
}

func TestJaro(t *testing.T) {
	s1 := "hello"
	s2 := "hollaaaa"
	j := Jaro(s1, s2)

	if j != 0.6833333333333332 {
		t.Errorf("J %v", j)
	}
}

func TestJaroWinkler(t *testing.T) {
	s1 := "LATE"
	s2 := "LACE"
	jw := JaroWinkler(s1, s2)

	if jw != 0.8666666666666667 {
		t.Errorf("JW %v", jw)
	}
}
