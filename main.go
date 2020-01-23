package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	lines := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, sc.Err()
}

func main() {

	flag.Parse()

	// Open and parse targets file into a slice
	targetsFile := flag.Arg(0)
	if targetsFile == "" {
		log.Fatalln("Must provide path to targets file")
	}

	targets, err := readLines(targetsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open targets file: %s\n", err)
		os.Exit(1)
	}

	// Open and parse dorks file into a slice
	dorksFile := flag.Arg(1)
	if dorksFile == "" {
		log.Fatalln("Must provide path to dorks file")
	}

	dorks, err := readLines(dorksFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open dorks file: %s\n", err)
		os.Exit(1)
	}

	// Combine dorks with targets to output URL list of clickable github dorks
	for _, d := range dorks {
		for _, t := range targets {
			fmt.Printf("https://github.com/search?q=%%22%s%%22+%s&type=Code\n",
				url.QueryEscape(t),
				url.QueryEscape(d))
		}
	}
}
