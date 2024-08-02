package controllers

import (
    "bufio"
    "fmt"
    "library_management/models"
    "library_management/services"
    "os"
    "strconv"
    "strings"
)

var library = services.NewLibrary()

func displayMenu() {
    fmt.Println("Library Management System")
    fmt.Println("0. Add Member")
    fmt.Println("1. Add Book")
    fmt.Println("2. Remove Book")
    fmt.Println("3. Borrow Book")
    fmt.Println("4. Return Book")
    fmt.Println("5. List Available Books")
    fmt.Println("6. List Borrowed Books")
    fmt.Println("7. Exit")
}

func getInput(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

func addMember() {
	id, _ := strconv.Atoi(getInput("Enter Member ID: "))
	name := getInput("Enter Member Name: ")
	member := models.Member{
		ID:   id,
		Name: name,
	}
	
	library.AddMember(member)
	fmt.Println("Member added successfully!")
}

func addBook() {
    id, _ := strconv.Atoi(getInput("Enter Book ID: "))
    title := getInput("Enter Book Title: ")
    author := getInput("Enter Book Author: ")

    book := models.Book{
        ID:     id,
        Title:  title,
        Author: author,
        Status: "Available",
    }
    library.AddBook(book)
    fmt.Println("Book added successfully!")
}

func removeBook() {
    id, _ := strconv.Atoi(getInput("Enter Book ID to Remove: "))
    library.RemoveBook(id)
    fmt.Println("Book removed successfully!")
}

func borrowBook() {
    bookID, _ := strconv.Atoi(getInput("Enter Book ID to Borrow: "))
    memberID, _ := strconv.Atoi(getInput("Enter Member ID: "))
    err := library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book borrowed successfully!")
    }
}

func returnBook() {
    bookID, _ := strconv.Atoi(getInput("Enter Book ID to Return: "))
    memberID, _ := strconv.Atoi(getInput("Enter Member ID: "))
    err := library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book returned successfully!")
    }
}

func listAvailableBooks() {
    books := library.ListAvailableBooks()
    fmt.Println("Available Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}

func listBorrowedBooks() {
    memberID, _ := strconv.Atoi(getInput("Enter Member ID: "))
    books := library.ListBorrowedBooks(memberID)
    fmt.Println("Borrowed Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}

func Run() {
    for {
        displayMenu()
        choice, _ := strconv.Atoi(getInput("Enter your choice: "))
        switch choice {
        case 0:
        	addMember()
        case 1:
            addBook()
        case 2:
            removeBook()
        case 3:
            borrowBook()
        case 4:
            returnBook()
        case 5:
            listAvailableBooks()
        case 6:
            listBorrowedBooks()
        case 7:
            fmt.Println("Exiting the system. Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
