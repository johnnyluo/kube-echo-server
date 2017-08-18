package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "http listen port")
	flag.Parse()
	http.HandleFunc("/", handler)
	fmt.Printf("Listen on port %d\n", *port)
	if err := http.ListenAndServe(":"+strconv.Itoa(*port), nil); nil != err {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Printf("request:%s\n", r.URL.String())
}
