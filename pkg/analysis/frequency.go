package analysis

import "math"

// Score evaluates a byte slice against a reference frequency map.
// It returns a score representing the distance from the expected distribution.
// A lower score indicates a closer match to the reference text.
func Score(input []byte, freqMap map[byte]float64) float64 {
	// 1. Count the number of occurrences of each letter of the alphabet. Certain characters
	// are later penalized based on their improbability to appear.
	hardPenalty, softPenalty := 0, 0
	occ := make([]int, 26)
	for _, c := range input {
		switch {
		case c > 64 && c < 91:
			occ[int(c)-65]++
		case c > 96 && c < 123:
			occ[int(c)-97]++
		case c < 32 || c > 126:
			hardPenalty++
		case c == ' ' || c == '\'':
			// ignore
		default:
			softPenalty++
		}
	}

	// 2. Compute euclidean distance
	var sum, size float64
	size = float64(len(input))
	for i := range 26 {
		sum += math.Pow((float64(occ[i])/size)*100-freqMap[byte(i+97)], 2)
	}
	return math.Sqrt(sum + 50*float64(hardPenalty) + 20*float64(softPenalty))
}
