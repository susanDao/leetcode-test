package main

import (
	"fmt"
)

func sortAndConvert(nums []int) string {
	for i := len(nums) - 1; i > 0; i-- {
		for k := 0; k < i; k++ {
			if nums[k] > nums[k+1] {
				nums[k], nums[k+1] = nums[k+1], nums[k]
			}
		}
	}
	s := fmt.Sprintf("%d-%d-%d", nums[0], nums[1], nums[2])
	return s
}
func threeSum(nums []int) [][]int {
	strExist := make(map[string]int, 0)
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			result = append(result, nums)
			return result
		}
	}
	l := len(nums)
	for first := 0; first < l-2; first++ {
		for second := first + 1; second < l-1; second++ {
			for thrid := second + 1; thrid < l; thrid++ {
				if nums[first]+nums[second]+nums[thrid] == 0 {
					tmp := []int{nums[first], nums[second], nums[thrid]}
					s := sortAndConvert(tmp)
					if _, ok := strExist[s]; !ok {
						strExist[s] = 1
						result = append(result, tmp)
					}
				}
			}
		}
	}
	return result

}

func main() {
	r := []int{-1, 0, 1, 2, -1, -4}
	//r := []int{8, -15, -2, -13, 8, 5, 6, -3, -9, 3, 6, -6, 8, 14, -9, -8, -9, -6, -14, 5, -7, 3, -10, -14, -12, -11, 12, -15, -1, 12, 8, -8, 12, 13, -13, -3, -5, 0, 10, 2, -11, -7, 3, 4, -8, 9, 3, -10, 11, 5, 10, 11, -7, 7, 12, -12, 3, 1, 11, 9, -9, -4, 9, -12, -6, 11, -7, 4, -4, -12, 13, -8, -12, 2, 3, -13, -12, -8, 14, 14, 12, 9, 10, 12, -6, -1, 8, 4, 8, 4, -1, 14, -15, -7, 9, -14, 11, 9, 5, 14}
	s := threeSum(r)
	fmt.Println(s)
}
