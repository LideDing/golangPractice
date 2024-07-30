package sort

//冒泡排序
/*
冒泡排序是一种简单的排序算法。它重复地遍历要排序的列表，比较每对相邻的元素，并在必要时交换它们的位置。
如果要排序的列表有 n 个元素，则需要遍历 n-1 次。它的时间复杂度为 O(n^2)。
*/
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ { //从第一个元素开始
		for j := 0; j < len(arr)-1-i; j++ { //从第一个元素开始，到未排序的最后一个元素
			if arr[j] > arr[j+1] { //如果前面的元素比后面的元素大
				arr[j], arr[j+1] = arr[j+1], arr[j] //交换两个元素
			}
		}
	}
	return arr
}
