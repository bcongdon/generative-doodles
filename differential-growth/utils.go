package main

func Lerp(start, stop, amt float64) float64 {
	return amt*(stop-start) + start
}
