package main

func minimumAbsoluteDifference(arr []int32) int32 {
	arrLength := int32(len(arr))
	var minDiff int32
	var tempDiff int32

	minDiff = absVal(arr[0] - arr[1])

	arr = mergeSort(arr)

	for i := int32(0); i < arrLength-1; i++ {
		tempDiff = absVal(arr[i] - arr[i+1])

		if tempDiff < minDiff {
			minDiff = tempDiff
		}
	}

	return minDiff
}

func absVal(a int32) int32 {
	if a < 0 {
		return -a
	}

	return a
}

func mergeSort(arr []int32) []int32 {
	arrLength := int32(len(arr))

	if arrLength == 1 {
		return arr
	}

	arrMiddle := arrLength / 2
	var left = make([]int32, arrMiddle)
	var right = make([]int32, arrLength-arrMiddle)

	for i := int32(0); i < arrLength; i++ {
		if i < arrMiddle {
			left[i] = arr[i]
		} else {
			right[i-arrMiddle] = arr[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int32) []int32 {
	result := make([]int32, len(left)+len(right))

	i := int32(0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func main() {}
