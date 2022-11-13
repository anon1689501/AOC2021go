package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type TodoPageData struct {
	Cookie string
}

func Day1(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	aocInput := r.FormValue("aocInput")
	aocInputString := strings.Fields(aocInput)
	cookie := http.Cookie{
		Name:  "inputDay1",
		Value: "abc",

		MaxAge: 3600,

		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	log.Println("cookie set")

	total := 0

	for _, v := range aocInputString {
		myNum, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(w, "int conversion error")
		}
		total = total + day1_fuel(myNum)

	}
	params := mux.Vars(r)
	fmt.Fprintf(w, "\n"+params["id"])
	fmt.Fprintf(w, "\n%v", total)
	w.Write([]byte("cookie!"))

}

func day1_fuel(x int) int {
	x = (x / 3) - 2
	if x < 0 {
		return 0
	}
	return x + day1_fuel(x)

}

func GetForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/form.html")
}
func MakeIndex(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "static/index.html")
	cookie, _ := r.Cookie("inputDay1")
	var myCookie string
	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, http.ErrNoCookie):
	// 		// http.Error(w, "cookie not found", http.StatusBadRequest)
	// 		log.Println("no cookie")
	// 		return
	// 	default:
	// 		log.Println(err)
	// 		// http.Error(w, "server error", http.StatusInternalServerError)
	// 	}
	// 	return
	// }
	indexTemplate := template.Must(template.ParseFiles("static/index.html"))
	if cookie == nil {
		myCookie = "no cookie"
	} else {
		myCookie = cookie.Value
	}

	data := TodoPageData{
		Cookie: myCookie} //, err2 := r.Cookie("inputDay1")
	// if err2 != nil {
	// 	fmt.Fprintf(w, "cookie read error")
	// }
	indexTemplate.Execute(w, data)
}

// fmt.Println(total)
