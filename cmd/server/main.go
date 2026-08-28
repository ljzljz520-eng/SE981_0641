package main

import (
	"fmt"
	"warehouse5s/internal/store"
)

func main() {
	s, e := store.Open("warehouse.db")
	if e != nil {
		panic(e)
	}
	defer s.Close()
	fmt.Println("warehouse 5S service ready")
}
