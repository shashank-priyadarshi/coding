package main

func peak_n(arr []int) int {
	length := len(arr)
	for k, v := range arr {
		if k >= length-1 {
			return -1
		}

		if k == 0 {
			continue
		}

		if v > arr[k-1] && v > arr[k+1] {
			return k
		}
	}

	return -1
}
