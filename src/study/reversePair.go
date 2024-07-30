package study

//在一个数组中，左边的数如果比右边的数大，则这两个数构成一个逆序对，打印所有的逆序对，找出逆序对的数量

func reverseMergeSort(arr []int) (int, []int) { //
	if len(arr) <= 1 {
		return 0, arr
	}

	mid := len(arr) / 2
	leftNum, leftArr := reverseMergeSort(arr[:mid])
	rightNum, rightArr := reverseMergeSort(arr[mid:])
	mergedNum, mergedArr := reverseMerge(leftArr, rightArr)

	return leftNum + rightNum + mergedNum, mergedArr
}

func reverseMerge(leftArr, rightArr []int) (int, []int) {
	merged := make([]int, 0, len(leftArr)+len(rightArr))
	i, j, count := 0, 0, 0

	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			merged = append(merged, leftArr[i])
			i++
		} else {
			merged = append(merged, rightArr[j])
			j++
			count += len(leftArr) - i
		}
	}

	merged = append(merged, leftArr[i:]...)
	merged = append(merged, rightArr[j:]...)

	return count, merged
}

func reversePairs(arr []int) int {
	count, _ := reverseMergeSort(arr)
	return count
}
