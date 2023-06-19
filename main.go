package main

import (
	"fmt"
	"sort"
)

func main() {
	slicesInt := []int{1, 5, 2, 5, 3}

	fmt.Println("Check duplicate by sorting")
	if containsDuplicateBySorting(slicesInt) {
		fmt.Println("Duplicate in list")
	} else {
		fmt.Println("List value is unique")
	}

	fmt.Println("Check duplicate by map")
	if containsDuplicateByMap(slicesInt) {
		fmt.Println("Duplicate in list")
	} else {
		fmt.Println("List value is unique")
	}

	var inputString string
	fmt.Printf("Input string for validate: ")
	fmt.Scan(&inputString)
	fmt.Printf("Is valid input: %v", isValid(inputString))
}

// Chủ trương: add các ký tự mở vào slices
// nếu gặp ký tự đóng, so sánh với ký tự trước nó
// nếu == nhau, clear đi kể từ ký tự mở tương ứng với ký tự đó
//		tiếp tục append nếu là mở, nếu đóng, tiếp tục so sánh và clear đi data từ chỗ mở
// nếu != -> return false

// Case 0: have character not in list -> không
// input: 0
// stack: 1 -> return false

// Case 1: only have 1 open character -> chưa kịp bắt đầu đã kết thúc
// input: )
//stack: ( -> return false

// Case 2: only have 1 close character -> chúng ta đã sai từ đầu
// input: (
// stack: go in if ->
//					len(stack) == 0 -> return false

// Case 3: have close and open character but not match -> kết thúc không trọn vẹn
// input: (]
// stack: append: (. continue: find ] -> but stack[0] != ] -> return false

// Case 4: người thứ ba
// input: (})
// stack: append (
// val = {
// stack[0] = ( != }
// return false

// Case 5: tình yêu trọn vẹn
// input: ({})
// stack: append (
// stack: append {
// stack[({] -> val = {
// stack = stack[:(2-1)] -- now -> stack[(]
// val = ) --- stack = stack(:0) => empty slices
// return true

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if val, ok := mapping[(char)]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			//stack = []rune{}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func containsDuplicateByMap(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func containsDuplicateBySorting(nums []int) bool {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
