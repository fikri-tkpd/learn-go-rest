package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadFromBody(req *http.Request, result interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(result)
	fmt.Print("result:", req.Body)

	if err != nil {
		panic(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)

	if err != nil {
		panic(err)
	}

}
