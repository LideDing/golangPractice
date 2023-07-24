package src
import(
	"sort"
	// "fmt"
)
/*
三数之和
给一个数组，问数组中是否存在三个元素相加为0，如果存在，返回这三个数字
*/

func ThreeSum(nums []int)[][]int{
	res:=[][]int{}
	n:=len(nums)
	if n<3{
		return res
	}
	sort.Ints(nums)
	
	for first:=0;first<n;first++{
		if first>0 && nums[first]==nums[first-1]{
			continue // 遇到相同的跳过，避免重复计算
		}
		target:=-nums[first] //second+third的目标值
		third:=n-1
		for second:=first+1;second<n;second++{
			if second>first+1 && nums[second]==nums[second-1]{
				continue // 遇到相同的跳过，避免重复计算
			}
			for third>second && nums[second]+nums[third]>target{
				third--
			}
			if second==third{
				break
			}
			if nums[second]+nums[third]==target{
				res = append(res, []int{nums[first],nums[second],nums[third]})
			}
		}
	}
	return res
}