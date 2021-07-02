package main

import "fmt"

func main() {
	var winner, player, temp string
	
	var i, ronde int
	var nilai, answer int

	winner = "A"
	player = "B"

	ronde = 1

	fmt.Printf("Ronde %d:\n", ronde)
	fmt.Printf("%s - masukan angka rahasia: ", winner)
	fmt.Scanln(&nilai)

	for nilai != -101 {
		i = 1

		fmt.Printf("%s - masukan angka tebakan ke-%d: ", player, i)
		fmt.Scanln(&answer)

		for i < 3 && answer != nilai {
			i++

			fmt.Printf("%s - masukan angka tebakan ke-%d: ", player, i)
			fmt.Scanln(&answer)
		}

		if nilai == answer {
			temp = winner
			winner = player
			player = temp
		}

		fmt.Printf("%s adalah pemenangnya\n", winner)

		ronde++ 

		fmt.Printf("\nRonde %d:\n", ronde)
		fmt.Printf("%s - masukan angka rahasia: ", winner)
		fmt.Scanln(&nilai)
	}

	fmt.Println("Permainan Selesai")
}