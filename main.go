package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Owner  string `json:"owner"`
}

var db *sql.DB

func main() {
	mydb, err := sql.Open("mysql",
		"<username>:<password>@tcp(127.0.0.1:3306)/gravity_books")
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	} else {
		fmt.Println("Database connected")
	}
	db = mydb
	//close db when api ends
	defer mydb.Close()

	router := mux.NewRouter()

	router.HandleFunc("/getbook/{name}/{owner}", getbook).Methods("GET")
	router.HandleFunc("/posts/{name}", getPosts).Methods("GET")
	router.HandleFunc("/add/{name}/{author}/{ISBN}/{owner}", addbook).Methods("POST")
	fmt.Println("hosting on: http://localhost:8000")
	http.ListenAndServe(":8000", router)
}

//  from /posts/{name}
//  get book by title, panic and return nothing if there is any error
func getPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var books []Book
	var name = params["name"]
	fmt.Println("getPosts call with" + name)
	var name1 = "%" + name + "%"
	var name2 = name + "%"
	result, err := db.Query("SELECT book.book_id , book.title, author_name , book.isbn13 ,store_name FROM book INNER JOIN book_author on book_author.book_id = book.book_id INNER JOIN author on author.author_id = book_author.author_id INNER JOIN store_book on store_book.book_id = book.book_id INNER JOIN store on store.store_id = store_book.store_id WHERE book.title LIKE ? ORDER BY (CASE WHEN book.title = ? THEN 1 WHEN book.title LIKE ? THEN 2 ELSE 3 END),book.title;", name1, name, name2)
	if err != nil {
		fmt.Println("search error")
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var book Book
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Owner)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}
	booksJSON, _ := json.Marshal(books)
	w.WriteHeader(200)
	w.Write([]byte(booksJSON))
}

// from /getbook/{name}/{owner}
// return books with like name and like owner
func getbook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var books []Book
	var name = params["name"]
	var owner = params["owner"]
	fmt.Println("get book and owner call with" + name + " " + owner)
	var name1 = "%" + name + "%"
	var name2 = name + "%"
	var owner1 = "%" + owner + "%"
	result, err := db.Query("SELECT book.book_id , book.title, author_name , book.isbn13 ,store_name FROM book INNER JOIN book_author on book_author.book_id = book.book_id INNER JOIN author on author.author_id = book_author.author_id INNER JOIN store_book on store_book.book_id = book.book_id INNER JOIN store on store.store_id = store_book.store_id WHERE book.title LIKE ? AND store.store_name like ? ORDER BY (CASE WHEN book.title = ? THEN 1 WHEN book.title LIKE ? THEN 2 ELSE 3 END),book.title;", name1, owner1, name, name2)
	if err != nil {
		fmt.Println("search error")
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var book Book
		err := result.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Owner)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
	}
	booksJSON, _ := json.Marshal(books)
	w.WriteHeader(200)
	w.Write([]byte(booksJSON))
}

// from /add/{name}/{author}/{ISBN}/{owner}
// add book if book not exist
// add book and owner pair if not exist
func addbook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var name = params["name"]
	var author = params["author"]
	var isbn = params["ISBN"]
	var owner = params["owner"]
	var authorid = 0
	var bookid = 0
	var ownerid = 0
	var message = "sucess"
	fmt.Println("addbook call with" + name)
	authorresult, err := db.Query("SELECT author_id FROM gravity_books.author WHERE author_name = ?", author)
	if err != nil {
		fmt.Println("search error")
		panic(err.Error())
	}
	if authorresult.Next() {
		err := authorresult.Scan(&authorid)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		fmt.Println("author found")
	} else { // add author if not exist
		authorq, err := db.Prepare("INSERT INTO author (author_name) VALUES(?)")
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		addauthorresult, err := authorq.Exec(author)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		temp, err := addauthorresult.LastInsertId()
		authorid = int(temp)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
	}
	fmt.Println("author id is:")
	fmt.Println(authorid)
	defer authorresult.Close()

	//add book
	//as ISBN is unique, we use it as identifier to see if the book exist in our db
	bookresult, err := db.Query("SELECT book.book_id FROM gravity_books.book WHERE isbn13 = ?", isbn)
	if err != nil {
		fmt.Println("search error")
		panic(err.Error())
	}
	if bookresult.Next() {
		err := bookresult.Scan(&bookid)
		fmt.Println("book name found")
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
	} else { // add book if not exist
		addbook, err := db.Prepare("INSERT INTO book (title, isbn13) VALUES(?, ?)")
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		addbookresult, err := addbook.Exec(name, isbn)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		temp, err := addbookresult.LastInsertId()
		bookid = int(temp)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
	}
	defer bookresult.Close()
	fmt.Println("book id is:")
	fmt.Println(bookid)

	bookauthorresult, err := db.Query("SELECT * FROM gravity_books.book_author WHERE book_id = ? AND author_id = ?", bookid, authorid)
	if err != nil {
		fmt.Println("error")
	}
	if !bookauthorresult.Next() {
		addpair, err := db.Prepare("INSERT INTO book_author (book_id, author_id) VALUES(?, ?)")
		fmt.Println("added book")
		if err != nil {
			fmt.Println("Pair exist")
		}
		res, err := addpair.Exec(bookid, authorid)
		if err != nil {
			fmt.Println("book_author_Scan error")
			panic(err.Error())
		}
		fmt.Println(res)
	} else {
		message = "book exist"
	}
	defer bookauthorresult.Close()

	//book owner relation

	//if owner name is unqiue and exist
	ownerresult, err := db.Query("SELECT store_id FROM store WHERE store_name = ?", owner)
	if err != nil {
		fmt.Println("search error")
		panic(err.Error())
	}
	if ownerresult.Next() {
		err := ownerresult.Scan(&ownerid)
		fmt.Println("store name found")
		if err != nil {
			fmt.Println("store_Scan error")
			panic(err.Error())
		}
	} else { // add owner if not exist
		addowner, err := db.Prepare("INSERT INTO `store` (`store_name`)VALUES (?)")
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		addownerresult, err := addowner.Exec(owner)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
		temp, err := addownerresult.LastInsertId()
		ownerid = int(temp)
		if err != nil {
			fmt.Println("Scan error")
			panic(err.Error())
		}
	}

	bookownerresult, err := db.Query("SELECT * FROM store_book WHERE book_id = ? AND store_id = ?", bookid, ownerid)
	if err != nil {
		fmt.Println("error")
	}
	if !bookownerresult.Next() {
		addpair, err := db.Prepare("INSERT INTO store_book (book_id, store_id) VALUES(?, ?)")
		fmt.Println("added book_store")
		if err != nil {
			fmt.Println("Pair exist")
		}
		res, err := addpair.Exec(bookid, ownerid)
		if err != nil {
			fmt.Println("Pair exist")
			panic(err.Error())
		}
		fmt.Println(res)
	} else {
		message = "Owner and book pair exist"
	}
	defer bookauthorresult.Close()

	messageJSON, _ := json.Marshal(message)
	w.WriteHeader(200)
	w.Write([]byte(messageJSON))
}

/*
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":80", r)
}*/
