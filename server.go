package main

import (
	"html/template"
	"log"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Quote string
	Hobby string
}

func Home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./templates/header.html", "./templates/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func IdCard(w http.ResponseWriter, r *http.Request, info []Student) {
	template, err := template.ParseFiles("./idCard.html", "./templates/header.html", "./templates/footer.html", "./templates/student.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, info)
}
func main() {
	var tabInfos []Student
	tabInfos = append(tabInfos, Student{"Julien", 23, "ok", "fais pas chier"})
	tabInfos = append(tabInfos, Student{"David", 17, "tg", "nul à valo"})
	tabInfos = append(tabInfos, Student{"Joseph", 22, "oui", "stuck plat sur valo"})
	http.HandleFunc("/", Home)
	http.HandleFunc("/idcard", func(w http.ResponseWriter, r *http.Request) {
		IdCard(w, r, tabInfos)
	})
	fs := http.FileServer(http.Dir("/css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.ListenAndServe(":8080", nil)
}
