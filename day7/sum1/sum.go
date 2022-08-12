package sum1

func Nums(vs ...int) int {
	return nums(vs)
}

func nums(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return nums(vs[1:]) + vs[0]
}
