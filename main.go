package main

import (
	"java-hackery/pkg/classloader"
	"java-hackery/pkg/intrinsics"
	"java-hackery/pkg/jvm"
	"log"
	"os"
)

func main() {
	// Create the class map
	classes := make(map[string]*classloader.JavaClass, 1)

	// Load the intrinsics
	for _, class := range intrinsics.Classes() {
		classes[class.Name] = &class
	}

	// Load the class from the CLI
	mainClass, err := classloader.LoadFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	classes[mainClass.Name] = mainClass

	if err := jvm.Invoke(mainClass,"<clinit>"); err != nil {
		log.Fatal(err)
	}
	if err := jvm.Invoke(mainClass,"main"); err != nil {
		log.Fatal(err)
	}
}
