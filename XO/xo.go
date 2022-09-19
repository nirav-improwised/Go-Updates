package main

import "fmt"

func main() {
	fmt.Println("----------------------Welcome to X O game----------------------------")
	var player1i string
	var player1j string
	var player2i string
	var player2j string
	var playground [3][3]string
	var available []string
	fmt.Println("Player 1-X: Enter single space separated row and column. Ex: 1 1")
	fmt.Scan(&player1i, &player1j)

	fmt.Println("Player 2-O: Enter single space separated row and column. Ex: 1 1")
	fmt.Scan(&player2i, &player2j)

}
