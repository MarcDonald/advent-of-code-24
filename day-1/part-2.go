package day1

func Part2() int {
	listOne, listTwo := populateLists()

	totalSimilarity := 0

	for _, val := range listOne {
		numberOfRepeats := 0
		for i := 0; i < len(listTwo); i++ {
			if listTwo[i] == val {
				numberOfRepeats++
			}
		}

		similarityScore := val * numberOfRepeats
		totalSimilarity += similarityScore
	}

	return totalSimilarity
}
