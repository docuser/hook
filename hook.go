package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	var payload interface{}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	b, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	os.Stdout.Write(b)

	m := payload.(map[string]interface{})

	resp, err := sendReply(m["callback_url"].(string))
	if err != nil {
		fmt.Println("ERROR callback:", err)
	}
	fmt.Println("callback to ", m["callback_url"].(string))
	body, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Printf("%s", body)
	} else {
		fmt.Println("no callback BODY")
	}

	fmt.Fprintf(w, "This is a hook processor\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func sendReply(url string) (*http.Response, error) {
	var jsonStr = []byte(`{"state":"success"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
