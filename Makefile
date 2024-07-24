# Default target
.PHONY: all
all: run

# Run the Go program
.PHONY: run
run:
	@echo "Running main.go in $(filter-out $@,$(MAKECMDGOALS))..."
	@cd $(filter-out $@,$(MAKECMDGOALS)) && go run main.go

# Test the Go program
.PHONY: test
test:
	@echo "Running tests in $(filter-out $@,$(MAKECMDGOALS))..."
	@cd $(filter-out $@,$(MAKECMDGOALS)) && go test -v

# Generate a new problem directory with positional parameters
.PHONY: generate
generate:
	@echo "Generating new problem directory..."
	@./scripts/generate.sh $(filter-out $@,$(MAKECMDGOALS))

# Display usage information
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make [target] [PROBLEM_NUMBER=<number>] [PROBLEM_NAME=<name>]"
	@echo "  make generate <problem_number> <problem_name>"
	@echo ""
	@echo "Targets:"
	@echo "  all       - Run the default target (run)"
	@echo "  run       - Run the Go program in the specified problem directory"
	@echo "  test      - Run tests in the specified problem directory"
	@echo "  generate  - Generate a new problem directory with custom arguments"

# Ignore any make goals that are numbers or problem names
%:
	@: