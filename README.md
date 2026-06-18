# calculator-go

A simple, lightweight command-line calculator written in Go. It supports basic arithmetic operations: addition, subtraction, multiplication, and division.

## Features

- Basic arithmetic: `+`, `-`, `*`, `/`
- Floating-point number support
- Clear error handling for invalid inputs and division by zero
- Minimal and fast – no external dependencies

## Installation

### Option 1: Using `go install`

```bash
go install github.com/MisaghMomeniB/calculator-go/src@latest
```

### Option 2: Build from source

Clone the repository and build the binary:

```bash
git clone https://github.com/MisaghMomeniB/calculator-go.git
cd calculator-go/src
go build -o calculator
```

## Usage

Run the calculator with two numbers and an operator:

```bash
calculator <num1> <operator> <num2>
```

### Supported Operators

| Operator | Operation     |
|----------|---------------|
| `+`      | Addition      |
| `-`      | Subtraction   |
| `*`      | Multiplication|
| `/`      | Division      |

### Examples

```bash
calculator 5 + 3
# Output: 5 + 3 = 8

calculator 10 - 4
# Output: 10 - 4 = 6

calculator 7 \* 6
# Output: 7 * 6 = 42

calculator 15 / 3
# Output: 15 / 3 = 5
```

> **Note:** On some shells, `*` may need to be escaped (`\*`) or quoted to prevent glob expansion.

### Error Handling

- **Invalid number**: Prints an error message and exits.
- **Division by zero**: Prints `Error: division by zero` and exits.
- **Unsupported operator**: Prints `Unsupported operator: <op>` and exits.
- **Incorrect arguments**: Shows usage instructions.

```
Usage: calculator <num1> <operator> <num2>
Operators: + - * /
```

## Project Structure

```
calculator-go/
├── src/
│   └── main.go      # Main application entry point
├── LICENSE
├── .gitignore
└── README.md
```

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file in the repository.

---
