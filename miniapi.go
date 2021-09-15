package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	tm := time.Now()
	h := tm.Hour()
	m := tm.Minute()
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "%dh%d", h, m)
	}
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		fmt.Fprintf(w, "%s: %s",req.FormValue("author"), req.FormValue("entry"))
		createFile(req.FormValue("entry"), w)
	}
}

func entriesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if _, err := os.Stat("result.txt"); os.IsNotExist(err){
			fmt.Printf("The file don't exist")
		}

		readFile, err := os.ReadFile("result.txt")
		if err != nil {
			fmt.Fprintf(w,"There is an issue on the file")
		}
		fmt.Fprintf(w, "%s", string(readFile))
	}
}

func createFile(entry string, w http.ResponseWriter){
	if _, err := os.Stat("result.txt"); os.IsNotExist(err){
		d1 := []byte("")
		os.WriteFile("result.txt", d1, 0644)
	}

	saveFile, err := os.OpenFile("./result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("There is an issue on the file")
		fmt.Fprintf(w, "There is an issue on the file")
	}
	defer saveFile.Close()
	saveFile.WriteString(entry + "\n")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/entries", entriesHandler)
	http.ListenAndServe(":4567", nil)
}
