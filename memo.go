package memo

import (
	"fmt"
	"reflect"
	"runtime"
)

var cache = map[string]any{}

// Memoize is a function that takes a function and its dependencies and returns the result of the function.
// If the function has already been called with the same dependencies, the result is returned from the cache.
//
// # Example
//
//	func fib(n int) int {
//		return Memoize(func() int {
//			if n < 2 {
//				return n
//			}
//
//			return fib(n-1) + fib(n-2)
//		}, n)
//	}
func Memoize[Return any](f func() Return, dependencies ...any) Return {
	key := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	key += fmt.Sprint(dependencies...)

	if value, ok := cache[key]; ok {
		return value.(Return)
	}

	value := f()
	cache[key] = value
	return value
}
