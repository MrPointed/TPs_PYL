package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	ID     int    `json:"ID"`
	Name   string `json:"Name"`
	Amount int    `json:"Amount"`
}

func List(items []Item) {

	if len(items) == 0 {
		fmt.Println("No items added. To add a new item use the command (add)")
		return
	}

	for _, item := range items {
		fmt.Printf("\nID: %d\nName: %s\nAmount: %d\n", item.ID, item.Name, item.Amount)
	}
}

func Add(items []Item, name string, amount int) []Item {

	newItem := Item{
		ID:     GetNextID(items),
		Name:   name,
		Amount: amount,
	}

	return append(items, newItem)
}

func UpdateFile(f *os.File, items []Item) {
	//convert to json-encoded bytes
	bytes, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = f.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(f)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush() //por las dudas queda algo en el buffer
	if err != nil {
		panic(err)
	}
}

// get the next available id for the next task
func GetNextID(items []Item) int {
	if len(items) == 0 {
		return 1
	}
	return items[len(items)-1].ID + 1
}
