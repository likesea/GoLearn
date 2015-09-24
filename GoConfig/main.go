package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	MachineName string
	Environment string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.getSettings()
	fmt.Println(s.Servers[0])
	fmt.Println(os.Hostname())
	return
}

func (s *Serverslice) getSettings() Serverslice {
	fi, err := os.Open("settings.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	json.Unmarshal(fd, &s)
	return *s
}
