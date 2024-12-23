package util

func Abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
