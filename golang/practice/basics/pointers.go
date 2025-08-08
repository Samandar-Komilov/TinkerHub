package basics

func Swap(a *int, b *int) (*int, *int) {
	var tmp int = *a
	*a = *b
	*b = tmp

	return a, b
}
