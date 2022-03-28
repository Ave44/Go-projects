package main

import (
	"flag"
	"fmt"
)

func main() {
	// go run nicknames.go -imie John -nazwisko Doe -dzień 15

	def := "Default"
	wskaźnikImie := flag.String("imie", def, "Imie")
	wskaźnikNazwisko := flag.String("nazwisko", def, "Nazwisko")
	wskaźnikDzień := flag.Int("dzień", 0, "Dzień")
	flag.Parse()

	imie := *wskaźnikImie
	nazwisko := *wskaźnikNazwisko
	dzień := *wskaźnikDzień

	for imie == def && imie != "" {
		fmt.Print("Podaj swoje imie: ")
		// fmt.Scanf("%s\n", &imie) // inna opcja pobierania parametrów
		fmt.Scan(&imie)
	}

	for nazwisko == def && nazwisko != "" {
		fmt.Print("Podaj nazwisko: ")
		fmt.Scan(&nazwisko)
	}

	for dzień <= 0 {
		fmt.Print("Podaj dzień: ")
		fmt.Scan(&dzień)
	}

	fmt.Println(imie, nazwisko, dzień)

	imieLitera := imie[:1]
	nazwiskoLitera := nazwisko[:1]
	tablicaDni := [31]string{
		"szybki","jednoręki","bystry","przyczjony","wszechwidzący","hardy","cichy","pierwszy","żelaznoręki","zwinny",
		"zielonoki","rudy","trwały","wolny","szary","czarny","głośny","biały","wielki","mały",
		"przeklęty","ślepy","chrobry","złoty","blady","pomylony","szalony","cwany","ślepy","łysy",
		"ostatni"}

	mapaImion = map[string]string{
		"a":"",
		"b":"",
		"c":"",
		"d":"",
		"e":"",
		"f":"",
		"g":"",
		"h":"",
		"i":"",
		"j":"",
		"k":"",
		"l":"",
		"ł":"",
		"m":"",
		"n":"",
		"o":"",
		"p":"",
		"r":"",
		"s":"",
		"t":"",
		"u":"",
		"w":"",
		"x":"",
		"y":"",
		"z":"",
		"ź":"",
		"ż":""
	}

	fmt.Println(imieLitera, nazwiskoLitera, tablicaDni[dzień-1])
}