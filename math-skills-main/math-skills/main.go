package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go data.txt")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if num, err := strconv.ParseFloat(line, 64); err == nil {
			data = append(data, num)
		}
	}

	if len(data) == 0 {
		fmt.Println("No data found in file.")
		return
	}

	avg := average(data)
	med := median(data)
	variance := variance(data, avg)
	stdDev := math.Sqrt(variance)

	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	fmt.Printf("Median: %d\n", int(math.Round(med)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}

func average(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}

func variance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, v := range data {
		diff := v - mean
		sum += diff * diff
	}
	return sum / float64(len(data))
}
