package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

type coords struct {
  Lat float32 `json:"lat"`
  Lng float32 `json:"lng"`
}

var position *coords = &coords{
  Lat: 40.0,
  Lng: -83.0,
}

func positionHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    // Just send out the JSON version of 'position'
    j, _ := json.Marshal(position)
    w.Write(j)
  case "POST":
    // Decode the JSON in the body and overwrite 'position' with it
    d := json.NewDecoder(r.Body)
    p := &coords{}
    err := d.Decode(p)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    position = p
  default:
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintf(w, "I can't do that.")
  }
}

func main() {
  http.Handle("/", http.FileServer(http.Dir("./httpd")))
  http.HandleFunc("/position", positionHandler)

  log.Println("Go!")
  if err := http.ListenAndServe(":80", nil); err != nil {
    panic(err)
  }
}
