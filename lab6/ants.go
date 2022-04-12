package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Map struct {
	fields map[string]Field
	ants   []Ant
	leafs  []Leaf
	width  int
	hight  int
}

type Field struct {
	ant  bool
	leaf bool
}

type Ant struct {
	x        int
	y        int
	haveLeaf bool
}

type Leaf struct {
	x int
	y int
}

func (simMap Map) draw() {
	for i := 0; i < simMap.hight; i++ {
		for j := 0; j < simMap.width; j++ {
			coordinate := strconv.Itoa(j) + "." + strconv.Itoa(i)
			drawField(simMap.fields[coordinate])
		}
		fmt.Print("\n\n")
	}
}

func drawField(f Field) {
	if f.ant {
		if !f.leaf {
			fmt.Print(string("\033[31m"), " Ӂ")
			fmt.Print(string("\033[37m"), "  ")
		} else {
			fmt.Print(string("\033[31m"), " Ӂ")
			fmt.Print(string("\033[32m"), "❀")
			fmt.Print(string("\033[37m"), " ")
		}

	} else if f.leaf {
		fmt.Print(string("\033[32m"), " ❀")
		fmt.Print(string("\033[37m"), "  ")
	} else {
		fmt.Print(string("\033[37m"), "░░░ ")
	}
}

func getNewRandAntPosition(m Map, ants []Ant) string {
	foundNewPosition := false
	coordinate := "0.0"
	for !foundNewPosition {
		foundNewPosition = true
		x, y := getRandPosition(m)

		for _, value := range ants {
			if value.x == x && value.y == y {
				foundNewPosition = false
				break
			}
		}
		coordinate = strconv.Itoa(x) + "." + strconv.Itoa(y)
	}

	return coordinate
}

func getRandPosition(m Map) (int, int) {
	rand.Seed(time.Now().UnixNano())
	xMax := m.width
	yMax := m.hight

	x := rand.Intn(xMax)
	y := rand.Intn(yMax)

	return x, y
}

func getCoordinates(s string) (int, int) {
	arr := strings.Split(s, ".")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	return x, y
}

func getNewRandLeafPosition(m Map, leafs []Leaf) string {
	foundNewPosition := false
	coordinate := "0.0"
	for !foundNewPosition {
		foundNewPosition = true
		x, y := getRandPosition(m)

		for _, value := range leafs {
			if value.x == x && value.y == y {
				foundNewPosition = false
				break
			}
		}
		// uniemożliwienie liściu zespawnowanie się na mrówce
		for _, value := range m.ants {
			if value.x == x && value.y == y {
				foundNewPosition = false
				break
			}
		}
		coordinate = strconv.Itoa(x) + "." + strconv.Itoa(y)
	}

	return coordinate
}

func (simMap Map) tick() {
	fmt.Print("\033[H\033[2J")
	for i, ant := range simMap.ants {
		ant.move(simMap, i)
	}
	simMap.draw()
}

func (ant Ant) move(simMap Map, antIndex int) {
	//rand.Seed(time.Now().UnixNano())
	direction := rand.Intn(4)
	ant = moveIfCan(ant, direction, simMap, antIndex)
}

func moveIfCan(ant Ant, d int, simMap Map, antIndex int) Ant {
	oldCoordinates := strconv.Itoa(ant.x) + "." + strconv.Itoa(ant.y)
	newCoordinates := ""
	switch d {
	case 0:
		newCoordinates = strconv.Itoa(ant.x) + "." + strconv.Itoa(ant.y-1)
	case 1:
		newCoordinates = strconv.Itoa(ant.x+1) + "." + strconv.Itoa(ant.y)
	case 2:
		newCoordinates = strconv.Itoa(ant.x) + "." + strconv.Itoa(ant.y+1)
	case 3:
		newCoordinates = strconv.Itoa(ant.x-1) + "." + strconv.Itoa(ant.y)
	}
	x, y := getCoordinates(newCoordinates)
	if x >= 0 && y >= 0 && x < simMap.width && y < simMap.hight && simMap.fields[newCoordinates].ant == false {
		ant.x = x
		ant.y = y
		if value, ok := simMap.fields[oldCoordinates]; ok {
			value.ant = false
			simMap.fields[oldCoordinates] = value
		}
		if value, ok := simMap.fields[newCoordinates]; ok {
			value.ant = true
			simMap.fields[newCoordinates] = value
		}
		simMap.ants[antIndex] = ant
	}

	return ant
}

func main() {

	///// inicjalizacja mapy /////
	var simMap Map
	fields := make(map[string]Field)

	simMap.width = 10
	simMap.hight = 5

	for i := 0; i < simMap.width; i++ {
		for j := 0; j < simMap.hight; j++ {
			coordinate := strconv.Itoa(i) + "." + strconv.Itoa(j)
			fields[coordinate] = Field{}
		}
	}
	simMap.fields = fields

	///// inicjalizacja mrówek /////
	antAmount := 3
	ants := make([]Ant, antAmount)

	// ustawienie "defaultowych wartości" to pozwala mi na zespawnowanie mrówki na polu 0.0
	for i := range ants {
		ants[i].x = -1
		ants[i].y = -1
	}

	for i := 0; i < antAmount; i++ {
		coordinate := getNewRandAntPosition(simMap, ants)
		x, y := getCoordinates(coordinate)

		// zmiana wartości pola w mapie
		if value, ok := simMap.fields[coordinate]; ok {
			value.ant = true
			simMap.fields[coordinate] = value
		}
		ants[i] = Ant{x: x, y: y}
	}

	simMap.ants = ants

	///// inicjalizacja liści /////
	leafAmount := 4
	leafs := make([]Leaf, leafAmount)

	// ustawienie "defaultowych wartości" to pozwala mi na zespawnowanie liścia na polu 0.0
	for i := range leafs {
		leafs[i].x = -1
		leafs[i].y = -1
	}

	for i := 0; i < leafAmount; i++ {
		coordinate := getNewRandLeafPosition(simMap, leafs)
		x, y := getCoordinates(coordinate)

		// zmiana wartości pola w mapie
		if value, ok := simMap.fields[coordinate]; ok {
			value.leaf = true
			simMap.fields[coordinate] = value
		}
		leafs[i] = Leaf{x: x, y: y}
	}

	simMap.leafs = leafs

	simMap.draw()
	for i := 0; i < 50; i++ {
		time.Sleep(200 * time.Millisecond)
		simMap.tick()
	}

}
