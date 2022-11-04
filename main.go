package main

import (
	"cryptostocksprices/repo"
	"fmt"
)

func main() {
	repo := repo.NewRepo()
	fmt.Println(repo.Method())
}
