package sort

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) >> 1
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	res := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)
	return res
}
