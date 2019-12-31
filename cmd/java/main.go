package main

import (
	"fmt"
	"github.com/ucosty/jvm/pkg/classloader"
	"github.com/ucosty/jvm/pkg/intrinsics"
	"github.com/ucosty/jvm/pkg/jvm"
	"log"
	"os"
	"strings"
)

func main() {
	metaspace := jvm.NewMetaspace()

	// Load the classpath
	if err := classloader.LoadFromClasspath(os.Args[1], metaspace); err != nil {
		log.Fatal(err)
	}

	if err := intrinsics.PatchSystem(metaspace); err != nil {
		log.Fatal(err)
	}

	metaspace.DumpHeap()
	if err := metaspace.InitClasses(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Main Class: %s\n", os.Args[2])
	mainClassName := strings.Replace(os.Args[2], ".", "/", -1)

	class, err := metaspace.GetClass(mainClassName)
	if err != nil {
		log.Fatal(err)
	}
	//if err := class.Invoke("<clinit>"); err != nil {
	//	log.Fatal(err)
	//}
	if err := class.Invoke("main"); err != nil {
		log.Fatal(err)
	}
}
