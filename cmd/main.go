package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Engine started")

	for {
		var input string
		fmt.Scan(&input)

		switch(input) {
		case "isready":
			fmt.Println("readyok")
		case "uci":
			fmt.Println("id name vhcess")
		case "quit":
			os.Exit(0)
		}
	}
}
