package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Stap 1: Lees de volgorde uit input2.txt
	var rankOrder []string
	file2, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Fout bij openen van input2.txt:", err)
		return
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		line := strings.TrimSpace(scanner2.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}
		// De naam is alles behalve de laatste twee delen (scores)
		name := strings.Join(parts[:len(parts)-2], " ")
		rankOrder = append(rankOrder, name)
	}

	// Stap 2: Lees input1.txt en maak een map van volledige namen naar regels
	engineMap := make(map[string]string)
	file1, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Fout bij openen van input1.txt:", err)
		return
	}
	defer file1.Close()

	scanner1 := bufio.NewScanner(file1)
	for scanner1.Scan() {
		line := strings.TrimSpace(scanner1.Text())
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}
		name := strings.TrimSpace(parts[0])
		engineMap[name] = line
	}

	// Stap 3: Schrijf de output in de exacte volgorde van input2
	outputFile, err := os.Create("sortEngines.txt")
	if err != nil {
		fmt.Println("Fout bij maken van sortEngines.txt:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, name := range rankOrder {
		if line, exists := engineMap[name]; exists {
			fmt.Fprintln(writer, line)
		}
	}
	writer.Flush()
}
