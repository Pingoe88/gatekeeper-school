package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func checkName(kenteken string) bool {
	allowedkenteken := []string{"82nnk5", "nk76b5", "nnk825", "825nnk"}

	kenteken = strings.ToLower(kenteken)
	for _, allowed := range allowedkenteken {
		if kenteken == allowed {
			return true
		}
	}

	return false
}

func main() {
	welkomStr := "Welkom bij Fonteyn Vakantieparken"
	currentTime := time.Now().Format("15:04")

	fmt.Print("Voer je kenteken in: ")
	reader := bufio.NewReader(os.Stdin)
	kenteken, _ := reader.ReadString('\n')
	kenteken = strings.TrimSpace(kenteken)

	if checkName(kenteken) {
		if currentTime >= "07:00" && currentTime <= "12:00" {
			fmt.Println("Goedemorgen!:", welkomStr)
		} else if currentTime > "12:00" && currentTime <= "18:00" {
			fmt.Println("Goedemiddag:", welkomStr)
		} else if currentTime > "18:00" && currentTime <= "23:00" {
			fmt.Println("Goedenavond:", welkomStr)
		} else {
			fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten.")
		}
	} else {
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein.")
	}
}
