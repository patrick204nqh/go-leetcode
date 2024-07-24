#!/bin/bash

# ANSI color codes for prettifying the output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to convert problem name to camelCase
to_camel_case() {
    IFS='-' read -r -a array <<< "$1"
    local output="${array[0]}"
    for element in "${array[@]:1}"; do
        output+=$(tr '[:lower:]' '[:upper:]' <<< "${element:0:1}")${element:1}
    done
    echo "$output"
}

# Check if the correct number of arguments is provided
if [ $# -lt 2 ]; then
    echo -e "${RED}Usage: $0 <problem_number> <problem_name>${NC}"
    exit 1
fi

# Read input arguments
PROBLEM_NUMBER=$1
PROBLEM_NAME=$2
PROBLEM_DIR="algorithms/${PROBLEM_NUMBER}-${PROBLEM_NAME// /-}"

# Check if the directory already exists
if [ -d "$PROBLEM_DIR" ]; then
    echo -e "${YELLOW}Directory $PROBLEM_DIR already exists.${NC}"
    exit 1
fi

# Create the directory
mkdir -p "$PROBLEM_DIR"

# Convert problem name to camelCase
CAMEL_CASE_NAME=$(to_camel_case "$PROBLEM_NAME")
CAPITALIZED_CAMEL_CASE_NAME="$(tr '[:lower:]' '[:upper:]' <<< "${CAMEL_CASE_NAME:0:1}")${CAMEL_CASE_NAME:1}"

# Prompt user for function parameters and expected output
echo -e "${GREEN}Enter function parameters (comma separated, e.g., 'nums []int, target int'):${NC}"
read -r FUNCTION_PARAMETERS
echo -e "${GREEN}Enter expected output type (e.g., '[]int'):${NC}"
read -r EXPECTED_OUTPUT
echo -e "${GREEN}Enter a test case input (comma separated, e.g., '[]int{2, 7, 11, 15}, 9'):${NC}"
read -r TEST_CASE_INPUT
echo -e "${GREEN}Enter expected output for the test case (e.g., '[]int{0, 1}'):${NC}"
read -r TEST_CASE_EXPECTED_OUTPUT

# Create main.go
cat <<EOL > "$PROBLEM_DIR/main.go"
package main

import (
    "fmt"
    "github.com/fatih/color"
)

func $CAMEL_CASE_NAME($FUNCTION_PARAMETERS) $EXPECTED_OUTPUT {
    // Your implementation here
    return nil
}

func main() {
    result := $CAMEL_CASE_NAME($TEST_CASE_INPUT)
    color.Green("Result: %v", result) // Your expected output here
}
EOL

# Create main_test.go
cat <<EOL > "$PROBLEM_DIR/main_test.go"
package main

import (
    "github.com/fatih/color"
    "reflect"
    "testing"
)

func Test${CAPITALIZED_CAMEL_CASE_NAME}(t *testing.T) {
    testCases := []struct {
        input    []interface{}
        expected $EXPECTED_OUTPUT
    }{
        // Add your test cases here
        {input: []interface{}{$TEST_CASE_INPUT}, expected: $TEST_CASE_EXPECTED_OUTPUT},
    }

    for _, tc := range testCases {
        result := $CAMEL_CASE_NAME(tc.input...)
        if !reflect.DeepEqual(result, tc.expected) {
            color.Red("FAIL: $CAMEL_CASE_NAME(%v) = %v; expected %v", tc.input, result, tc.expected)
            t.Errorf("FAIL: $CAMEL_CASE_NAME(%v) = %v; expected %v", tc.input, result, tc.expected)
        } else {
            color.Green("PASS: $CAMEL_CASE_NAME(%v) = %v", tc.input, result)
        }
    }
}
EOL

echo -e "${GREEN}Problem directory $PROBLEM_DIR created.${NC}"
