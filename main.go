package main

import (
	"duplicateExcercise/exercise"
	"fmt"
)

func main() {
	slicesInt := []int{1, 5, 2, 5, 3}

	fmt.Println("Check duplicate by sorting")
	if exercise.ContainsDuplicateBySorting(slicesInt) {
		fmt.Println("Duplicate in list")
	} else {
		fmt.Println("List value is unique")
	}

	fmt.Println("Check duplicate by map")
	if exercise.ContainsDuplicateByMap(slicesInt) {
		fmt.Println("Duplicate in list")
	} else {
		fmt.Println("List value is unique")
	}

	var inputString string
	fmt.Printf("Input string for validate: ")
	fmt.Scan(&inputString)
	fmt.Printf("Is valid input: %v\n", exercise.IsValid(inputString))

	fmt.Println("===================Exerise 2===================")
	fmt.Println("Using sorting")
	fmt.Printf("Is anagram: %v\n", exercise.IsAnagramUsingSorting("iloveyou", "youlovei"))

	fmt.Println("Using map")
	fmt.Printf("Is anagram: %v\n", exercise.IsAnagramUsingMap("anagram", "managra"))
}
