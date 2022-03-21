package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	wskaźnik1 := flag.Float64("tmp", 0, "Temperatura")
	wskaźnik2 := flag.String("typ", "C", "Skala")
	flag.Parse()

	tmp := *wskaźnik1
	typ := *wskaźnik2

	if len(os.Args) >= 2 {
		if s, err := strconv.ParseFloat(os.Args[1], 32); err == nil {
			tmp = s
		} else {
			fmt.Print("Podaj temperaturę: ")
			fmt.Scanf("%f\n", &tmp)
		}
	} else {
		fmt.Print("Podaj temperaturę: ")
		fmt.Scanf("%f\n", &tmp)
	}

	if len(os.Args) >= 3 {
		typ = os.Args[2]
	} else {
		fmt.Print("Podaj domyślną skalę [C, F, K]: ")
		fmt.Scanf("%s\n", &typ)
	}

	if typ == "C" {
		fmt.Printf("%.1f °C\n", tmp)
		fmt.Printf("%.1f °F\n", tmp*1.8+32)
		fmt.Printf("%.1f °K\n", tmp+273.15)
	} else if typ == "F" {
		fmt.Printf("%.1f °C\n", (tmp-32)/1.8)
		fmt.Printf("%.1f °F\n", tmp)
		fmt.Printf("%.1f °K\n", (tmp+459.67)*5/9)
	} else {
		fmt.Printf("%.1f °C\n", tmp-273.15)
		fmt.Printf("%.1f °F\n", (tmp*1.8)-459.67)
		fmt.Printf("%.1f °K\n", tmp)
	}
}
