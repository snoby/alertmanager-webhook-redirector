package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/snoby/spark-pivot/alertmgr"
)

const spark_pivot_test_room = "Y2lzY29zcGFyazovL3VzL1JPT00vNDVhMjFkZjAtNDlmZi0xMWU3LTg2NmQtOGRmOWI2ZjlhNGM1"

// Read data from a file
func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 2 {
		fmt.Println("Need the JSON passed as an argument and the name of the receiver sent")
		os.Exit(1)
	}
	//	fileName := os.Args[1]
	//	data, err := ioutil.ReadFile(fileName)
	//	if err != nil {
	//		fmt.Println("Error: Reading file: %s", fileName)
	//		os.Exit(1)
	//	}
	// The entire JSON payload is being passed to us as an argument
	payload := os.Args[1]
	if payload == "" {
		fmt.Println("first argument didn't contain the alert message")
		os.Exit(1)
	}

	//TODO: need some sanitizing of the input here
	receiver := os.Args[2]

	authTOKEN := os.Getenv("AUTHTOKEN")
	if authTOKEN == "" {
		fmt.Println("You must export AUTHTOKEN in your environment!!! below is what I found")
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			fmt.Println(pair[0])
		}
		os.Exit(1)
	}
	fmt.Println("payload, receiver and authTOKEN received and accepted, parsing...")

	// now call the alertmgr package to handle unmarshalling the data
	var alert alertmgr.AlertMgr
	alert.LoadReceiverMappings()

	fmt.Println("Loaded file")
	dat := []byte(payload)
	err := alert.LoadRawData(dat)
	if err != nil {
		fmt.Println("Not able to load data to be processed")
		os.Exit(1)
	}

	err = alert.SaveReceiver(receiver)
	if err != nil {
		fmt.Print(err)
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

	sparkRoom := alert.GetRoomId()
	fmt.Printf("Sending to roomID: %s because receiver was: %s \n", sparkRoom, receiver)

	err = alert.SendToSparkRoom(authTOKEN, sparkRoom, Markdown)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
