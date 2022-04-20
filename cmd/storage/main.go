package main

import (
	"fmt"
	"storage2/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("It's work", st)
}
