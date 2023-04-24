package main

import "fmt"

// Iterate using a callback function
func main() {
	// Callback approach is push based Iterator approach.
	// use IterateBooks to iterate via a callback function
	// lib.IterateBooks(myBookCallback)

	// Use IterateBooks to iterate via anonymous function
	// lib.IterateBooks(func(b Book) error {
	// 	fmt.Println("Book Author:", b.Author)
	// 	return nil
	// })

	// Pull based Iterator approach.
	// TODO: create a BookIterator
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
