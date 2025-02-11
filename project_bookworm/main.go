package main

import (
	"encoding/json"
	"fmt"
	"os"
  "sort"
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
func bookcount (bookworm []Bookworm) map[Book]uint{
  count:=make(map[Book]uint,451)
  for _,bookworm:=range bookworm {
    for _,book:=range bookworm.Books{
       count[book]++
    }

  }
  return count
}

func findCommonBooks(bookworms []Bookworm) []Book {
booksOnShelves := bookcount(bookworms) 
var commonBooks []Book
for book, count := range booksOnShelves { 
if count > 1 { 
commonBooks = append(commonBooks, book)
}
}
return commonBooks
}


func sortBooks(books []Book) []Book {
sort.Slice(books, func(i, j int) bool { 
if books[i].Author != books[j].Author {
return books[i].Author < books[j].Author
}
return books[i].Title < books[j].Title 
})
return books
}

func displayBooks(books []Book) {
for _, book := range books {
fmt.Println("-", book.Title, "by", book.Author)
}
}

func main() {
bookworms, err := loaddata("/home/aviral/learning_go/project_bookworm/testdata/bookworm.json")
if err != nil {
_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
os.Exit(1)
}
commonBooks := findCommonBooks(bookworms)
fmt.Println("Here are the books in common:")
displayBooks(commonBooks)
}
