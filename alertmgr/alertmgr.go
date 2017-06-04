package alertmgr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type AlertMgr struct {
	buf                []byte
	RawJSON            string
	ParsedNotification AlertManagerNative
}

//
// AlertMgr_LoadRawData just saves the JSON buffer to this instance.
//
func (alert *AlertMgr) LoadRawData(buf []byte) (err error) {
	fmt.Println(" Loading buffer into AlertMgr memory")
	alert.buf = buf
	alert.RawJSON = string(buf)
	return nil
}

//
// AlertMgr_PrintRawJSON is a helper funtion that will format and print the passed in
// JSON buffer.  The Buffer will be beautified.
//
func (alert *AlertMgr) PrintRawJSON() (err error) {
	if alert.buf == nil {
		return errors.New("AlertMgr: No data loaded.  Load data first!!!")
	}

	var out bytes.Buffer
	err = json.Indent(&out, alert.buf, "", "  ")
	if err != nil {
		fmt.Println("Not able to Indent the buffer properly")
		fmt.Print(err)
		return err
	}
	alert.RawJSON = out.String()

	return nil

}

//
// AlertMgr_unMarshallJSON - Initial entry point into hanlding parsing a JSON alert from Alert Manager
// Format should be something like
//  list of alerts
//  - list  commonAnnotations
//  - list  commonLabels
//  - externalURL
//  - groupKey
//  - GroupLables list
//
//
func (alert *AlertMgr) UnMarshallJSON() (err error) {
	if alert.buf == nil {
		return errors.New("AlertMgr: No data loaded.  Load data first!!!")
	}
	var theAlert AlertManagerNative
	err = json.Unmarshal(alert.buf, &theAlert)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// Save the beautified JSON
	alert.ParsedNotification = theAlert

	return nil

}

//
//GenerateMarkDown will craft a MarkDown message from the Parsed Notification we have decoded
//
func (alert *AlertMgr) GenerateMarkDown() (Markdown string, err error) {
	// TODO: Parse out cluster name from External URL
	// Generate the MarkDown document that will be sent in the message to the alert room
	Markdown += "# Cluster Alert\n"
	Markdown += fmt.Sprintf("cluster: %s \n", alert.ParsedNotification.ExternalURL)
	Markdown += fmt.Sprintf("## %s \n", alert.ParsedNotification.CommonAnnotations.Summary)
	Markdown += fmt.Sprintf("[Alert URL](%s)\n", alert.ParsedNotification.Alerts[0].GeneratorURL)
	Markdown += fmt.Sprintf("* Status: %s\n", alert.ParsedNotification.Status)
	Markdown += fmt.Sprintf("* Service: %s\n", alert.ParsedNotification.CommonLabels.Service)
	Markdown += fmt.Sprintf("* Severity: %s\n", alert.ParsedNotification.CommonLabels.Severity)
	Markdown += "### Raw Alert\n"
	Markdown += "```JSON\n"
	Markdown += alert.RawJSON
	Markdown += "```\n"

	//	fmt.Printf("%s\n", Markdown)

	return Markdown, nil
}
