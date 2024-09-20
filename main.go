package main

import (
	"log"
	"os"

	"network-scan/database"
	"network-scan/oui"
	"network-scan/utils"
)

const dbPath string = "./oui.db"

func init() {
	err := database.CreateDbIfNotExists(dbPath)
	if err != nil {
		log.Fatalf("Error creating database file: %v", err)
	}

	db, err := database.New(dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	err = db.CreateOuiTable()
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func main() {
	db, err := database.New(dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	} else {
		utils.PrintInColor("Database opened successfully", 28)
	}
	defer db.Close()

	file, err := os.Open("standards-oui.ieee.org.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	oui.ScanDataFromTextFile("/Users/brijesh/projects/ongoing/network-scan/standards-oui.ieee.org.txt", db)
	utils.PrintInColor("OUI data imported successfully", 28)
}
