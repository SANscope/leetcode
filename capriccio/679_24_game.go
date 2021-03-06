package main

import (
	"fmt"
)

func judgePoint24(cards []int) bool {
	// write code here
	subResult := make([]float64, 4)
	for i := 0; i < len(cards); i++ {
		subResult[i] = float64(cards[i])
	}
	return Game24Backtracking(&subResult, "")
}

func Game24Backtracking(nums *[]float64, str string) bool {
	// Termination condition.
	if len(*nums) == 1 {
		if (*nums)[0] > 24-1e-9 && (*nums)[0] < 24+1e-9 {
			return true
		} else {
			return false
		}
	}

	// 游戏的第一步是挑出两个数，算出一个新数替代这两个数。
	// 然后，在三个数中玩 24 点，再挑出两个数，算出一个数替代它们。
	// 然后，在两个数中玩 24 点。
	// Each time to pick two cards.
	for pickedOne := 0; pickedOne < len(*nums)-1; pickedOne++ {
		for pickedTwo := pickedOne + 1; pickedTwo < len(*nums); pickedTwo++ {
			// Get all combinations.
			left := (*nums)[pickedOne]
			right := (*nums)[pickedTwo]

			// Need backtracking.
			*nums = append((*nums)[:pickedTwo], (*nums)[pickedTwo+1:]...)
			*nums = append((*nums)[:pickedOne], (*nums)[pickedOne+1:]...)

			// +.
			*nums = append(*nums, left+right)
			if Game24Backtracking(nums, str+fmt.Sprintf("(%f + %f)", left, right)) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]

			// -.
			*nums = append(*nums, left-right)
			if Game24Backtracking(nums, str+fmt.Sprintf("(%f - %f)", left, right)) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]
			// -.
			*nums = append(*nums, right-left)
			if Game24Backtracking(nums, str+fmt.Sprintf("(%f - %f)", right, left)) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]

			// *.
			*nums = append(*nums, left*right)
			if Game24Backtracking(nums, str+fmt.Sprintf("(%f * %f)", left, right)) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]

			// /.
			*nums = append(*nums, left/right)
			if Game24Backtracking(nums, str+fmt.Sprintf("(%f / %f)", left, right)) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]
			// /.
			*nums = append(*nums, right/left)
			if Game24Backtracking(nums, str+fmt.Sprintf("(%f / %f)", right, left)) {
				return true
			}
			*nums = (*nums)[:len(*nums)-1]

			// 复原到原来的位置。
			// 需要申请新的空间。
			temp := make([]float64, len((*nums)[pickedOne:]))
			copy(temp, (*nums)[pickedOne:])
			*nums = append((*nums)[:pickedOne], left)
			*nums = append(*nums, temp...)

			temp = make([]float64, len((*nums)[pickedTwo:]))
			copy(temp, (*nums)[pickedTwo:])
			*nums = append((*nums)[:pickedTwo], right)
			*nums = append(*nums, temp...)
		}
	}
	return false
}
