package src

//插入排序
/*
标准的插入排序算法将数组分为已排序和未排序两部分，然后依次将未排序部分的每个元素插入到已排序部分中的正确位置。
*/
func InsertSort(arr []int) []int {
    n := len(arr)
    if n < 2{
        return arr
    }
    for i := 1; i < n; i++{ //从1开始，因为0到0已经有序
        for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j--{ //j和j+1的位置比较，j>=0表示不越界
            arr[j], arr[j+1] = arr[j+1], arr[j]
        }
    }
    return arr

}



/*
    for i := 1; i < len(arr); i++ {  //从第二个元素开始
        key := arr[i] //待排序的元素
        j := i - 1   //已经排好序的元素

        for j >= 0 && arr[j] > key { //从后往前比较，如果已经排好序的元素大于待排序的元素，就将已经排好序的元素往后移动
            arr[j+1] = arr[j] //将已经排好序的元素往后移动
            j = j - 1 //继续往前比较
        }
        arr[j+1] = key //将待排序的元素插入到合适的位置
    }

    return arr
*/