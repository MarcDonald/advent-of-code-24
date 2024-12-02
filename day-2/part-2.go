package day2

import (
	"fmt"
	"math"
	"strings"

	"marcdonald.com/advent-of-code/common"
)

func Part2() int {
	reports := readReports()

	var safeReports = 0

	for _, report := range reports {
		firstRunSafe, badIndex := checkReport(report)

		if firstRunSafe {
			safeReports++
		} else {
			secondRunSafe, _ := checkReport(common.RemoveFromSlice(report, badIndex))
			if secondRunSafe {
				safeReports++
			}
		}
	}

	return safeReports
}

func checkReport(report []int) (isSafe bool, badIndex int) {
	var prevVal = -1
	var ascOrDesc = ""

	for index, val := range report {
		if prevVal != -1 {
			// If the previous value is the same as the current value, then it's not safe
			if prevVal == val {
				return false, index
			}

			// Determine if ascending or descending
			if ascOrDesc == "" {
				if val > prevVal {
					ascOrDesc = "asc"
				} else if val < prevVal {
					ascOrDesc = "desc"
				} else {
					// If the values are the same, it's not safe
					return false, index
				}
			}

			// Ensure the sequence is consistent (either ascending or descending)
			if (ascOrDesc == "asc" && val <= prevVal) || (ascOrDesc == "desc" && val >= prevVal) {
				return false, index
			}

			// If the difference between the current value and the previous value is less than 1 or greater than 3, it's not safe
			diff := math.Abs(float64(val - prevVal))
			if diff < 1 || diff > 3 {
				return false, index
			}
		}

		prevVal = val
	}

	println("Report is safe: " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(report)), ","), "[]"))
	return true, -1
}
