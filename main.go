package main

import (
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/elastic/beats/filebeat/harvester/encoding"
	"github.com/elastic/beats/filebeat/harvester/reader"
	"github.com/elastic/beats/libbeat/common/match"
	"github.com/fatih/color"
)

var (
	strPattern = flag.String("pattern", "", "Rexexp pattern to test")
	logfile    = flag.String("logfile", "", "Logfile example")
	negateFlag = flag.Bool("negate", false, "Negate result")
	matchFlag  = flag.String("match", "after", "after or before")

	colors = []*color.Color{
		color.New(color.FgGreen),
		color.New(color.FgCyan),
	}
)

func main() {
	flag.Parse()

	if *logfile == "" {
		log.Fatal("Logfile cannot be empty")
	}

	if *strPattern == "" {
		log.Fatal("Pattern cannot be empty")
	}

	pattern := match.MustCompile(*strPattern)

	cfg := reader.MultilineConfig{
		Pattern: &pattern,
		Negate:  *negateFlag,
		Match:   *matchFlag,
	}

	f, err := os.Open(*logfile)
	if err != nil {
		log.Fatal("Cannot read logfile", err)
	}

	codecFactory, _ := encoding.FindEncoding("utf-8")

	buffer := bytes.NewBuffer(nil)
	codec, _ := codecFactory(buffer)

	var r reader.Reader
	r, err = reader.NewEncode(f, codec, 4096)
	if err != nil {
		log.Fatalln("Failed to initialize line reader: %v", err)
	}

	r, err = reader.NewMultiline(reader.NewStripNewline(r), "\n", 1<<20, &cfg)
	if err != nil {
		log.Fatalln("failed to initializ reader: %v", err)
	}

	i := 0
	for {
		message, err := r.Next()
		if err != nil {
			break
		}
		i++
		colors[i%len(colors)].Println(string(message.Content))
	}

}
