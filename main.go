package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	pattern = flag.String("pattern", "", "Rexexp pattern to test")
	logfile = flag.String("logfile", "", "Logfile example")
	negate  = flag.Bool("negate", false, "Negate result")
)

func main() {
	flag.Parse()

	if *logfile == "" {
		log.Fatal("Logfile cannot be empty")
	}

	if *pattern == "" {
		log.Fatal("Pattern cannot be empty")
	}

	content, err := ioutil.ReadFile(*logfile)
	if err != nil {
		log.Fatal("Cannot read logfile", err)
	}

	regex, err := regexp.Compile(*pattern)
	if err != nil {
		log.Fatal("Failed to compile pattern: ", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	fmt.Printf("matches\tline\n")
	for _, line := range lines {
		matches := regex.MatchString(line)
		if *negate {
			matches = !matches
		}
		if matches {
			color.Green(line)
		} else {
			color.Red(line)
		}
	}
}
