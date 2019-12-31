package main

import (
	"github.com/ucosty/jvm/pkg/classloader"
	"github.com/ucosty/jvm/pkg/jvm"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {

}

// Java classpath browser
func main() {
	metaspace := jvm.NewMetaspace()

	// Load the classpath
	if err := classloader.LoadFromClasspath(os.Args[1], metaspace); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
