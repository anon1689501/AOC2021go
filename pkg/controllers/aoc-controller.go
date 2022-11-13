package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Day1(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	aocInput := r.FormValue("aocInput")
	aocInputString := strings.Fields(aocInput)

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

// fmt.Println(total)
