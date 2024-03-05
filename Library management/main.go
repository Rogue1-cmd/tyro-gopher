package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Book struct
type Book struct {
	ID            int
	Title, Author string
	IsTaken       bool
}

// Library Struct
type Library struct {
	books []Book
}

//Library Methods

// Adding a book
func (l *Library) Add(title string, author string) {
	//new book instance
	book := Book{
		ID:      len(l.books) + 1,
		Title:   title,
		Author:  author,
		IsTaken: false,
	}
	l.books = append(l.books, book)
}

// Listing books
func (l *Library) List() {
	for _, book := range l.books {
		status := "Available"
		if book.IsTaken {
			status = "Borrowed"
		}
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, status)
		//fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, status)
	}
}

// Borrowing
func (l *Library) BorrowBook(id int) Book {
	borrowedBook := Book{}
	for index, book := range l.books {
		if book.ID == id {
			l.books[index].IsTaken = true
			borrowedBook = l.books[index]
			break
		}
	}

	return borrowedBook
}

// Return Book
func (l *Library) ReturnBook(id int) {
	for index, book := range l.books {
		if book.ID == id {
			l.books[index].IsTaken = false
			break
		}
	}
	//return Book
}

// Delete
func (l *Library) DeleteBook(id int) {
	books := l.books
	for i, book := range l.books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			l.books = books
			return
		}
	}
}

var id int
var title, author string

func addingBook(lib Library) Library {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nEnter Book Title: ")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("\nEnter Author's Name: ")
	author, _ = reader.ReadString('\n')
	author = strings.TrimSpace(author)

	lib.Add(title, author)
	fmt.Println("\nAdded ", title, "by ", author, "to library")
	lib.List()

	return lib
}

func returning(lib Library) Library {
	fmt.Println("\n\tWhich book are you returning?")
	lib.List()

	fmt.Println("\n\t Enter Book ID: ")
	fmt.Scanln(&id)

	lib.ReturnBook(id)
	time.Sleep(1 * time.Second)

	fmt.Println("\n\t Updated List ")
	lib.List()

	return lib
}

func borrow(lib Library) (Library, error) {
	fmt.Println("\n\tBorrowing Book\nSelect from one of the following")
	lib.List()

	fmt.Println("\nEnter ID: ")
	//fmt.Scanln(&id)

	_, err := fmt.Scanln(&id)
	if err != nil || id <= 0 || id > len(lib.books) {
		fmt.Println("Invalid ID. Please enter a valid ID.")
		return lib, err
	}

	if lib.books[id-1].IsTaken {
		fmt.Println("The Book is unavailable")
	} else {
		borrowedBook := lib.BorrowBook(id)
		fmt.Println("You have borrowed: ", borrowedBook)
		time.Sleep(500 * time.Millisecond)

		fmt.Println("\nPress Enter to Continue...")
		fmt.Scanln()
	}

	return lib, nil
}

func slow() {
	fmt.Println("\nPress Enter to Continue ")
	fmt.Scanln()
	time.Sleep(1 * time.Second)

}

func populateLib() Library {
	lib := Library{}
	lib.Add("A Tale of Two Cities", "Charles Dickens")
	lib.Add("The Little Prince (Le Petit Prince)", "Antoine de Saint-Exupéry")
	lib.Add("Harry Potter and the Philosopher's Stone", "J. K. Rowling")
	lib.Add("And Then There Were None", "Agatha Christie")
	lib.Add("Dream of the Red Chamber (紅樓夢)", "Cao Xueqin")

	return lib

}

func main() {

	var choice int

	lib := Library{}
	lib = populateLib()

	fmt.Println("\n\t Welcome to the Library")
	for {
		fmt.Println("\n\t Select what you would like to do? ")
		fmt.Println("\n1. Add Book\n2. Borrow Book\n3. Return Book\n4. List Books \n5. Delete Book Entry \n6. Exit Library ")
		fmt.Scanln(&choice)

		if choice <= 0 || choice > 6 {
			fmt.Println("\n\tEnter choice between 1 - 6!!")
			continue
		}

		if choice == 6 {
			break
		}

		switch choice {
		case 1: //Adding a new book
			lib = addingBook(lib)
			slow()

		case 2: //Borrowing A book
			lib, err := borrow(lib)
			if err != nil {
				fmt.Println("Error Borrowing Book: ", err.Error())
			}
			slow()

			fmt.Println("\n\tUpdated List ")
			lib.List()

		case 3: //Returning a book
			lib = returning(lib)
			slow()
			fmt.Println("\n\tUpdated List ")
			lib.List()

		case 4:
			fmt.Println("\n\t Books: ")
			lib.List()

			slow()

		case 5:
			var id int
			fmt.Println("\n\t Enter Book ID to Delete: ")
			fmt.Scanln(&id)

			lib.DeleteBook(id)

			//updated book status
			fmt.Println("\n\tUpdated List ")
			lib.List()

		}

	}

}
