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
