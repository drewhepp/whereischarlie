package main

import (
  "bytes"
  "encoding/json"
  "encoding/binary"
  "fmt"
  "math"
  "net"
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

func float32FromBytes(bytes []byte) float32 {
    bits := binary.LittleEndian.Uint32(bytes)
    float := math.Float32frombits(bits)
    return float
}

func main() {
  endpoint := "https://whereischarlie.org/position"
  s, err := net.ResolveUDPAddr("udp4", ":5000")
  if err != nil {
    panic(err)
  }
  conn, err := net.ListenUDP("udp", s)
  if err != nil {
    panic(err)
  }

  buf := make([]byte, 65535)
  for {
    n, _, err := conn.ReadFrom(buf)
    if err != nil {
      panic(err)
    }
    if n != 8 {
      fmt.Printf("Recv'ed %s bytes, should be 8\n", string(n))
      return
    }

    lat := float32FromBytes(buf[:4])
    lng := float32FromBytes(buf[4:])

    position.Lat = lat
    position.Lng = lng
    fmt.Println(lat)
    fmt.Println(lng)
    fmt.Println()

    val, _ := json.Marshal(position)
    r, err := http.Post(endpoint, "application/json", bytes.NewBuffer(val))
    if(err != nil) {
      panic(err)
    }
    fmt.Println(r)
  }
}

