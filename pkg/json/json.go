package json

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpRequestUnmarshal(r *http.Request, res interface{}) interface{} {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	log.Println("\n create order body = ", string(body))
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("Error marshalling json data :", err)
		return nil
	}
	return res
}

func JsonResposeEncoder(w http.ResponseWriter, r *http.Request, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func JsonEncode(data interface{}) string {
	s, _ := json.Marshal(data)
	return string(s)
}
