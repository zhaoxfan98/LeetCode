package easy

//定义两个指针fast和slow分别为快指针和慢指针，快指针表示遍历数组到达的下标位置，慢指针表示下一个不同元素要填入的下标位置
//fast依次遍历从1到n-1的每个位置，对于每个位置，如果nums[fast] != nums[fast-1],
//说明nums[fast]和之前的元素都不同，因此将nums[fast]的值复制到nums[slow]，然后将slow的值加1，即指向下一个位置
//遍历结束之后，从nums[0]到nums[slow-1]的每个元素都不相同且包含原数组中的每个不同的元素，因此新的长度即为slow

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
