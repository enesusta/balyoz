package math

func Swap(x *int, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}
