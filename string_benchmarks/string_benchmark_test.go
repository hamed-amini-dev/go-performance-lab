package stringbenchmarks_test

import (
	"regexp"
	"runtime"
	"strings"
	"testing"
)

// Functions to be benchmarked
func removeParenthesesManual(input string) string {
	start := strings.Index(input, "(")
	if start != -1 {
		end := strings.Index(input, ")")
		if end != -1 && end > start {
			return input[:start] + input[end+1:]
		}
	}
	return input
}

func removeParenthesesRegex(input string) string {
	re := regexp.MustCompile(`\s*$begin:math:text$.*?$end:math:text$`)
	return re.ReplaceAllString(input, "")
}

// Benchmark functions
func BenchmarkRemoveParenthesesManual(b *testing.B) {
	runtime.GOMAXPROCS(1) // Ensure using 1 thread
	input := "USDT-TRC20(NILE)"
	for i := 0; i < b.N; i++ {
		removeParenthesesManual(input)
	}
}

func BenchmarkRemoveParenthesesRegex(b *testing.B) {
	runtime.GOMAXPROCS(1) // Ensure using 1 thread
	input := "USDT-TRC20(NILE)"
	for i := 0; i < b.N; i++ {
		removeParenthesesRegex(input)
	}
}
