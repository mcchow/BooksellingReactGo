# BooksellingReactGo

<b>Setup the backend:</b><br>
make sure you have npm install<br>
make sure you have mysql install and setup a empty database<br>
change this line in main.go with your username and password of the db, if you are using root account change <username> to root:<br>
mydb, err := sql.Open("mysql",
		"<username>:<password>@tcp(127.0.0.1:3306)/gravity_books")<br>
run all sql code in the DBSetup file<br>
run npm run server and if you see database connect and host:8000 and it should working<br>

<b>Frontend:</b><br>
Run npm start for frontend<br>

Reference<br>
https://hugo-johnsson.medium.com/rest-api-with-golang-mux-mysql-c5915347fa5b<br>
https://stackoverflow.com/questions/10070508/sqlite-like-order-by-match-query<br>