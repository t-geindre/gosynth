package normalizer

import "math"

const freqOctaves float64 = 4 // +-4 octaves
const freqRef float64 = 440   // std A4

// CvToFrequency Normalize a frequency value between -1 and 1 to hz
func CvToFrequency(value float64) float64 {
	return freqRef * math.Pow(2, value*freqOctaves)
}

// FrequencyToCv Denormalize a frequency (hz) to -1 and 1
func FrequencyToCv(value float64) float64 {
	return math.Log2(value/freqRef) / freqOctaves
}
