package main

import (
	"flag"
	"fmt"
	"github.com/inferal73/filezipper/internal/app/filezipper"
	"github.com/inferal73/filezipper/internal/app/logger"
	"log"
	"os"
)

var (
	Version string
	entry string
	out string
)

func init()  {
	flag.StringVar(&entry, "entry", "", "path to entry file or folder")
	flag.StringVar(&out, "out", "", "path to output directory")
}

func main()  {
	fmt.Println("filezipper", Version)
	flag.Parse()
	validateFlags([]string{"entry", "out"})
	logger.GetLogger()
	config := filezipper.NewConfig(entry, out)
	if err := filezipper.Zip(config); err != nil {
		log.Fatal(err)
	}
}

func validateFlags(required []string) {
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			_, _ = fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n", req)
			os.Exit(2)
		}
	}
}