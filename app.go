package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Messages struct {
	State bool		`json:state`
	Data []string `json:data`
}

func main() {
	for {
		resp, _ := http.Get("https://pikabot.herokuapp.com/switch/")

		body, _ := ioutil.ReadAll(resp.Body)
		var m Messages
		json.Unmarshal(body, &m)
		if m.State==true{
				fmt.Println("Hay acciones")
		}

		time.Sleep(5000)
	}
}
