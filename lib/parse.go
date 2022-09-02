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

// var errorCount int
// var warnCount int

// var errorDrives []string
// var warnDrives []string

func Unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

// LogParse parses logs
func LogParse(path string) (warnCount int, errorCount int, warnDrives []string, errorDrives []string, error error) {
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
			warnDrv = append(warnDrv, match[0])
			warnDrives = Unique(warnDrv)
			color.Set(color.FgYellow, color.Underline)
			fmt.Println(line)
			color.Unset()
		} else if strings.Contains(line, "Puncturing bad") {
			errorCount++
			match := pdRegex.FindStringSubmatch(line)
			// fmt.Println(match)
			errorDrv = append(errorDrv, match[0])
			errorDrives = Unique(errorDrv)
			color.Set(color.FgRed, color.Bold)
			fmt.Println(line)
			color.Unset()
		}
	}
	if warnDrv != nil {
		fmt.Printf("%d total Warnings found on drives %s \n", warnCount, warnDrives)
	}
	if errorDrives != nil {
		fmt.Printf("%d total Errors found on drives %s \n", errorCount, errorDrives)
	}
	return warnCount, errorCount, warnDrives, errorDrives, err
}
