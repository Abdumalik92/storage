package main

import (
	"fmt"
	"github.com/abdumalik92/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("f.txt", []byte("Hello world"))
	if err != nil {
		log.Fatal(err)
	}
	restoredFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(st)
	fmt.Println("It file restored ", restoredFile)
}
