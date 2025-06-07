package main

import (
	"fmt"
	"my-redis/internal"
)

func main() {
	db := store.NewDatabase()
	db.Set("Key1", "Value1")
	db.Set("Key2", "Value2")


	// Persist the database to a file
	err := db.Persist("mydb.gob")
	if err != nil {
		fmt.Println("Error persisting database:", err)
		return
	}

	// Load the database from the file
	err = db.Load("mydb.gob")
	if err != nil {
		fmt.Println("Error loading database:", err)
		return
	}

	// Retrieve values from the database
	value1, exists1 := db.Get("Key1")
	fmt.Println("Key1:", value1, "Exists:", exists1)
	value2, exists2 := db.Get("Key2")
	fmt.Println("Key2:", value2, "Exists:", exists2)
}