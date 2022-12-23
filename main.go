package main

import ("fmt"
		"net/http"
		)

type User struct{
	name string
	age uint16
	money int16
	avg_grades, happiness float32
}

func home_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello YOU!!!")
}

func contacts_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "contacts page !!!")
}

func handleRequest()  {
	fmt.Println("hello you")
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main()  {
	handleRequest()
}