package main

import (
 "encoding/json"
 "math/rand"
 "net/http"
 "time"
)

type RandomResponse struct {
 Number int `json:"number"`
}

func RandomHandler(w http.ResponseWriter, r *http.Request) {
 rand.Seed(time.Now().UnixNano())
 num := rand.Intn(10) 
 response := RandomResponse{Number: num}

 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(response)
}