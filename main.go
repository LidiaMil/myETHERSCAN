package main

import ("fmt"
		"net/http"
		"html/template"
		)

type User struct{
	Name string
	Age uint16
	Money int16
	Avg_grades, happiness float32
}

func (u User) getAllInfo() string{
	return fmt.Sprintf("User name is: %s. He is %d. He has money = %d", u.Name, u.Age, u.Money)

}

func (u *User) setNewName(newName string){
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request){
	bob := User{"Bob", 10, 100, 4.2, 1000}
	bob.setNewName("Alex")
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _:= template.ParseFiles("template/home_page.html")
	tmpl.Execute(w, bob)
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
	// var bob User =
	// bob := User{name: "Bob", age: 10, money: 100, avg_grades: 4.2, happiness: 1000}
	// bob := User{"Bob", 10, 100, 4.2, 1000}
	handleRequest()
}