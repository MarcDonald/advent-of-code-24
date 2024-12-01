package day1

import (
	"math"
	"sort"
)

func Part1() int {
	listOne, listTwo := populateLists()

	sort.Slice(listOne, func(i2, j int) bool {
		return listOne[i2] < listOne[j]
	})

	sort.Slice(listTwo, func(i2, j int) bool {
		return listTwo[i2] < listTwo[j]
	})

	var distances = make([]int, 0)

	for i, value := range listOne {
		distanceBetween := math.Abs(float64(listTwo[i] - value))
		distances = append(distances, int(distanceBetween))
	}

	totalDistance := 0

	for _, val := range distances {
		totalDistance += val
	}

	return totalDistance
}
