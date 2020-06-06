package main

import (
	"github.com/myidapp/backend/api"
	"log"
)
import "github.com/google/uuid"

func main(){
	backend := new(api.MYIDAPI)
	backend.Init("http://localhost:8000/api")
	uuid := uuid.New()

	clearResult, err := backend.Clear()
	log.Println("Clear:", err, clearResult)

	postResult, err := backend.PostSignature(uuid.String(), "4342423432", "AFBDB1")
	log.Println("Post:", err, postResult)

	resp, err := backend.GetSignature(uuid.String())
	log.Println("Get:",err, resp)
}
