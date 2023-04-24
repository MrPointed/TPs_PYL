package data

import "fmt"

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

// get the next available id for the next task
func GetNextID(items []Item) int {
	if len(items) == 0 {
		return 1
	}
	return items[len(items)-1].ID + 1
}
