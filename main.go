package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type xfix uint

const (
	prefix xfix = iota
	infix
	postfix
)

type cFunc struct {
	opType xfix
	fn     func(a, b float64) string
}

func myFormat(x float64) string {
	return strconv.FormatFloat(x, 'g', -1, 64)
}

var F = map[string]cFunc{

	"sqrt": {
		prefix,
		func(_, b float64) string {
			if b >= 0 {
				return "±" + myFormat(math.Sqrt(b))
			} else {
				return "You can't take the square root of a negative number."
			}
		},
	},

	"sin": {
		prefix,
		func(_, b float64) string {
			return myFormat(math.Sin(b))
		},
	},

	"cos": {
		prefix,
		func(_, b float64) string {
			return myFormat(math.Cos(b))
		},
	},

	"tan": {
		prefix,
		func(_, b float64) string {
			if (math.Mod(b, math.Pi/2) == 0) && (math.Mod(b, math.Pi) != 0) {
				return "Undefined."
			} else {
				return myFormat(math.Tan(b))
			}
		},
	},

	"+": {
		infix,
		func(a, b float64) string {
			return myFormat(a + b)
		},
	},

	"-": {
		infix,
		func(a, b float64) string {
			return myFormat(a - b)
		},
	},

	"*": {
		infix,
		func(a, b float64) string {
			return myFormat(a * b)
		},
	},

	"/": {
		infix,
		func(a, b float64) string {
			if b == 0 {
				return "You can't divide by zero."
			} else {
				return myFormat(a / b)
			}
		},
	},

	"squared": {
		postfix,
		func(a, _ float64) string {
			return myFormat(a * a)
		},
	},
}

func notF(r rune) bool {
	return !strings.ContainsRune("-.0123456789", r)
}

func outWait(s string) {
	fmt.Println(s)
	fmt.Scanln()
}

func main() {

	input := os.Args[1]

	fmt.Println(input)

	// Find the operator.

	opStart := strings.IndexFunc(input, notF)
	opEnd := strings.LastIndexFunc(input, notF) + 1

	// Validate the operator.

	if opStart == -1 {
		outWait("Your input contains no operator.")
		return
	}
	op := strings.TrimSpace(input[opStart:opEnd])
	_, ok := F[op]
	if !ok {
		outWait("Unknown operator: " + op)
		return
	}

	// Get the operands.

	as := input[:opStart]
	bs := input[opEnd:]

	// Validate the operands.

	if (as != "" && F[op].opType == prefix) || (bs != "" && F[op].opType == postfix) {
		outWait("Too many operands.")
		return
	}

	var a, b float64
	var err error

	if (F[op]).opType != prefix {
		if a, err = strconv.ParseFloat(as, 64); err != nil {
			outWait(as + "is not a well-formatted number")
			return
		}
	}

	if (F[op]).opType != postfix {
		if b, err = strconv.ParseFloat(bs, 64); err != nil {
			outWait(bs + " is not a well-formatted number.")
			return
		}
	}

	// Output.
	
	outWait(F[op].fn(a, b))
}