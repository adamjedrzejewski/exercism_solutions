package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	triplets := []Triplet{}
	for a := min; a <= max; a++ {
		aSquare := math.Pow(float64(a), 2)
		for b := a + 1; b <= max; b++ {
			bSquare := math.Pow(float64(b), 2)
			c := int(math.Sqrt(aSquare + bSquare))
			if c <= max && aSquare+bSquare == math.Pow(float64(c), 2) {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	triplets := Range(1, p)
	result := []Triplet{}

	for _, triplet := range triplets {
		a, b, c := triplet[0], triplet[1], triplet[2]
		if a+b+c == p {
			result = append(result, triplet)
		}
	}

	return result
}
