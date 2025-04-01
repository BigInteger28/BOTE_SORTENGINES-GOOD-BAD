package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// **Deel 1: Update de levels in engines13.txt op basis van bote.txt**
	
	// Stap 1.1: Open en lees bote.txt om levels op te slaan in een map
	boteFile, err := os.Open("bote.txt")
	if err != nil {
		fmt.Println("Fout bij het openen van bote.txt:", err)
		return
	}
	defer boteFile.Close()

	// Map om engine-naam te koppelen aan level
	nameToLevel := make(map[string]string)
	scanner := bufio.NewScanner(boteFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "---") {
			continue // Sla scheidingslijnen over
		}
		parts := strings.SplitN(line, "   ", 4) // Splits op drie spaties
		if len(parts) >= 3 {
			name := parts[0]  // Engine-naam
			level := parts[1] // Level
			nameToLevel[name] = level
		}
	}

	// Stap 1.2: Open en lees engines13.txt om levels te updaten
	enginesFile, err := os.Open("engines13.txt")
	if err != nil {
		fmt.Println("Fout bij het openen van engines13.txt:", err)
		return
	}
	defer enginesFile.Close()

	scanner = bufio.NewScanner(enginesFile)
	var updatedLines []string
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 3) // Splits op ":" (naam:level:cijfers)
		if len(parts) < 3 {
			continue // Ongeldige regel, sla over
		}
		name := parts[0] // Engine-naam
		if newLevel, exists := nameToLevel[name]; exists {
			// Update level als de engine in bote.txt voorkomt
			updatedLine := fmt.Sprintf("%s:%s:%s", name, newLevel, parts[2])
			updatedLines = append(updatedLines, updatedLine)
		} else {
			// Behoud originele regel als de engine niet in bote.txt staat
			updatedLines = append(updatedLines, line)
		}
	}

	// Stap 1.3: Schrijf bijgewerkte regels naar updated_engines13.txt
	updatedFile, err := os.Create("updated_engines13.txt")
	if err != nil {
		fmt.Println("Fout bij het aanmaken van updated_engines13.txt:", err)
		return
	}
	defer updatedFile.Close()

	writer := bufio.NewWriter(updatedFile)
	for _, line := range updatedLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Println("Deel 1 klaar: Levels zijn bijgewerkt in updated_engines13.txt.")

	// **Deel 2: Sorteer de bijgewerkte engines op basis van bote.txt**
	
	// Stap 2.1: Lees bote.txt opnieuw om de volgorde van namen te bepalen
	boteFile.Seek(0, 0) // Terug naar het begin van bote.txt
	scanner = bufio.NewScanner(boteFile)
	var orderNames []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "---") {
			continue
		}
		parts := strings.SplitN(line, "   ", 4)
		if len(parts) >= 3 {
			name := parts[0]
			orderNames = append(orderNames, name)
		}
	}

	// Stap 2.2: Lees updated_engines13.txt en koppel namen aan regels
	updatedEnginesFile, err := os.Open("updated_engines13.txt")
	if err != nil {
		fmt.Println("Fout bij het openen van updated_engines13.txt:", err)
		return
	}
	defer updatedEnginesFile.Close()

	scanner = bufio.NewScanner(updatedEnginesFile)
	engineLines := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			continue
		}
		name := parts[0]
		engineLines[name] = line
	}

	// Stap 2.3: Maak een gesorteerde lijst volgens de volgorde in bote.txt
	var sortedLines []string
	for _, name := range orderNames {
		if line, exists := engineLines[name]; exists {
			sortedLines = append(sortedLines, line)
		} else {
			fmt.Printf("Naam niet gevonden in updated_engines13.txt: %s\n", name)
		}
	}

	// Stap 2.4: Schrijf de gesorteerde lijst naar sorted_updated_engines.txt
	sortedFile, err := os.Create("sortedEngines.txt")
	if err != nil {
		fmt.Println("Fout bij het aanmaken van sorted_updated_engines.txt:", err)
		return
	}
	defer sortedFile.Close()

	writer = bufio.NewWriter(sortedFile)
	for _, line := range sortedLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Println("Deel 2 klaar: De gesorteerde lijst is opgeslagen in sorted_updated_engines.txt.")
}
