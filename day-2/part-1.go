package day2

import (
	"math"
)

func Part1() int {
	reports := readReports()

	var safeReports = 0

	for _, report := range reports {
		var isSafe = true
		var prevVal = -1
		var ascOrDesc = ""

		for num, val := range report {
			if isSafe {
				if num == 1 {
					if val > prevVal {
						ascOrDesc = "asc"
					} else if val < prevVal {
						ascOrDesc = "desc"
					} else {
						isSafe = false
					}
				}

				if prevVal != -1 {
					if (ascOrDesc == "asc" && val <= prevVal) || (ascOrDesc == "desc" && val >= prevVal) {
						isSafe = false
					} else if math.Abs(float64(val-prevVal)) > 3 {
						isSafe = false
					}
				}
				prevVal = val
			} else {
				break
			}
		}

		if isSafe {
			safeReports++
		}
	}

	return safeReports
}
