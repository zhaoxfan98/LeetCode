package easy

//暴力枚举
func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		// 每一个位于x之前的元素都已经和x匹配过，因此不需要再进行匹配
		// 而每一个元素不能被使用两次，所以我们只需要在 x 后面的元素中寻找 target - x
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//哈希表
//对于每一个x，首先查询哈希表中是否存在target-x，然后将x插入到哈希表中，
//即可保证不会让x和自己匹配
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
