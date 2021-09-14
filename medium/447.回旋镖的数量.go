package medium

func numberOfBoomerangs(points [][]int) (res int) {
	//遍历每一个点和它所有点的距离，用hash表记录
	for _, p := range points {
		//必须放置在for循环内  每次大的循环重置一次map
		cnt := map[int]int{}
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cnt[dis]++
		}
		//如果重复距离相同的点 说明形成回旋镖 累计个数
		for _, m := range cnt {
			res += m * (m - 1)
		}
	}

	return
}
