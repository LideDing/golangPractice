package sort

func SelectSort(arr []int) []int { //选择排序
	for i := 0; i < len(arr); i++ { //从第一个元素开始
		min := i                            //最小元素的下标
		for j := i + 1; j < len(arr); j++ { //从第二个元素开始
			if arr[j] < arr[min] { //如果后面的元素比最小元素小
				min = j //更新最小元素的下标
			}
		}
		arr[i], arr[min] = arr[min], arr[i] //将最小元素放到最前面
	}
	return arr
}
