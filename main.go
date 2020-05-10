package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/arwineap/gorate/gorate"
)

func main() {
	client := gorate.NewClient()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "converting stdin to int:", err)
			continue
		}
		fmt.Print(client.NewEntry(i))
	}

}
