## Usage

To run the benchmark tests, clone the repository and use the Go testing framework:

```sh
git clone https://github.com/yourusername/go-performance-tests.git
cd go-performance-tests
go test -bench=.

BenchmarkRemoveParenthesesManual-11 134282155 8.776 ns/op
testing: BenchmarkRemoveParenthesesManual-11 left GOMAXPROCS set to 1
BenchmarkRemoveParenthesesRegex-11 388644 2947 ns/op
testing: BenchmarkRemoveParenthesesRegex-11 left GOMAXPROCS set to 1
```
