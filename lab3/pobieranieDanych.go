package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// *wskaźnik oznacza wartość wskazaną przez znacznik
	// &wskażnik oznacza wskaźnik danej zmiennej
	// uruchamianie z parametrami:    go run .\pobieranieDanych.go -liczba1 7 -liczba2 6

	def := -9999999999999
	wskaźnik1 := flag.Int("liczba1", def, "Liczba numer 1")
	wskaźnik2 := flag.Int("liczba2", def, "Liczba numer 2")
	flag.Parse()

	liczba1 := *wskaźnik1
	liczba2 := *wskaźnik2

	if liczba1 == def {
		fmt.Print("Podaj pierwszą liczbę: ")
		fmt.Scanf("%d\n", &liczba1)
	}

	if liczba2 == def {
		fmt.Print("Podaj drugą liczbę: ")
		fmt.Scanf("%d\n", &liczba2)
	}

	suma := liczba1 + liczba2
	fmt.Println("Liczby to", liczba1, "i", liczba2)
	fmt.Println("Suma", suma)

	// to już nie jest związane z programem, niemniej ciekawe
	fmt.Println(len(os.Args))
	fmt.Println(os.Args)
}
