package session1

import "sort"

func IsAnagramUsingSorting(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Convert string to a slice of runes
	firstSlices := []rune(s)
	secondSlices := []rune(t)

	// Sort the slice
	sort.Slice(firstSlices, func(i, j int) bool {
		return firstSlices[i] < firstSlices[j]
	})

	sort.Slice(secondSlices, func(i, j int) bool {
		return secondSlices[i] < secondSlices[j]
	})

	// Convert the sorted slice back to a string
	sortedFirstStrs := string(firstSlices)
	sortedSecondStrs := string(secondSlices)

	for i := 0; i < len(sortedFirstStrs); i++ {
		if sortedFirstStrs[i] != sortedSecondStrs[i] {
			return false
		}
	}

	return true
}

func IsAnagramUsingMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	slicesS := make(map[rune]int)
	slicesT := make(map[rune]int)

	for _, value := range s {
		slicesS[value]++
	}

	for _, value := range t {
		slicesT[value]++
	}

	for k, v := range slicesS {
		if v != slicesT[k] {
			return false
		}
	}
	return true
}
