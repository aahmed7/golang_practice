package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Server struct which config for a server
type Server struct {
	Server Config `json:"server"`
}

// Config struct
type Config struct {
	Host          string `json:"host"`
	Port          string `json:"port"`
	ServicePrefix string `json:"servicePrefix"`
	ConsulServer  string `json:"consulServer"`
	ServiceName   string `json:"serviceName"`
}

func getSvtPortFromFile() Server {
	// jsonFile, _ := os.Open()
	// defer jsonFile.Close()

	byteValue, _ := ioutil.ReadFile("memsvc_config.json")

	svr := Server{}

	json.Unmarshal([]byte(byteValue), &svr)
	return svr
}

func main() {
	serverinfo := getSvtPortFromFile()
	fmt.Print("Server port :" + serverinfo.Server.Port)
}
