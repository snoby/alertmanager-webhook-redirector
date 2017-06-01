package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/snoby/alertmanager-webhook-redirector/alertmgr"
)

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

	// now call the alertmgr package to handle unmarshalling the data
	var alert alertmgr.AlertMgr

	fmt.Println("Loaded file")
	err = alert.AlertMgr_LoadRawData(dat)
	if err != nil {
		fmt.Println("Not able to load data to be processed")
		os.Exit(1)
	}
	err = alert.AlertMgr_PrintRawJSON()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

}
