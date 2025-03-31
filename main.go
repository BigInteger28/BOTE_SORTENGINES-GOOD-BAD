package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Lees bote.txt en verzamel de volgorde van namen
	file, err := os.Open("bote.txt")
	if err != nil {
		fmt.Println("Fout bij het openen van bote.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var orderNames []string
	for scanner.Scan() {
		line := scanner.Text()
		// Negeer lijnen die beginnen met ---
		if strings.HasPrefix(line, "---") {
			continue
		}
		// Splits de lijn op "   " (drie spaties), met een maximum van 4 delen
		parts := strings.SplitN(line, "   ", 4)
		if len(parts) >= 3 {
			name := parts[0] // De naam is het eerste deel vóór de drie spaties
			orderNames = append(orderNames, name)
		} else {
			fmt.Println("Ongeldige lijn:", line)
		}
	}

	// 2. Lees engines13.txt en maak een map van naam naar volledige lijn
	file2, err := os.Open("engines13.txt")
	if err != nil {
		fmt.Println("Fout bij het openen van engines13.txt:", err)
		return
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)
	engineLines := make(map[string]string)
	for scanner2.Scan() {
		line := scanner2.Text()
		// Splits op de eerste : om de naam te krijgen
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue // Ongeldige lijn
		}
		name := parts[0]
		engineLines[name] = line
	}

	// 3. Verzamel de gesorteerde lijnen
	var sortedLines []string
	for _, name := range orderNames {
		if line, exists := engineLines[name]; exists {
			sortedLines = append(sortedLines, line)
		} else {
			fmt.Printf("Naam niet gevonden in engines13.txt: %s\n", name)
		}
	}

	// 4. Schrijf naar sortedEngines.txt
	outputFile, err := os.Create("sortedEngines.txt")
	if err != nil {
		fmt.Println("Fout bij het aanmaken van sortedEngines.txt:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range sortedLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Println("Klaar! De gesorteerde lijst is opgeslagen in sortedEngines.txt.")
}
