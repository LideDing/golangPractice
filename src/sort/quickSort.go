package sort

func QuickSort(arr []int, low, high int) {
	if low >= high { // 递归终止条件
		return
	}
	pivot := partition(arr, low, high) // 分区
	QuickSort(arr, low, pivot-1)
	QuickSort(arr, pivot+1, high)
}

func partition(arr []int, low, high int) int { // 分区
	pivot := arr[high]            // 选取最后一个元素作为基准
	i := low                      // i指向小于基准的最后一个元素
	for j := low; j < high; j++ { // j指向当前遍历的元素
		if arr[j] < pivot { // 如果当前元素小于基准
			arr[i], arr[j] = arr[j], arr[i] // 交换当前元素和小于基准的最后一个元素
			i++                             // 小于基准的最后一个元素的下标+1
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // 将基准元素放到小于基准的最后一个元素的后面
	return i                              // 返回基准元素的下标
}
