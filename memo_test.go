package memo

import "testing"

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fibMemoize(n int) int {
	return Memoize(func() int {
		if n < 2 {
			return n
		}

		return fibMemoize(n-1) + fibMemoize(n-2)
	}, n)
}

var fib10Result = 6765

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := fib(20)
		if result != fib10Result {
			b.Fatalf("Expected %d, got %d", fib10Result, result)
		}
	}
}

func BenchmarkFibMemoize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := fibMemoize(20)
		if result != fib10Result {
			b.Fatalf("Expected %d, got %d", fib10Result, result)
		}
	}
}
