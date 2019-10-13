package main

import (
	"math"
	"math/big"
	"crypto/rand"
)

var sampleMaxCache = make(map[int]map[int]int64)
var sampleCache = make(map[int]map[int]int)

type sampling struct {
	max int64
	size int
}

// sampleMax computes the number of samples to get from a single call to rand.Int() to minimize rejection.
func sampleMax(base int, size int) sampling {
	if m, ok := sampleMaxCache[base][size]; ok {
		return sampling{m, sampleCache[base][size]}
	}
	//fmt.Printf("base\tsamples\texpected_rejections\n")
	bestMax := int64(1)
	bestSampling := 1
	minRejections := math.Inf(1)
	var max float64
	for samples := 1; samples <= size && max < 1 << 50; samples++ {
		max := math.Pow(float64(base), float64(samples))
		collision := 1 - max / (math.Pow(2, math.Ceil(math.Log2(max))))
		expectedRejections := collision * math.Ceil(float64(size/samples))
		if expectedRejections < minRejections {
			bestMax = int64(max)
			bestSampling = samples
			minRejections = expectedRejections
		}
		//fmt.Printf("%d\t%d\t%.3f\n", base, samples, expectedRejections)
	}
	if _, ok := sampleMaxCache[base]; !ok {
		sampleMaxCache[base] = make(map[int]int64)
		sampleCache[base] = make(map[int]int)
	}
	sampleMaxCache[base][size] = bestMax
	sampleCache[base][size] = bestSampling
	return sampling{bestMax, bestSampling}
}

func sample(base int, size int) string {
	s := ""	
	sampling := sampleMax(base, size)
	for i := 0; i < sampling.size; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(sampling.max))
		s += string(n.String())
	}
	return s[:size]
}

func rawSample(base int, size int) string {
	s := ""
	for i := 0; i < size; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(base)))
		s += string(n.String())
	}
	return s
}
