package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile := "C:\\Users\\aleks\\Documents\\_workbench\\input\\SCR18319202411.TXT" // Nom du fichier à analyser
	outputFile := "C:\\Users\\aleks\\Documents\\_workbench\\output\\duplicates.txt"   // Fichier de sortie

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Erreur ouverture fichier:", err)
		return
	}
	defer file.Close()

	duplicates := make(map[string][]int) // Stocke les valeurs extraites et leurs numéros de ligne
	scanner := bufio.NewScanner(file)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		if len(line) < 384 { // Ignorer les lignes trop courtes
			continue
		}

		key := line[26:32] // Extraire la portion à comparer (colonnes 27-33)
		duplicates[key] = append(duplicates[key], lineNumber)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lecture fichier:", err)
		return
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Erreur création fichier de sortie:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for key, lines := range duplicates {
		if len(lines) > 1 {
			fmt.Fprintf(writer, "Valeur extraite: \"%s\"\nLignes: %s\n---\n", key, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(lines)), ", "), "[]"))
		}
	}
	writer.Flush()

	fmt.Println("Analyse terminée. Résultats dans", outputFile)
}
