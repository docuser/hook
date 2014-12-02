package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	var payload interface{}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	printMap(payload)

	m := payload.(map[string]interface{})

	resp, err := sendReply(m["callback_url"].(string))
	if err != nil {
		fmt.Println("ERROR callback: ", err)
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

// print the json payload - good for debugging
func printMap(payload interface{}) {
	printMapIndent(payload, "")
}
func printMapIndent(payload interface{}, indent string) {
	m := payload.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(indent, k, "is string", vv)
		case nil:
			fmt.Println(indent, k, "is nil")
		case bool:
			fmt.Println(indent, k, "is bool", vv)
		case int:
			fmt.Println(indent, k, "is int", vv)
		case float64:
			fmt.Println(indent, k, "is float64", vv)
		case []interface{}:
			fmt.Println(indent, k, "is an array:")
			for i, u := range vv {
				fmt.Println(indent, i, u)
			}
		default:
			//        fmt.Println(indent, k, "is of a type I don't know how to handle")
			fmt.Println(indent, k)
			printMapIndent(v, "  "+indent)
			fmt.Println("")
		}
	}
}
