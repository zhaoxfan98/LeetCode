package hard

import "strings"

/*
对于每一行，首先确定最多可以放置多少单词，这样可以得到该行的空格个数，从而确定该行单词之间的空格个数

题目中填充空格的细节  分以下三种情况讨论
- 当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格；
- 当前行不是最后一行，且只有一个单词：该单词左对齐，在行末填充空格；

- 当前行不是最后一行，且不只一个单词：设当前行单词数为numWords，空格数为numSpaces,我们需要将空格均匀分配在单词之间，
则单词之间应至少有
	avgSpaces = [numSpaces/(numWords-1)]  向下取值
个空格，对于多出来的
	extraSpaces = numSpaces mod (numWords-1)
个空格，应填在前extraSpaces个单词之间。因此，前extraSpaces个单词之间填充avgSpaces+1个单词，其余单词之间填充avgSpaces个空格
*/

func fullJustify(words []string, maxWidth int) (ans []string) {
	right, n := 0, len(words)
	for {
		left := right //当前行的第一个单词在words的位置
		sumLen := 0   //统计这一行单词长度之和
		//循环确定当前行可以放多少单词，注意单词之间应至少有一个空格
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		//当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格
		if right == n {
			s := strings.Join(words[left:], " ")
			ans = append(ans, s+blank(maxWidth-len(s)))
			return
		}

		numWords := right - left
		numSpaces := maxWidth - sumLen

		//当前行只有一个单词：该单词左对齐，在行末填充剩余空格
		if numWords == 1 {
			ans = append(ans, words[left]+blank(numSpaces))
			continue
		}

		//当前行不只一个单词
		avgSpaces := numSpaces / (numWords - 1)
		extraSpaces := numSpaces % (numWords - 1)
		s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1)) //拼接额外加一个空格的单词
		s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))  //拼接其余单词
		ans = append(ans, s1+blank(avgSpaces)+s2)
	}
}

//blank返回长度为n的由空格组成的字符串
func blank(n int) string {
	return strings.Repeat(" ", n)
}
