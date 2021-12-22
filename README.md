# BooksellingReactGo

This is a go backend and react frontend project. The goal is to build a simple app that handle:<br>
The application should meet the following criteria via API:

## A store owner can upload their inventory data to the application<br>
we have a add book function that add a book to the DB
## A store owner can view their inventory data in a browser<br>
we can see the by search book button on the react page
## A store owner can search their inventory data<br>
user can search by keyword title using the search book button
## (Bonus) A store owner can search their inventory data using multiple fields at the same time<br>
user can search by keyword title and owner using the search book and owner button
## (Bonus) The application can support multiple store owners at the same time<br>
different user can upload their data using their name/id


# what is looking for in this code:<br>

1.     Does it work

2.     Clean, well-organized, easy to read code

3.     Well thought-out architectural decisions that scale and are performant

4.     Bonus points for writing tests

Here is my Design:<br>
This is a book storage that store different books record for each owner<br>
user can search the system and see who own the book and what book the system have<br>

DB design:<br>
![Capture](https://user-images.githubusercontent.com/44908668/147015232-5c0fee89-cd03-4413-8d5c-d583f6588d29.PNG)

<b>book:</b> a list of all books available in the System. As different people may have the same book, we make it one table<br>
<b>author:</b> a list of all authors. As books may have same author, we make a table to reduce storage space by FK<br>
<b>book_author:</b> stores the authors for each book, which is a many-to-many relationship.<br>
<b>book_language:</b> a list of possible languages of books.(exist but not use in current development)<br>
<b>publisher:</b> a list of publishers for books.(exist but not use in current development)<br>
<b>store:</b> a list of the owners of the Gravity Bookstore. Password can add if we want user verification later. For now password function have not implement yet<br>
<b>store_book:</b> stores the owner for each book, which is a many-to-many relationship.<br>


<b>Setup the backend:</b><br>
<ol>
<li>make sure you have npm install<br></li>
<li>make sure you have npm install<br></li>
<li>make sure you have go install<br></li>
<li>make sure you have mysql install and setup a empty database<br></li>
<li>change this line in main.go with your username and password of the db, if you are using root account change username to root:<br>
mydb, err := sql.Open("mysql",
		"(username):(password)@tcp(127.0.0.1:3306)/gravity_books")<br></li>
<li>run all sql code in the DBSetup file<br></li>
<li>run <b>npm run server</b> and if you see database connect and host:8000 and it should working<br></li>
</ol>
<b>Frontend:</b><br>
Run <b>npm start</b> for frontend<br>

<b>Testing:</b><br>
Frontend:<br>
input validation: ISBN cannot longer than 13 and other cannot longer than 100<br>

Backend:<br>
I use postman to test my API. I skip the security stuff(SQL injection) as it is not going to open to public. Here is what I test:<br>
Insert and search should work: <br>
GET http://localhost:8000/posts/test<br>
GET http://localhost:8000/getbook/test/test<br>
POST http://localhost:8000/add/test/test/test/test<br>






Reference<br>
https://hugo-johnsson.medium.com/rest-api-with-golang-mux-mysql-c5915347fa5b<br>
https://stackoverflow.com/questions/10070508/sqlite-like-order-by-match-query<br>
