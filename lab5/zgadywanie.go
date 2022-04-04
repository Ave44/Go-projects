package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getUserNumber(komunikat string) (int, bool) {
	liczbaString := ""
	fmt.Print(komunikat)
	fmt.Scan(&liczbaString)

	if liczbaString == "koniec" {
		return -1, false
	}

	liczba, err := strconv.Atoi(liczbaString)
	if err == nil && liczba >= 0 && liczba <= 1000 {
		return liczba, true
	}
	return getUserNumber("Należy podać liczbę od 0 do 1000: ")
}

func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func victory() bool {
	anwser := ""
	fmt.Print("Gratulacje zgadłeś!\n")
	for anwser != "T" && anwser != "N" {
		fmt.Print("Czy chcesz kontynuować? [T/N]: ")
		fmt.Scan(&anwser)
	}
	if anwser == "T" {
		return true
	}
	return false
}

func game() {
	play := true

MainLoop:
	for play {
		randNum := getRandomNumber(0, 5)
		liczba, continueGame := getUserNumber("Podaj liczbę od 0 do 1000: ")
		if !continueGame {
			break MainLoop
		}

		for liczba != randNum {
			if liczba < randNum {
				liczba, continueGame = getUserNumber("Podano za małą liczbę: ")
			} else {
				liczba, continueGame = getUserNumber("Podano za dużą liczbę: ")
			}
			if !continueGame {
				break MainLoop
			}
		}

		play = victory()
	}

	fmt.Print("Dziękuję za grę!")
}

func main() {

	game()
}
