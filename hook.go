package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Printf("BODY: %s\n", body)
	}
	var payload interface{}
	err = json.Unmarshal(body, &payload)
	if err == nil {
	printMap(payload)

	m := payload.(map[string]interface{})
	// push_data := m["push_data"] //.(map[string]interface{})

	fmt.Println("callbackURL: ", m["callback_url"])

	data := url.Values{}
	data.Set("state", "success")

	resp, err := http.PostForm(m["callback_url"].(string), data)
	if err != nil {
		fmt.Println("ERROR callback: ", err)
	} else {
		fmt.Println("SUCCESS callback: ", resp)
	}
	

	fmt.Fprintf(w, "This is a hook processor\n")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}


// print the json payload - good for debugging
func printMap(payload interface{}) {
	m := payload.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case nil:
			fmt.Println(k, "is nil")
		case bool:
			fmt.Println(k, "is bool", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			//        fmt.Println(k, "is of a type I don't know how to handle")
			fmt.Printf("-- %s: ", k)
			printMap(v)
			fmt.Println("")
		}
	}
}
