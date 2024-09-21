package main

import (
	"fmt"
	"log"
	"os"

	"network-scan/arp"
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

	utils.ConfirmBeforeRunning("Do you want to seed sqlite from IEEE data?", func() {
		file, err := os.Open("standards-oui.ieee.org.txt")
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()

		oui.ScanDataFromTextFile("/Users/brijesh/projects/ongoing/network-scan/standards-oui.ieee.org.txt", db)
		utils.PrintInColor("OUI data imported successfully", 28)
	})

	utils.ConfirmBeforeRunning("Do you want to scan for devices in your network?", func() {
		devices, err := arp.GetDevices()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Devices found:")
		for _, device := range devices {
			fmt.Printf("IP: %s, MAC: %s\n", device.IP, device.MAC)
		}
	})

}
