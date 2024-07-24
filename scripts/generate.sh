#!/bin/bash

if [ $# -lt 2 ]; then
    echo "Usage: $0 <problem_number> <problem_name>"
    exit 1
fi

PROBLEM_NUMBER=$1
PROBLEM_NAME=$2
PROBLEM_DIR="algorithms/$PROBLEM_NUMBER-$PROBLEM_NAME"

if [ -d "$PROBLEM_DIR" ]; then
    echo "Directory $PROBLEM_DIR already exists."
    exit 1
fi

mkdir -p $PROBLEM_DIR

# Convert problem_name to camelCase for function names
to_camel_case() {
    IFS='-' read -r -a array <<< "$1"
    local output="${array[0]}"
    for element in "${array[@]:1}"
    do
        output+=$(tr '[:lower:]' '[:upper:]' <<< ${element:0:1})${element:1}
    done
    echo "$output"
}

CAMEL_CASE_NAME=$(to_camel_case $PROBLEM_NAME)
CAPITALIZED_CAMEL_CASE_NAME="$(tr '[:lower:]' '[:upper:]' <<< ${CAMEL_CASE_NAME:0:1})${CAMEL_CASE_NAME:1}"

# Create main.go
cat <<EOL > $PROBLEM_DIR/main.go
package main

import (
    "fmt"
    "github.com/fatih/color"
)

func $CAMEL_CASE_NAME(nums []int, target int) []int {
    // Your implementation here
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := $CAMEL_CASE_NAME(nums, target)
    color.Green("Result: %v", result) // Your expected output here
}
EOL

# Create main_test.go
cat <<EOL > $PROBLEM_DIR/main_test.go
package main

import (
    "github.com/fatih/color"
    "reflect"
    "testing"
)

func test$CAPITALIZED_CAMEL_CASE_NAME(t *testing.T) {
    testCases := []struct {
        nums     []int
        target   int
        expected []int
    }{
        // Add your test cases here
        {nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
    }

    for _, tc := range testCases {
        result := $CAMEL_CASE_NAME(tc.nums, tc.target)
        if !reflect.DeepEqual(result, tc.expected) {
            color.Red("FAIL: $CAMEL_CASE_NAME(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
            t.Errorf("FAIL: $CAMEL_CASE_NAME(%v, %d) = %v; expected %v", tc.nums, tc.target, result, tc.expected)
        } else {
            color.Green("PASS: $CAMEL_CASE_NAME(%v, %d) = %v", tc.nums, tc.target, result)
        }
    }
}
EOL

echo "Problem directory $PROBLEM_DIR created."
