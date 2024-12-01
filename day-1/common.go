package day1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func populateLists() ([]int, []int) {
	var listOne = make([]int, 0)
	var listTwo = make([]int, 0)

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

	return listOne, listTwo
}
