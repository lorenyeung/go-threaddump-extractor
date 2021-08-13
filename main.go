package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lorenyeung/go-threaddump-extractor/helpers"

	log "github.com/sirupsen/logrus"
)

var gitCommit string
var version string

func printVersion() {
	fmt.Println("Current build version:", gitCommit, "Current Version:", version)
}

func main() {
	versionFlag := flag.Bool("v", false, "Print the current version and exit")
	flags := helpers.SetFlags()
	switch {
	case *versionFlag:
		printVersion()
		return
	}
	helpers.SetLogger(flags.LogLevelVar)

	if flags.LogFileVar == "" {
		log.Error("Please provide file name with -file")
		os.Exit(1)
	}
	for {
		if _, err := os.Stat(flags.LogFileVar); os.IsNotExist(err) {
			// path/to/whatever does not exist
			log.Error("File:", flags.LogFileVar, " does not exist, please try again:")
			reader := bufio.NewReader(os.Stdin)
			downloadIn, _ := reader.ReadString('\n')
			flags.LogFileVar = strings.TrimSuffix(downloadIn, "\n")
			if flags.LogFileVar == "n" {
				os.Exit(0)
			}
		} else {
			break
		}
	}
	log.Info("Scanning:", flags.LogFileVar)
	scanForLines(flags.LogFileVar, flags)
}

func scanForLines(path string, flags helpers.Flags) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	lineNum := 1
	fileCount := 0
	write := false
	// https://golang.org/pkg/bufio/#Scanner.Scan
	prevline := ""
	var fo *os.File
	var datawriter *bufio.Writer
	for scanner.Scan() {
		//need to get line before
		line := scanner.Text()
		if strings.Contains(scanner.Text(), flags.TdBeginStringVar) {
			fileCount++
			write = true
			log.Info("Begin line found:", lineNum, ",", scanner.Text(), ",", prevline)
			//begin file buffer
			fo, err = os.Create(flags.TdFilePrefixVar + "-" + strconv.Itoa(fileCount))
			if err != nil {
				panic(err)
			}
			datawriter = bufio.NewWriter(fo)
			_, _ = datawriter.WriteString(prevline + "\n")
		}
		if write {
			_, _ = datawriter.WriteString(scanner.Text() + "\n")
			//stream line into file
		}
		if strings.Contains(scanner.Text(), flags.TdEndStringVar) {
			log.Info("End line found:", lineNum, ",", scanner.Text())
			datawriter.Flush()
			fo.Close()
			write = false
			//end file buffer
		}
		prevline = line
		lineNum++
	}
	log.Info("Tds found:", fileCount)

	if err := scanner.Err(); err != nil {
		// Handle the error
		return 0, err
	}
	return 0, err
}
