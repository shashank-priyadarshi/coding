package main

func firstAndLastPosition(arr []int, target int) (int, int) {
	length := len(arr) - 1
	start, end := -1, -1

	for k, v := range arr {
		if start == -1 && v == target {
			start = k
		}

		if end == -1 && arr[length] == target {
			end = length
		}

		length--

		if start != -1 && end != -1 {
			return start, end
		}

	}

	return start, end
}
