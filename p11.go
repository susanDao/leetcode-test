package main

import "fmt"

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	if len(height) == 2 {
		return min(height[0], height[1])
	}

	lHeight := len(height)
	h, t := height[0], height[lHeight-1]
	area := (len(height) - 1) * min(h, t)
	areaTmp := 0

	i, j := 0, lHeight-1

	for i < j {
		areaTmp = (j - i) * min(height[i], height[j])
		if areaTmp > area {
			area = areaTmp
		}

		if height[i] < height[j] {
			if height[i] < height[i+1] {
				i++
			} else {
				i = i + 2
			}
		} else {
			if height[j-1] > height[j] {
				j--
			} else {
				j = j - 2
			}
		}
	}

	return area
}

func main() {
	t1 := []int{76, 155, 15, 188, 180, 154, 84, 34, 187, 142, 22, 5, 27, 183, 111, 128, 50, 58, 2, 112, 179, 2, 100, 111, 115, 76, 134, 120, 118, 103, 31, 146, 58, 198, 134, 38, 104, 170, 25, 92, 112, 199, 49, 140, 135, 160, 20, 185, 171, 23, 98, 150, 177, 198, 61, 92, 26, 147, 164, 144, 51, 196, 42, 109, 194, 177, 100, 99, 99, 125, 143, 12, 76, 192, 152, 11, 152, 124, 197, 123, 147, 95, 73, 124, 45, 86, 168, 24, 34, 133, 120, 85, 81, 163, 146, 75, 92, 198, 126, 191}
	fmt.Println(maxArea(t1))
}
