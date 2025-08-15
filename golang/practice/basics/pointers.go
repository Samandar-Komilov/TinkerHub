package basics

func Swap(a *int, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp
}
