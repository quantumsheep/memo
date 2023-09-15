# Memo

`memo` is a Go library for memoization.

## Installation

```bash
go get github.com/quantumsheep/memo
```

## Example

This is an example of how to use `memo` to memoize the Fibonacci function.

```go
func fib(n int) int {
    return memo.Memoize(func() int {
        if n < 2 {
            return n
            }

        return fib(n-1) + fib(n-2)
    }, n)
}
```

The first argument to `memo.Memoize` is a closure that returns a value. The next arguments are dependencies to be compared when determining whether to use the cached value or not.
