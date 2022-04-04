package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func getPlayerName() string {
	name := ""
	fmt.Print("Podaj swoje imie: ")
	fmt.Scan(&name)

	if name != "" {
		return name
	}
	return getPlayerName()
}

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

func victory(player string, guessCount int, num int) bool {
	anwser := ""
	if guessCount == 1 {
		fmt.Print("Gratulacje ", player, " zgadłeś w pierwszej próbie!\n")
	} else {
		fmt.Print("Gratulacje ", player, " zgadłeś w ", guessCount, " próbach!\n")
	}
	for anwser != "T" && anwser != "N" {
		fmt.Print("Czy chcesz kontynuować? [T/N]: ")
		fmt.Scan(&anwser)
	}
	if anwser == "T" {
		return true
	}
	return false
}

func game(results map[string]int) {
	play := true

MainLoop:
	for play {
		randNum := getRandomNumber(0, 5)
		player := getPlayerName()
		liczba, continueGame := getUserNumber("Podaj liczbę od 0 do 1000: ")
		if !continueGame {
			break MainLoop
		}
		guessCount := 1

		for liczba != randNum {
			guessCount++
			if liczba < randNum {
				liczba, continueGame = getUserNumber("Podano za małą liczbę: ")
			} else {
				liczba, continueGame = getUserNumber("Podano za dużą liczbę: ")
			}
			if !continueGame {
				break MainLoop
			}
		}

		record, includes := results[player]
		if !includes || record > guessCount {
			results[player] = guessCount
		}
		play = victory(player, guessCount, randNum)
	}

	fmt.Print("Dziękuję za grę!\n")
}

func main() {
	results := make(map[string]int)

	game(results)

	fmt.Print("-----Wyniki----- \n")

	hallOfFame := make([]string, 0, len(results))
	for key := range results {
		hallOfFame = append(hallOfFame, key)
	}

	sort.Slice(hallOfFame, func(i, j int) bool {
		return results[hallOfFame[i]] < results[hallOfFame[j]]
	})

	for _, name := range hallOfFame {
		fmt.Print(name, "\t", results[name], "\n")
	}
}
