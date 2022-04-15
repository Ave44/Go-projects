package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	min := 0
	max := 1000
	currentGuess := (max + min) / 2

	fmt.Fprintf(os.Stderr, "Program zgadujący uruchominy\n")

	for scanner.Scan() {
		time.Sleep(3 * time.Second)
		fmt.Fprintf(os.Stderr, "Pętla została uruchomiona\n")

		line := scanner.Text()
		fmt.Fprintf(os.Stderr, line, "\n")

		if strings.Contains(line, "name") {
			fmt.Println("GuessBot")
			return
		}
		if strings.Contains(line, "next") {
			fmt.Println("N")
			return
		}
		if strings.Contains(line, "number") {
			fmt.Println(currentGuess)
			fmt.Fprintf(os.Stderr, "Pierwsza proba: "+strconv.Itoa(min)+" "+strconv.Itoa(currentGuess)+" "+strconv.Itoa(max)+"\n")
			continue
		}
		if strings.Contains(line, "m") {
			min = currentGuess + 1
			if (max+min)%2 == 0 {
				currentGuess = (max + min) / 2
			} else {
				currentGuess = (max + min + 1) / 2
			}
			fmt.Println(currentGuess)
			fmt.Fprintf(os.Stderr, "m "+strconv.Itoa(min)+" "+strconv.Itoa(currentGuess)+" "+strconv.Itoa(max)+"\n")
			continue
		}
		if strings.Contains(line, "d") {
			max = currentGuess - 1
			if (max+min)%2 == 0 {
				currentGuess = (max + min) / 2
			} else {
				currentGuess = (max + min + 1) / 2
			}
			fmt.Println(currentGuess)
			fmt.Fprintf(os.Stderr, "d "+strconv.Itoa(min)+" "+strconv.Itoa(currentGuess)+" "+strconv.Itoa(max)+"\n")
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		return
	}
}
