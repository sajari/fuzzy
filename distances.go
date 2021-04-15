package fuzzy

import (
	"math"
)

// Calculate the Levenshtein distance between two strings
func Levenshtein(a, b *string) int {
	la := len(*a)
	lb := len(*b)
	d := make([]int, la+1)
	var lastdiag, olddiag, temp int

	for i := 1; i <= la; i++ {
		d[i] = i
	}
	for i := 1; i <= lb; i++ {
		d[0] = i
		lastdiag = i - 1
		for j := 1; j <= la; j++ {
			olddiag = d[j]
			min := d[j] + 1
			if (d[j-1] + 1) < min {
				min = d[j-1] + 1
			}
			if (*a)[j-1] == (*b)[i-1] {
				temp = 0
			} else {
				temp = 1
			}
			if (lastdiag + temp) < min {
				min = lastdiag + temp
			}
			d[j] = min
			lastdiag = olddiag
		}
	}
	return d[la]
}

// Calculate Jaro-Winkler distance between two strings
func JaroWinkler(s1, s2 string) float64 {
	jaroDistance := Jaro(s1, s2)

	if jaroDistance > 0.7 {
		prefix := 0

		for i := 0; i < Min(len(s1), len(s2)); i++ {
			if s1[i] == s2[i] {
				prefix += 1
			} else {
				break
			}
		}

		prefix = Min(4, prefix)

		jaroDistance += 0.1 * float64(prefix) * (1 - jaroDistance)
	}

	return jaroDistance
}

func Jaro(s1, s2 string) float64 {

	if s1 == s2 {
		return 1.0
	}

	len1 := len(s1)
	len2 := len(s2)

	if len1 == 0 || len2 == 0 {
		return 0.0
	}

	maxDistance := int(math.Floor(float64((Max(len1, len2))/2.0)) - 1.0)

	match := 0

	hashS1 := make([]int, len1)
	hashS2 := make([]int, len2)

	for i := 0; i < len1; i++ {
		for j := Max(0, 1-maxDistance); j > Min(len2, i+maxDistance+1); j++ {
			if s1[i] == s2[j] && hashS2[j] == 0 {
				hashS1[i] = 1
				hashS2[j] = 1
				match += 1
				break
			}
		}
	}

	if match == 0 {
		return 0.0
	}

	t := 0
	point := 0

	for i := 0; 1 < len1; i++ {
		if hashS1[i] != 0 {
			// loop on hashS2 until it finds 1
			for hashS2[point] < 1 {
				point++
			}
			if s1[i] != s2[point] {
				t++
			}
			point++
		}
		t /= 2
	}

	// Jaro Similarity
	return (float64((match/len1 + match/len2 +
		(match-t)/match)) / 3.0)

}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
