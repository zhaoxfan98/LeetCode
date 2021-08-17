//package main
package easy

func findMaxConsecutiveOnes(nums []int) int {
	count,max := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			count = 0
		}
		if max < count {
			max = count
		}

	}
	return max
}

//func main() {
//	nums := []int{1,1,0,1,1,1}
//	fmt.Println(findMaxConsecutiveOnes(nums))
//}
