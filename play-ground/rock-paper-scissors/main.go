// package main

// import (
// 	"log"
// 	"net/http"
// 	"text/template"
// )

// func main() {
// 	http.HandleFunc("/home", homePage)
// 	log.Println("Listening on port 8183...")
// 	http.ListenAndServe(":8183", nil)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "index.html")
// }

// func renderTemplate(w http.ResponseWriter, page string) {
// 	t, err := template.ParseFiles(page)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	err = t.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", handleHome)
	fmt.Println("Listening on port 8082...")
	http.ListenAndServe(":8082", nil)

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	handleTemplate(w, "index.html")
}

func handleTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
