package easy

func twoSum(numbers []int, target int) []int {
	var res = []int{0, 0}
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			res[0], res[1] = i+1, j+1
			break
		} else if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		}
	}
	return res
}

//func main() {
//	var num = []int{2, 7, 11, 15}
//	fmt.Println(twoSum(num, 9))
//}
