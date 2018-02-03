package sum

func Ints(vs ...int) int {
	return ints(vs...)
}

func ints(vs ...int) int {
	if len(vs) == 0 {
		return 0
	}
	return vs[0] + ints(vs[1:]...)
}
