package main

import (
	"fmt"
	"log"
	"storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("f.txt", []byte("Hello world"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("It's work", file)
}
