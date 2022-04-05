package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BestResult struct {
	guessCount int
	date       string
	randNum    int
}

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

func game(results map[string]BestResult) {
	play := true

MainLoop:
	for play {
		randNum := getRandomNumber(0, 1000)
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
		if !includes || record.guessCount > guessCount {
			if record.guessCount > guessCount {
				fmt.Print("Gratulacje ", player, " udało ci się pobić swój rekord!\n")
			}
			var res BestResult
			res.guessCount = guessCount
			res.date = time.Now().Format("02.01.2006")
			res.randNum = randNum
			results[player] = res
		}
		play = victory(player, guessCount, randNum)
	}

	fmt.Print("Dziękuję za grę!\n")
}

func writeToFile(log string) {
	file, err := os.Create("logs.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(log)
}

func getResults() map[string]BestResult {
	results := make(map[string]BestResult)
	logs, err := os.ReadFile("logs.txt")
	if err == nil {
		text := fmt.Sprintf("%s", logs)
		table := strings.Fields(text)
		for i := 0; i < len(table); i += 4 {
			var res BestResult
			res.guessCount, _ = strconv.Atoi(table[i+1])
			res.date = table[i+2]
			res.randNum, _ = strconv.Atoi(table[i+3])
			results[table[i]] = res
		}
	}
	return results
}

func main() {
	results := getResults()

	game(results)

	fmt.Print("-----Wyniki----- \n")

	hallOfFame := make([]string, 0, len(results))
	for key := range results {
		hallOfFame = append(hallOfFame, key)
	}

	sort.Slice(hallOfFame, func(i, j int) bool {
		return results[hallOfFame[i]].guessCount < results[hallOfFame[j]].guessCount
	})

	log := ""

	for _, name := range hallOfFame {
		row := fmt.Sprintf("%-10v %-3v %v %v\n", name, results[name].guessCount, results[name].date, results[name].randNum)
		if len(name) > 10 {
			row = fmt.Sprintf("%-10v %-3v %v %v\n", name[:9], results[name].guessCount, results[name].date, results[name].randNum)
		}
		fmt.Print(row)
		log += row
	}

	writeToFile(log)
}
