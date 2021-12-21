# BooksellingReactGo

<b>Setup the backend:</b><br>
<ol>
<li>make sure you have npm install<br></li>
<li>make sure you have mysql install and setup a empty database<br></li>
<li>change this line in main.go with your username and password of the db, if you are using root account change <username> to root:<br>
mydb, err := sql.Open("mysql",
		"<username>:<password>@tcp(127.0.0.1:3306)/gravity_books")<br></li>
<li>run all sql code in the DBSetup file<br></li>
<li>run <b>npm run server</b> and if you see database connect and host:8000 and it should working<br></li>
</ol>
<b>Frontend:</b><br>
Run <b>npm start</b> for frontend<br>

Reference<br>
https://hugo-johnsson.medium.com/rest-api-with-golang-mux-mysql-c5915347fa5b<br>
https://stackoverflow.com/questions/10070508/sqlite-like-order-by-match-query<br>