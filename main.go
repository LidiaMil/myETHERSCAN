package main

import ("fmt"
		"net/http"
		)

func home_page(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello YOU!!!")
}

func main()  {
	fmt.Println("hello you")
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
}