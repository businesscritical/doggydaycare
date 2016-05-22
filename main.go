package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := flag.String("name", "", "")
	flag.Parse()

	http.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
		resp.Write([]byte(fmt.Sprintf("Hello from the %s process!", *name)))
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
