package main

import (
	"net/http"
	"os"

	"github.com/TsukiGva2/pfn"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/execute", runpfn)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

func runpfn(wr http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	noprelude := false

	np, exists := query["noprelude"]
	if exists && np[0] == "true" {
		noprelude = true
	}

	res := pfn.Run(r.FormValue("code"), false, noprelude).Output
	wr.Write([]byte(res))
}
