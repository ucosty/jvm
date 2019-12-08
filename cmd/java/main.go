package main

import (
	"fmt"
	"java-hackery/pkg/classloader"
	"java-hackery/pkg/intrinsics"
	"java-hackery/pkg/jvm"
	"log"
	"os"
	"strings"
)

func main() {
	metaspace := jvm.NewMetaspace()

	//Load the intrinsics
	for _, class := range intrinsics.Classes() {
		if err := metaspace.LoadFromStruct(&class); err != nil {
			log.Fatalf("failed to load intrinsics: %v\n", err)
		}
	}

	// Load the classpath
	if err := classloader.LoadFromClasspath(os.Args[1], metaspace); err != nil {
		log.Fatal(err)
	}

	metaspace.DumpHeap()

	fmt.Printf("Main Class: %s\n", os.Args[2])
	mainClassName := strings.Replace(os.Args[2], ".", "/", -1)

	class, err := metaspace.GetClass(mainClassName)
	if err != nil {
		log.Fatal(err)
	}
	if err := class.Invoke("<clinit>"); err != nil {
		log.Fatal(err)
	}
	if err := class.Invoke("main"); err != nil {
		log.Fatal(err)
	}
}
