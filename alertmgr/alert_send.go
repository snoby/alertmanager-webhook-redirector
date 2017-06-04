package alertmgr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

type SparkMessage struct {
	RoomId   string `json:"roomId"`
	Markdown string `json:"markdown"`
}

// formatRequest generates ascii representation of a request

func (alert *AlertMgr) SendToSparkRoom(auth string, sparkRoom string, MarkDown string) (err error) {

	const PostURL string = "https://api.ciscospark.com/v1/messages"

	msg := &SparkMessage{
		RoomId:   sparkRoom,
		Markdown: MarkDown,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	// Create HTTP interface
	client := &http.Client{}

	// returns a pointer to a *Request
	req, err := http.NewRequest("POST", "https://api.ciscospark.com/v1/messages", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	token := "Bearer "
	token += auth
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	// print header out for debugging
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}
	fmt.Println(string(requestDump))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respbody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response respBody:", string(respbody))

	return nil
}
