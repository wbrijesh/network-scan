package oui

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"scan-network/database"
)

func ScanDataFromTextFile(filePath string, db *database.Database) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regexBase16 := regexp.MustCompile(`([0-9A-Fa-f]{6})\s+\(base 16\)\s+(.*)`)

	var organisation string
	index := 1
	for scanner.Scan() {
		line := scanner.Text()

		if matches := regexBase16.FindStringSubmatch(line); matches != nil {
			macPrefix := matches[1]
			organisation = matches[2]

			err := db.InsertOUI(macPrefix, organisation)
			if err != nil {
				fmt.Printf("%d) Error inserting this record => Organisation: %s Assignment Base16: %s\n", index, organisation, macPrefix)
				fmt.Println("Error inserting OUI:", err)
			} else {
				fmt.Printf("Inserted %d records successfully\n", index)
			}

			index++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
