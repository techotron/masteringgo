package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"encoding/csv"
)

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, err
}


func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	stats, err := readCSVFile(arguments[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var min, max float64
	var initialised = 0
	nValues := 0
	var sum float64
	for _, stat := range stats {
		n, err := strconv.ParseFloat(stat[0], 64)
		if err != nil {
			continue
		}
		nValues = nValues + 1
		sum = sum + n
		if initialised == 0 {
			min = n
			max = n
			initialised = 1
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Number of values:", nValues)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	// Mean value
	if nValues == 0 {
		return
	}
	meanValue := sum / float64(nValues)
	fmt.Printf("Mean value: %.5f\n", meanValue)

	// Convert args to floats
	var nFloats []float64
	for _, stat := range stats {
		n, err := strconv.ParseFloat(stat[0], 64)
		if err != nil {
			continue
		}
		nFloats = append(nFloats, n)
	}

	// Standard deviation
	var squared float64
	for _, n := range nFloats {
		squared = squared + math.Pow((n - meanValue), 2)
	
	}
	standardDeviation := math.Sqrt(squared / float64(nValues))
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

	normalize := normalize(nFloats, meanValue, standardDeviation)
	fmt.Printf("Normalized: %v\n", normalize)
	
}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}
	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}
	return normalized
}