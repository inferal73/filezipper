package main

import (
	"flag"
	"fmt"
	"github.com/inferal73/filezipper/internal/app/filezipper"
	"log"
	"os"
)

var (
	Version string
	entry string
	out string
)

func init()  {
	flag.StringVar(&entry, "entry", "./files", "path to entry file or folder")
	flag.StringVar(&out, "out", "./zip", "path to output directory")
}

func main()  {
	fmt.Println("filezipper", Version)
	flag.Parse()
	config := filezipper.NewConfig(entry, out)
	if err := filezipper.ZipFiles(config, os.Stdout); err != nil {
		log.Fatal(err)
	}
}