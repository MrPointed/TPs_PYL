package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"main/data"
	"os"
	"strconv"
	"strings"
)

func main() {

	var items []data.Item //array of item

	file, err := os.OpenFile("items.json", os.O_RDWR|os.O_CREATE, 0666) //Open the file called items, if it doesn't exist then create it
	if err != nil {
		panic(err)
	}

	defer file.Close() // close the file to avoid memory trash

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 { //check if there's content within the file
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &items)
		if err != nil {
			panic(err)
		}
	} else { //if not
		items = []data.Item{}
	}

	if len(os.Args) < 2 {
		fmt.Println(showUsage())
	} else {
		switch os.Args[1] {
		case "list":
			data.List(items)
		case "add":
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			reader = bufio.NewReader(os.Stdin)
			fmt.Println("Amount: ")
			amount, _ := reader.ReadString('\n')
			amount = strings.TrimSpace(amount)

			amnt, _ := strconv.Atoi(amount) //I've added this to convert string to integer

			if amnt <= 0 {
				fmt.Println("Please enter a correct value")
				return
			}
			items = data.Add(items, name, amnt)

			data.UpdateFile(file, items)

			fmt.Println("New item added")

		case "remove":
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("ID:")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			idd, _ := strconv.Atoi(id) //I've added this to convert string to integer

			items = data.Remove(idd, items)
			data.UpdateFile(file, items)
		}
	}

}

func showUsage() string {
	return "list of commands (list|add|remove)"
}
