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
	res := pfn.Run(r.FormValue("code"), false).Output
	wr.Write([]byte(res))
}
