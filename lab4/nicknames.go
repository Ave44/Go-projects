package main

import (
	"flag"
	"fmt"
	"strings"
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

	imieLitera := strings.ToLower(imie[:1])
	nazwiskoLitera := strings.ToLower(nazwisko[:1])
	tablicaDni := [31]string{ // przymiotnik
		"Szybki", "Jednoręki", "Bystry", "Przyczjony", "Wszechwidzący", "Hardy", "Cichy", "Pierwszy", "Żelaznoręki", "Zwinny",
		"Zielonoki", "Rudy", "Trwały", "Wolny", "Szary", "Czarny", "Głośny", "Biały", "Wielki", "Mały",
		"Przeklęty", "Ślepy", "Chrobry", "Złoty", "Blady", "Pomylony", "Szalony", "Cwany", "Ślepy", "Łysy",
		"Ostatni"}

	mapaImion := map[string]string{ // rzeczownik
		"a": "Anioł",
		"b": "Baron",
		"c": "Cesaż",
		"d": "Dyktator",
		"e": "Elizjanin",
		"f": "Fabryator",
		"g": "Gracz",
		"h": "Hohlik", // poprawnie pisze się Chohlik ale tu umyślnie dałem niepoprawną wersję
		"i": "Indyk",
		"j": "Jaguar",
		"k": "Kapłan",
		"l": "Lord",
		"ł": "Łotr",
		"m": "Maruder",
		"n": "Nożownik",
		"o": "Opiekun",
		"p": "Płomień",
		"r": "Ryceż",
		"s": "Smok",
		"t": "Trep",
		"u": "Uciekinier",
		"w": "Wieśniak",
		"x": "Kżyżowiec",
		"y": "Yeti",
		"z": "Zdrajca",
		"ź": "Kłamca",
		"ż": "Żniwiaż"}

	mapaNazwisk := map[string]string{ // przydomek
		"a": "Co okrążył ziemię",
		"b": "Kapłan prawdy",
		"c": "Na wieczność wygnany",
		"d": "Dziecie lasu",
		"e": "Wybaca z Salobaru",
		"f": "Co pokonał hydrę",
		"g": "Przyciągający pecha",
		"h": "Z pod góry Targon",
		"i": "Nenkany przez cienie",
		"j": "Wiecznie wracający",
		"k": "Wysłannik niebios",
		"l": "Syn wiedźmy",
		"ł": "Postrach goblinów",
		"m": "Niesłusznie oskarżony",
		"n": "Nie z tego świata",
		"o": "Tańczący na kurhanach wrogów",
		"p": "Bochater z pod Waterloo",
		"r": "Knujący spiski",
		"s": "Następca tronu",
		"t": "Poddany diabła",
		"u": "Innowator Raskanpolu",
		"w": "Co stracił oko",
		"x": "Co przetrwał atak Tarkian",
		"y": "Przybysz z Północy",
		"z": "Co uciekł przed śmiercią",
		"ź": "Przybysz z Południa",
		"ż": "Wybraniec niebois"}

	fmt.Println(tablicaDni[dzień-1], mapaImion[imieLitera], mapaNazwisk[nazwiskoLitera])
}
