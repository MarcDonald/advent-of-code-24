package day1

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var listOne = make([]int, 0)
var listTwo = make([]int, 0)
var distances = make([]int, 0)

func Part1() int {
	populateLists()

	sort.Slice(listOne, func(i2, j int) bool {
		return listOne[i2] < listOne[j]
	})

	sort.Slice(listTwo, func(i2, j int) bool {
		return listTwo[i2] < listTwo[j]
	})

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

func populateLists() {
	input, err := os.OpenFile("day-1/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lineContents := scanner.Text()
		splitString := strings.Split(lineContents, "   ")

		int1, _ := strconv.Atoi(splitString[0])
		int2, _ := strconv.Atoi(splitString[1])

		listOne = append(listOne, int1)
		listTwo = append(listTwo, int2)
	}
}
