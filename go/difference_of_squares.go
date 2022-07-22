package diffsquares

func SquareOfSum(n int) int {
	x := n * (n + 1) / 2
	return x * x
}

func SumOfSquares(n int) int {
	return n * (2*n + 1) * (n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
