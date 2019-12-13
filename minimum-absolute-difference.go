package main

func minimumAbsoluteDifference(arr []int32) int32 {
	arrLength := int32(len(arr))
	var minDiff int32
	var tempDiff int32
	var j int32
	minDiff = absVal(arr[0] - arr[1])

	for i := int32(0); i < arrLength; i++ {
		j = i + 1
		for j < arrLength {
			tempDiff = absVal(arr[i] - arr[j])

			if tempDiff < minDiff {
				minDiff = tempDiff
			}
			j++
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
func main() {}
