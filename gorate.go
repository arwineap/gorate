package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var initial_line bool = true
	var previous_value int
	var output string

	// Scan through stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		current_value, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Fprintln(os.Stderr, "converting stdin to int:", err)
			// If current_value is not an int, goto next line
			continue
		}
		if initial_line {
			// Initial line coming in; no differential
			output = fmt.Sprintf("%d", current_value)
			initial_line = false
		} else {
			difference := current_value - previous_value
			if difference < 0 {
				// The number is negative so difference will automagically be prefixed
				// with -
				output = fmt.Sprintf("%d ( %d )", current_value, difference)
			} else {
				// The number is positive so we need to force prefixing of +
				output = fmt.Sprintf("%d ( +%d )", current_value, difference)
			}

		}
		previous_value = current_value
		fmt.Println(output)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
