package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readReports() [][]int {
	var reports = make([][]int, 0)

	input, err := os.OpenFile("day-2/input.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lineContents := scanner.Text()
		splitString := strings.Split(lineContents, " ")

		report := make([]int, 0)

		for _, val := range splitString {
			intVal, _ := strconv.Atoi(val)
			report = append(report, intVal)
		}

		reports = append(reports, report)
	}

	return reports
}
