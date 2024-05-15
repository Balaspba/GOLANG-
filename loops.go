package main

import (
	"fmt"
)

func playerss() {

}
func games() {
	var ballgame string = "football"
	var person int = 11
	if person < 11 {
		fmt.Println(" Needed one player", ballgame)
		return

	} else if person >= 11 {

		fmt.Println(" all player available", ballgame)
		playerss()
	}

	var day int = 5
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
		playerss()
	default:
		fmt.Println("Invalid")

	}
	var a int = 0
	var b int
	for a = 0; a < 5; a++ {
		b = a
		a++
		fmt.Println("players", b)
		playerss()

	}
	var sub int = 0
	for sub = 0; sub < 5; sub++ {

		if sub < 4 {
			fmt.Println("sub players ok")
			break
		} else if sub <= 2 {
			fmt.Println("sub players not ok")
			continue
		}
	}
	var coach int

	var refree int
	for coach = 0; coach < 5; coach++ {
		if coach <= 5 {
			fmt.Println("coach all here")
			continue
		}
		refree = coach + 1
		for refree = 0; refree < coach; refree++ {
			if refree == 4 {
				fmt.Println("refree all here")

			}
			playerss()
		}
	}
	return
}
func main() {
	games()
}
