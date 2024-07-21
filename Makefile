# Variables
PROBLEM_DIR ?= algorithms/1-two-sum

# Default target
.PHONY: all
all: run

# Run the Go program
.PHONY: run
run:
	@echo "Running main.go in $(PROBLEM_DIR)..."
	@cd $(PROBLEM_DIR) && go run main.go

# Test the Go program
.PHONY: test
test:
	@echo "Running tests in $(PROBLEM_DIR)..."
	@cd $(PROBLEM_DIR) && go test -v

# Generate a new problem directory
.PHONY: generate
generate:
	@echo "Generating new problem directory..."
	@./scripts/generate.sh $(PROBLEM_NUMBER) $(PROBLEM_NAME)
