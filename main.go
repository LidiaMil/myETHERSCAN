package main

import ("fmt"
		"database/sql"
		_ "github.com/go-sql-driver/mysql"

)

func main()  {
	fmt.Println("работа с mysql")
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//add data

	insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Alex', 16)")
	if err != nil {
		panic(err)
	}

	defer insert.Close()
	fmt.Println("подключено к mysql")

}