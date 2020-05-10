package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/arwineap/gorate/gorate"
)

func main() {
	client := gorate.NewClient()

	formatter := flag.String("formatter", "full", "specify a formatter (full|cumulative|instantaneous)")
	flag.Parse()
	switch f := strings.ToLower(*formatter); f {
	case "full":
		client.SetFormatter(gorate.FullFormatter{})
	case "cumulative":
		client.SetFormatter(gorate.CumulativeFormatter{})
	case "instantaneous":
		client.SetFormatter(gorate.InstantaneousFormatter{})
	default:
		fmt.Fprintln(os.Stderr, "WARN: invalid formatter specified, defaulting to full")
	}

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
