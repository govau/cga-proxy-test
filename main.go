package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	log.Fatal(http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bb, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(bb)
		fmt.Fprintf(w, "\n\nYou are logged in as: %s\n", r.Header.Get("X-Forwarded-User"))
	})))
}
