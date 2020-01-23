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

	target := ""
	flag.StringVar(&target, "t", "", "")

	dorksFile := ""
	flag.StringVar(&dorksFile, "d", "", "")

	mdOut := ""
	flag.StringVar(&mdOut, "m", "", "")

	flag.Parse()

	if target == "" {
		log.Fatalln("Must provide target keyword or domain")
	}

	if dorksFile == "" {
		log.Fatalln("Must provide path to dorks file")
	}

	dorks, err := readLines(dorksFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open dorks file: %s\n", err)
		os.Exit(1)
	}

	// Combine dorks with targets to output URL list of clickable github dorks
	if mdOut != "" {
		f, err := os.Create(mdOut)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create markdown export file: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		w := bufio.NewWriter(f)

		for _, d := range dorks {
			dURL := fmt.Sprintf("https://github.com/search?q=%%22%s%%22+%s&type=Code",
				url.QueryEscape(target),
				url.QueryEscape(d))
			// Format as markdown list of hyperlinked tasks
			l := fmt.Sprintf("- [ ] [%s](%s)\n", d, dURL)

			_, err := w.WriteString(l)
			if err != nil {
				log.Fatalln("Error writing to markdown export file")
			}
		}

		w.Flush()
	} else {
		for _, d := range dorks {
			fmt.Printf("https://github.com/search?q=%%22%s%%22+%s&type=Code\n",
				url.QueryEscape(target),
				url.QueryEscape(d))
		}
	}
}
