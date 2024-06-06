package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var colorFlag = flag.String("color", "default", "Specify the color (black, red, green, yellow, blue, magenta, cyan, white, default)")

func convertInput(line string, lines []string, coloredLetters string) {
	if line != "" {
		for i := 1; i < 9; i++ {
			for _, char := range line {
				ascii := int(char) - 32
				if strings.ContainsRune(coloredLetters, char) {
					printColoredCharacter(lines[ascii*9+i], *colorFlag)
				} else {
					printColoredCharacter(lines[ascii*9+i], "default")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println()
	}
}

func printColoredCharacter(character string, color string) {
	colorMap := map[string]int{
		"default": 0,
		"black":   30,
		"red":     31,
		"green":   32,
		"yellow":  33,
		"blue":    34,
		"magenta": 35,
		"cyan":    36,
		"white":   37,
	}

	if code, ok := colorMap[color]; ok {
		fmt.Printf("\033[%dm%s\033[0m", code, character)
	} else {
		fmt.Println("Please choose from: black, red, green, yellow, blue, magenta, cyan, white, default")
		os.Exit(1)
	}
}

func main() {
	var fontFile string

	// Parse les drapeaux
	flag.Parse()

	// Vérifier si un argument est fourni
	if flag.NArg() < 2 {
		fmt.Println("Usage: go run . --color=<color> <letters to be colored> <string>")
		return
	}

	// Si un argument supplémentaire est fourni, utilisez-le comme nom de fichier de police
	if flag.NArg() == 3 {
		fontFile = flag.Arg(2)
	} else {
		fontFile = "standard" // fichier de police par défaut
	}

	data, err := ioutil.ReadFile(fontFile + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(data)

	lines := strings.Split(text, "\n")
	inputLines := strings.Split(flag.Arg(1), "\\n")
	coloredLetters := flag.Arg(0)

	for _, inputLine := range inputLines {
		convertInput(inputLine, lines, coloredLetters)
	}
}
