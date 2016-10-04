package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type server struct {
	host, port string
}

type health struct {
	clusterName string `json:"cluster_name"`
	status      string `json:"status"`
	keys        []string
}

func (s *server) health() {
	url := "http://" + s.host + ":" + s.port + "/_cluster/health"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
	var h health
	var d [][]byte
	fmt.Println(json.Unmarshal(body, &h))
	fmt.Println(h)
	fmt.Println(json.Unmarshal(body, &d))
	fmt.Println(d)
}

func main() {
	s := server{host: "localhost", port: "32769"}
	s.health()
}
