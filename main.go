package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Record struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

const fileName = "records.json"

func loadRecords() []Record {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Record{}
	}
	var records []Record
	json.Unmarshal(data, &records)
	return records
}

func saveRecords(records []Record) {
	data, _ := json.MarshalIndent(records, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func nextID(records []Record) int {
	max := 0
	for _, r := range records {
		if r.ID > max {
			max = r.ID
		}
	}
	return max + 1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tu [create|change|delete|list]")
		return
	}

	command := os.Args[1]
	records := loadRecords()

	switch command {

	case "create":
		if len(os.Args) < 3 {
			fmt.Println("Provide text")
			return
		}
		record := Record{
			ID:    nextID(records),
			Title: os.Args[2],
		}
		records = append(records, record)
		saveRecords(records)
		fmt.Println("Created")

	case "change":
		if len(os.Args) < 4 {
			fmt.Println("Usage: tu change <id> \"new text\"")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		newText := os.Args[3]

		for i, r := range records {
			if r.ID == id {
				records[i].Title = newText
				saveRecords(records)
				fmt.Println("Changed")
				return
			}
		}
		fmt.Println("Not found")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: tu delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		newRecords := []Record{}
		found := false

		for _, r := range records {
			if r.ID != id {
				newRecords = append(newRecords, r)
			} else {
				found = true
			}
		}

		if !found {
			fmt.Println("Not found")
			return
		}

		saveRecords(newRecords)
		fmt.Println("Deleted")

	case "list":
		for _, r := range records {
			fmt.Printf("[%d] %s\n", r.ID, r.Title)
		}

	default:
		fmt.Println("Unknown command")
	}
}