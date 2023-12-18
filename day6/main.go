package main

import "fmt"

func getWaysToWin(time int, recordDistance int) int {
	waysToWin := 0

	for i := 0; i < time; i++ {
		speed := i
		running := time - i

		distance := running * speed

		if distance > recordDistance {
			waysToWin += 1
		}
	}

	return waysToWin
}

func main() {
	time1, time2, time3, time4 := 45, 97, 72, 95
	dist1, dist2, dist3, dist4 := 305, 1062, 1110, 1695

	waysToWin1 := getWaysToWin(time1, dist1)
	waysToWin2 := getWaysToWin(time2, dist2)
	waysToWin3 := getWaysToWin(time3, dist3)
	waysToWin4 := getWaysToWin(time4, dist4)

	fmt.Println("Total ways to win for Part 1:", waysToWin1*waysToWin2*waysToWin3*waysToWin4)

	timeN := 45977295
	distN := 305106211101695

	waysToWinN := getWaysToWin(timeN, distN)

	fmt.Println("Total ways to win for Part 2:", waysToWinN)
}
