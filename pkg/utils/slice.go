package utils

func ArrContains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
