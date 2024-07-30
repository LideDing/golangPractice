package leetcode

/*
给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/single-number-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func singleNumber(arr []int) []int {
	/*
		设两个出现一次的数为a,b，剩下所有的数都出现了偶数次
	*/
	res := make([]int, 2)
	eor := 0
	for _, v := range arr {
		eor ^= v
	}
	/*此时eor=a^b，我们可以知道a和b的二进制表示中至少有一位不同,因为a!=b，
	假设a和b在第8位上不一样（一个为1，一个为0），我们此时创建一个变量eor'，
	eor'去和arr中的所有第8位为1的数异或，那么eor'就是a或者b中的一个，假设为a，
	a再去和eor异或，那么就是b了，这样就找到了a和b
	*/
	eor_ := 0
	rightOne := eor & (^eor + 1) //提取出最右侧的1，eor与上自己的相反数，再与上自己，就可以提取出最右侧的1
	for _, v := range arr {
		if v&rightOne != 0 { //第rightOne位为1的数
			eor_ ^= v
		}
	}
	res[0] = eor_
	res[1] = eor_ ^ eor
	return res
}
