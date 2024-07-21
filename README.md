# Go LeetCode

Welcome to the `go-leetcode` repository! This repository contains solutions to various LeetCode problems written in Go. Each problem is organized into a directory following a consistent naming convention for easy navigation and management.

## Directory Structure

```bash
go-leetcode/
├── algorithms/
│   └── 1-two-sum/
│       ├── main.go
│       └── main_test.go
├── scripts/
│   └── generate.sh
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
make generate PROBLEM_NUMBER=1 PROBLEM_NAME=two-sum
```

This will create a new directory `algorithms/1-two-sum` with `main.go` and `main_test.go` files pre-populated with boilerplate code.

### Running a Problem

To run a specific problem's main file, use the `run` target in the `Makefile`:

```sh
make run PROBLEM_DIR=algorithms/1-two-sum
```

### Testing a Problem

To run tests for a specific problem, use the `test` target in the `Makefile`:

```sh
make test PROBLEM_DIR=algorithms/1-two-sum
```

### Adding More Test Cases

You can add more test cases in the `main_test.go` file under the `testCases` array. Each test case is a struct containing the input and the expected output.

### Scripts

#### Generate Script

The `generate.sh` script automates the creation of new problem directories.

Usage:

```sh
./scripts/generate.sh <problem_number> <problem_name>
```

Example:

```sh
./scripts/generate.sh 1 two-sum
```

This command will create the directory `algorithms/1-two-sum` with the necessary boilerplate code.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/patrick204nqh/go-leetcode/tree/main?tab=MIT-1-ov-file) file for details.

---

Happy coding! 🔨🚀
