package basics

import (
	"cmp"
	"fmt"
)

func Main_generics() {
	// 1
	a, b := 1, 3
	a, b = SwapG(a, b)
	fmt.Println(a, b)
	SwapGPtr(&a, &b)
	fmt.Println(a, b)

	// 3
	mx := Max(a, b)
	fmt.Println(mx)
}

func SwapG[T any](a T, b T) (T, T) {
	a, b = b, a

	return a, b
}

func SwapGPtr[T any](a *T, b *T) {
	*a, *b = *b, *a
}

func Max[T cmp.Ordered](a T, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}

}
