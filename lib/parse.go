package lib

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"
)

var errorCount int
var warnCount int

var errorDrives []string
var warnDrives []string

// LogParse parses logs
func LogParse(path string) {
	pdRegex := regexp.MustCompile(`PD [0-9][0-9,A-F]`)
	var warnDrv []string
	var errorDrv []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	var line string
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, "Predictive failure") || strings.Contains(line, "3/11/00") {
			warnCount++
			match := pdRegex.FindStringSubmatch(line)
			warnDrv = append(warnDrives, match[0])
			color.Set(color.FgYellow, color.Underline)
			fmt.Println(line)
			color.Unset()
		} else if strings.Contains(line, "Puncturing bad") {
			errorCount++
			match := pdRegex.FindStringSubmatch(line)
			errorDrv = append(errorDrives, match[0])
			color.Set(color.FgRed, color.Bold)
			fmt.Println(line)
			color.Unset()
		}
	}
	fmt.Printf("Warnings found on drives %s", warnDrv)
	if errorDrv != nil {
		fmt.Printf("Errors found on drives %s", errorDrv)
	}
	return
}
