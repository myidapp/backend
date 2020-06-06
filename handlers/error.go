package handlers

import (
"encoding/json"
"net/http"
)


type Error int


type appError struct {
	Error string `json:"error"`
}

type appJsonResult struct {
	Result interface{} `json:"result"`
}

func ReturnError(w http.ResponseWriter, s int, e string) {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appError{e})

}


func ReturnResult(w http.ResponseWriter, r interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appJsonResult{r})
}

