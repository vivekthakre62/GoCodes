package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Items struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// func (i *Items) setData(id int, name string, quantity int, price float64) {
// 	i.id = id
// 	i.name = name
// 	i.quantity = quantity
// 	i.price = price
// }

type File struct {
	items map[int]Items
}

func (f *File) loadFromFile(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &f.items)
	if err != nil {
		panic(err)
	}
}
func (f *File) fileWrite(fileName string) string {
	sourceFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	bytes, err := json.Marshal(f.items) //Marshal for convert map into json formate

	if err != nil {
		panic(err)
	}
	_ = sourceFile
	_ = bytes
	_, e2 := sourceFile.Write(bytes)
	if e2 != nil {
		panic(e2)
	}
	return "Data added"
}

func (f *File) addItem(items Items) {
	f.items[items.ID] = items
	s := f.fileWrite("file.json")
	fmt.Println("Item added :", s)
}

func (f *File) removeItem(id int) {
	if _, exists := f.items[id]; !exists {
		fmt.Println("Data Not Found!")
		return
	}
	delete(f.items, id)

	b, err := json.MarshalIndent(f.items, " ", " ")
	if err != nil {
		panic(err)
	}
	err2 := os.WriteFile("file.json", b, 0644)
	if err2 != nil {
		fmt.Println("Error saving:", err)
		return
	}

	fmt.Println("Item removed successfully")

}

func (f *File) findItemById(id int) (Items, bool) {
	items, exits := f.items[id]
	return items, exits
}
func (f *File) updateItem(id int, name string, Quantity int, price float64) {
	item, exist := f.findItemById(id)
	if !exist {
		fmt.Println("Item not found for update")
		return
	}
	item.ID = id
	item.Name = name
	item.Quantity = Quantity
	item.Price = price
	f.items[id] = item
	data, err := json.MarshalIndent(f.items, "", " ")
	if err != nil {
		panic(err)
	}

	err1 := os.WriteFile("file.json", data, 0644)
	if err1 != nil {
		fmt.Println("File Does Not exist")
		return
	}
	fmt.Print("data updated")

}

// func main() {
// 	// file := File{
// 	// 	items: make(map[int]Items),
// 	// }
// 	// file.loadFromFile("file.json")
// 	// items := Items{
// 	// 	ID:       1,
// 	// 	Name:     "Mouse",
// 	// 	Quantity: 100,
// 	// 	Price:    300,
// 	// }

// 	// file.addItem(items)
// 	// item, exist := file.findItemById(5)
// 	// if !exist {
// 	// 	fmt.Println("Item does not exist")
// 	// } else {
// 	// 	fmt.Print(item)
// 	// }
// 	// file.removeItem(5)
// 	// file.updateItem(1, "Bottle", 12, 99)

//      for{

// 	 }
// }

func main() {
	file := File{
		items: make(map[int]Items),
	}
	file.loadFromFile("file.json")
	for {
		fmt.Println("\n===== Inventory Menu =====")
		fmt.Println("1. Add Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. Update Item")
		fmt.Println("4. Find Item By ID")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var id, quantity int
			var name string
			var price float64

			fmt.Print("Enter ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter Name: ")
			fmt.Scan(&name)

			fmt.Print("Enter Quantity: ")
			fmt.Scan(&quantity)

			fmt.Print("Enter Price: ")
			fmt.Scan(&price)

			item := Items{
				ID:       id,
				Name:     name,
				Quantity: quantity,
				Price:    price,
			}

			file.addItem(item)

		case 2:
			var id int
			fmt.Print("Enter ID to remove: ")
			fmt.Scan(&id)

			file.removeItem(id)

		case 3:
			var id, quantity int
			var name string
			var price float64

			fmt.Print("Enter ID to update: ")
			fmt.Scan(&id)

			fmt.Print("Enter New Name: ")
			fmt.Scan(&name)

			fmt.Print("Enter New Quantity: ")
			fmt.Scan(&quantity)

			fmt.Print("Enter New Price: ")
			fmt.Scan(&price)

			file.updateItem(id, name, quantity, price)

		case 4:
			var id int
			fmt.Print("Enter ID to search: ")
			fmt.Scan(&id)

			item, found := file.findItemById(id)
			if !found {
				fmt.Println("Item not found")
			} else {
				fmt.Printf("Item Found: %+v\n", item)
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, try again")
		}
	}
}
