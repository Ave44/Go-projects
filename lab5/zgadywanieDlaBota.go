package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Fprint(os.Stderr, "-----Gra uruchomiona----- \n")
	play := true

	for play {
		randNum := getRandomNumber(0, 1000)

		fmt.Fprint(os.Stderr, "name\n")
		fmt.Println("name")

		scanner.Scan()
		player := scanner.Text()

		guessCount := 1

		for scanner.Scan() {
			liczbaString := scanner.Text()
			fmt.Fprint(os.Stderr, "Podano: ", liczbaString, "\n")
			liczba, _ := strconv.Atoi(liczbaString)

			time.Sleep(3 * time.Second)
			guessCount++
			if liczba < randNum {
				fmt.Fprint(os.Stderr, "m", liczba, liczbaString, "|\n")
				fmt.Println("m")
				scanner.Scan()
				liczbaString := scanner.Text()
				liczba, _ = strconv.Atoi(liczbaString)
			}
			if liczba > randNum {
				fmt.Fprint(os.Stderr, "d\n")
				fmt.Println("d")
				scanner.Scan()
				liczbaString := scanner.Text()
				liczba, _ = strconv.Atoi(liczbaString)
			}
			if liczba == randNum {
				break
			}
		}

		fmt.Fprint(os.Stderr, player, " zgadł ", randNum, " w ", guessCount, " próbach\n")
		scanner.Scan()
		anwser := scanner.Text()

		if anwser == "T" {
			play = true
		}
		play = false
	}
}
