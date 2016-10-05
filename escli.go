package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sfrek/relearn/dto"
)

type server struct {
	host, port string
}

type Mini struct {
	ClusterName string `json:"cluster_name"`
	Status      string `json:"status"`
	TimeOut     bool   `json:"timed_out"`
	Nodes       uint   `json:"number_of_nodes"`
	DataNodes   uint   `json:"number_of_data_nodes"`
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
	var h dto.Health
	fmt.Println(json.Unmarshal(body, &h))
	fmt.Println(h)
	var m Mini
	json.Unmarshal(body, &m)
	fmt.Println(m)
}

func main() {
	s := server{host: "localhost", port: "32769"}
	s.health()
}
