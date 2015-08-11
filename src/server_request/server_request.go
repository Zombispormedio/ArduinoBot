package server_request

import (

  "objects"
	"encoding/json"

	"io/ioutil"
	"net/http"

)


func  Check()(state *objects.Messages){
  resp, _ := http.Get(objects.Url_pikabot)

  body, _ := ioutil.ReadAll(resp.Body)
  var m objects.Messages
  json.Unmarshal(body, &m)

return &m
}

func Advice(){
  http.Get("https://maker.ifttt.com/trigger/arduino/with/key/hVFVQTeDT9jMuSJYkR8GOgPT0xBX9dwpa00vD4TRELG")
}
