package test

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func TestStructV1(){
	var book1 Books
	var book2 Books

	book1.author = "大大大大"
	book1.title = "上线"
	book1.subject = "dda大大大"
	book1.book_id = 3

	book2.author = "大大大大2"
	book2.title = "上线2"
	book2.subject = "dda大大大2"
	book2.book_id = 4

	  /* 打印 Book1 信息 */
   fmt.Printf( "Book 1 title : %s\n", book1.title)
   fmt.Printf( "Book 1 author : %s\n", book1.author)
   fmt.Printf( "Book 1 subject : %s\n", book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", book1.book_id)

   /* 打印 Book2 信息 */
   fmt.Printf( "Book 2 title : %s\n", book2.title)
   fmt.Printf( "Book 2 author : %s\n", book2.author)
   fmt.Printf( "Book 2 subject : %s\n", book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", book2.book_id)
}
