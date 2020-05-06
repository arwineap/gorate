package main

import (
	"fmt"
	"time"

	"github.com/arwineap/gorate/gorate"
)

func handler(currentValue int64, difference int64, duration string, cumDifference int64, cumDuration string) {
	plusSign := ""
	if difference > 0 {
		plusSign = "+"
	}
	cumPlusSign := ""
	if cumDifference > 0 {
		cumPlusSign = "+"
	}

	fmt.Printf("%d ( %s%d per %s ) ( %s%d per %s )\n", currentValue, plusSign, difference, duration, cumPlusSign, cumDifference, cumDuration)
}

func main() {
	client := gorate.NewClient(handler)
	client.NewEntry(1000)
	time.Sleep(1 * time.Second)
	client.NewEntry(1001)
	time.Sleep(1 * time.Second)
	client.NewEntry(1002)
	time.Sleep(1 * time.Second)
	client.NewEntry(1003)
	time.Sleep(7 * time.Second)
	client.NewEntry(1010)
	time.Sleep(1 * time.Second)
	client.NewEntry(1002)

	/*
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			i, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "converting stdin to int:", err)
				continue
			}
			client.NewEntry(i)
		}
	*/
}
