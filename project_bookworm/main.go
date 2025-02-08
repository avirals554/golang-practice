package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func loaddata(filepath string) ([]Bookworm, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err

	}

	defer f.Close()
	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

func main() {
	//  filepath string :="(home/aviral/learning_go/project_bookworm/testdata/bookworm.json)"
	bookworms, err := loaddata("./testdata/bookworm.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(bookworms)

}
