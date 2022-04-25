package sumutils

type Numeric interface {
	int64 | float64
}

func SumMaps[K comparable, V Numeric](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}
