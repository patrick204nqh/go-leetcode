# Go LeetCode

[![Go Version](https://img.shields.io/github/go-mod/go-version/patrick204nqh/go-leetcode)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/patrick204nqh/go-leetcode)](https://goreportcard.com/report/github.com/patrick204nqh/go-leetcode)
[![License](https://img.shields.io/github/license/patrick204nqh/go-leetcode)](https://github.com/patrick204nqh/go-leetcode/blob/main/LICENSE)
[![Issues](https://img.shields.io/github/issues/patrick204nqh/go-leetcode)](https://github.com/patrick204nqh/go-leetcode/issues)

Welcome to the `go-leetcode` repository! This repository contains solutions to various LeetCode problems written in Go. Each problem is organized into a directory following a consistent naming convention for easy navigation and management.

## Directory Structure

```bash
go-leetcode/
├── algorithms/
│   ├── 1-two-sum/
│   │   ├── main.go
│   │   └── main_test.go
│   └── ...
├── scripts/
│   └── generate.sh
├── utils/
│   ├── perf/
│   │   └── analyzer.go
│   └── ...
├── Makefile
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.18 or higher
- Git

### Clone the Repository

```sh
git clone https://github.com/patrick204nqh/go-leetcode.git
cd go-leetcode
```

### Generate a New Problem Directory

To generate a new problem directory, use the `generate` target in the `Makefile`:

```sh
make generate 1 two-sum
```

This will create a new directory `algorithms/1-two-sum` with `main.go` and `main_test.go` files pre-populated with boilerplate code.

### Running a Problem

To run a specific problem's main file, use the `run` target in the `Makefile`:

```sh
make run algorithms/1-two-sum
```

### Testing a Problem

To run tests for a specific problem, use the `test` target in the `Makefile`:

```sh
make test algorithms/1-two-sum
```

### Performance Analysis

To use the performance analyzer, import the `perf` package and wrap your code block:

```go
import "github.com/patrick204nqh/go-leetcode/utils/perf"

perf.Analyze("Description", func() {
    // Your code here
})
```

#### Example Usage

```go
package main

import "path/to/utils/perf"

func exampleFunction() {
    sum := 0
    for i := 0; i < 1000000; i++ {
        sum += i
    }
}

func main() {
    perf.Analyze("Example Function", exampleFunction)
}
```

#### Sample Output

```
+=================== Performance Analysis ===================+
| Metric             | Value                                 |
+============================================================+
| Description        | Example Function                      |
+------------------------------------------------------------+
| Total Time (s)     | 0.012345                              |
| Memory Before (KB) | 2048                                  |
| Memory After (KB)  | 3072                                  |
| Memory Used (KB)   | 1024                                  |
| Memory Change (%)  | 50.00                                 |
| Goroutines Before  | 2                                     |
| Goroutines After   | 2                                     |
+============================================================+
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/patrick204nqh/go-leetcode/tree/main?tab=MIT-1-ov-file) file for details.

---

Happy coding! 🔨🚀
