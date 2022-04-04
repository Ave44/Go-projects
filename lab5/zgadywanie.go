package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func getUserNumber(komunikat string) int {
	liczbaString := ""
	fmt.Print(komunikat)
	fmt.Scan(&liczbaString)

	liczba, err := strconv.Atoi(liczbaString)
	if err == nil && liczba >= 0 && liczba <= 1000 {
		return liczba
	}
	return getUserNumber("Należy podać liczbę od 0 do 1000: ")
}

func main() {

	liczba := -1
	max := 1000
	min := 0

	rand.Seed(time.Now().UnixNano())
	loswa := rand.Intn(max - min) + min

	liczba = getUserNumber("Podaj liczbę od 0 do 1000: ")

	for liczba != loswa {
		if liczba < loswa {
			liczba = getUserNumber("Podano za małą liczbę: ")
		} else {
			liczba = getUserNumber("Podano za dużą liczbę: ")
		}
	}

	fmt.Print("Gratulacje zgadłeś!")

}