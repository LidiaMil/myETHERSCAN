package main

import ("fmt"
		"net/http"
		)

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