package math

func Swap(x *int, y *int) {
	*x = *x ^ *y
	*y = *x ^ *x
	*x = *x ^ *y
}
