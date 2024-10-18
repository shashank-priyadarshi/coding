package main

func binary_search(arr []int, target int) int {
	length := len(arr)
	high, low := 0, length
	mid := (high + low) / 2

	for i := 0; i < length; i++ {
		if target == arr[i] {
			return i
		}

		if target < arr[mid] {
			high = mid
			mid = (high + low) / 2
			continue
		}

		low = mid
		mid = (high + low) / 2
	}

	return -1
}

func linear_search(arr []int, target int) int {
	for key, val := range arr {
		if val == target {
			return key
		}
	}
	return -1
}
