package tools

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func Reverse[T any](arr []T) chan T {
	ret := make(chan T)
	go func() {
		for i := range arr {
			ret <- arr[len(arr)-1-i]
		}
		close(ret)
	}()
	return ret
}
