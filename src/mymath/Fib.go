package mymath

// Get Fib number
func GetFib(n int) int {
	var fib int = 1
	if n > 2 {
		fib = GetFib(n - 1) + GetFib(n - 2)
	}
	return fib
}