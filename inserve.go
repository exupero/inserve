package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var contentType = flag.String("content-type", "text/plain", "Content-Type to serve input as")

func main() {
	flag.Parse()

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Input:", err)
	}
	content := string(b)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", *contentType)
		io.WriteString(w, content)
	})

	log.Println("Starting server...")
	if err = http.ListenAndServe("0.0.0.0:" + flag.Arg(0), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
