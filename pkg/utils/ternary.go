package utils

func IIf[T any](condition bool, r1 T, r2 T) T {
	if condition {
		return r1
	} else {
		return r2
	}
}
