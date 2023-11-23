package main 

import "fmt"
import "time"

type Book struct{
  title  string
  author string
  numPages int 

  isSaved bool
  savedAt time.Time

}



func saveBook (book *Book){
  book.isSaved = true
  book.savedAt = time.Now()
  fmt.Println("saved")
}



func main(){
  book := Book{
		title:    "The Go Programming Language",
		author:   "Alan A. A. Donovan and Brian W. Kernighan",
		numPages: 380,
	}

	fmt.Println("The book:", book)
	saveBook(&book)
 }
