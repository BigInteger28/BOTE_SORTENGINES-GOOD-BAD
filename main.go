package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)


// Pas willekeurige posities na de laatste ':' aan naar '0'
func randomizePositions(line string, numZetten int) string {
    parts := strings.Split(line, ":")
    if len(parts) < 3 {
        return line // Ongeldige regel
    }
    lastPart := parts[len(parts)-1]
    if len(lastPart) == 0 {
        return line // Niets om aan te passen
    }
    // Beperk het aantal zetten tot de lengte van het laatste deel
    if numZetten > len(lastPart) {
        numZetten = len(lastPart)
    }
    // Kies willekeurige unieke posities
    indices := rand.Perm(len(lastPart))[:numZetten]
    for _, idx := range indices {
        lastPart = lastPart[:idx] + "0" + lastPart[idx+1:]
    }
    parts[len(parts)-1] = lastPart
    return strings.Join(parts, ":")
}

// Controleer of een string alleen uit cijfers bestaat
func isDigits(s string) bool {
    for _, c := range s {
        if c < '0' || c > '9' {
            return false
        }
    }
    return true
}

func main() {
    // Open en lees input.txt
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Fout bij openen input.txt:", err)
        return
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    // Controleer of de eerste regel een getal is
    var firstLine string
    if len(lines) > 0 {
        firstLine = lines[0]
        if isDigits(firstLine) {
            // Verwijder de eerste regel uit lines
            lines = lines[1:]
        } else {
            firstLine = ""
        }
    }

    // Vraag de gebruiker om het aantal willekeurige zetten
    fmt.Print("Hoeveel random zetten: ")
    var numZetten int
    fmt.Scan(&numZetten)

    // Stel een seed in voor willekeurige getallen
    rand.Seed(time.Now().UnixNano())

    // Verwerk de regels en maak aangepaste kopieën
    var kopieLines []string
    for _, line := range lines {
        parts := strings.Split(line, ":")
        if len(parts) < 3 {
            continue // Sla ongeldige regels over
        }
        lastPart := parts[len(parts)-1]
        if !isDigits(lastPart) {
            continue // Sla regels met letters na de laatste ':' over
        }
        // Maak een kopie met '.' voor de eerste ':'
        kopie := strings.Replace(line, ":", ":", 1)
        // Pas willekeurige posities aan naar '0'
        kopie = randomizePositions(kopie, numZetten)
        kopieLines = append(kopieLines, kopie)
    }

    // Schrijf naar output.txt
    output, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Fout bij maken output.txt:", err)
        return
    }
    defer output.Close()

    writer := bufio.NewWriter(output)
    // Schrijf de eerste regel als het een getal is
    if firstLine != "" {
        fmt.Fprintln(writer, firstLine)
    }
    // Schrijf de aangepaste kopieën
    for _, kopie := range kopieLines {
        fmt.Fprintln(writer, kopie)
    }
    // Schrijf de originele regels
    //for _, line := range lines {
    //    fmt.Fprintln(writer, line)
    //}
    writer.Flush()
}
