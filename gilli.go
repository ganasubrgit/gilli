package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	listenFlag = flag.String("listen", ":8080", "address and port to listen")
	textFlag   = flag.String("text", "<h2>Hello, this is Gilli!</h2>", "text to put on the static webpage")
	fileFlag   = flag.String("html", "", "html filename")
)

func handler(w http.ResponseWriter, r *http.Request) {

	if len(*fileFlag) != 0 {
		b, err := ioutil.ReadFile(*fileFlag) // just pass the file name
		str := string(b)
		fmt.Fprintf(w, str)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		fmt.Fprintf(w, *textFlag)
	}

}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*listenFlag, nil))
}
