package main

import "fmt"

func main() {
	x := make([]string, 50, 50)
	x = []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `NewHampshire`, `NewJersey`, `NewMexico`, `NewYork`, `NorthCarolina`, `NorthDakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `RhodeIsland`, `SouthCarolina`, `SouthDakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `WestVirginia`, `Wisconsin`, `Wyoming`}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
