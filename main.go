package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/snoby/spark-pivot/alertmgr"
)

const spark_pivot_test_room = "Y2lzY29zcGFyazovL3VzL1JPT00vNDVhMjFkZjAtNDlmZi0xMWU3LTg2NmQtOGRmOWI2ZjlhNGM1"

// Read data from a file
func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		fmt.Println("Need at exactly one file name to parse")
		os.Exit(1)
	}
	fileName := os.Args[1]
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: Reading file: %s", fileName)
		os.Exit(1)
	}

	authTOKEN := os.Getenv("AUTHTOKEN")
	if authTOKEN == "" {
		fmt.Println("You must export authTOKEN in your environment!!! exiting")
		os.Exit(1)
	}

	// now call the alertmgr package to handle unmarshalling the data
	var alert alertmgr.AlertMgr

	fmt.Println("Loaded file")
	err = alert.LoadRawData(dat)
	if err != nil {
		fmt.Println("Not able to load data to be processed")
		os.Exit(1)
	}
	err = alert.PrintRawJSON()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	err = alert.UnMarshallJSON()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var Markdown string
	Markdown, err = alert.GenerateMarkDown()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	sparkRoom := "Y2lzY29zcGFyazovL3VzL1JPT00vMmM1ZjYwMDAtMzk3ZS0xMWU3LTg4YzEtNGJjNjhkMWU5YWVj"

	err = alert.SendToSparkRoom(authTOKEN, sparkRoom, Markdown)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
