package main

import (
	"encoding/json"
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

func GetEnvironment() string {
	var s Serverslice
	envMap := s.getSettings()
	osName, _ := os.Hostname()
	if env, ok := envMap[osName]; ok {
		return env
	}
	if env, ok := envMap["default"]; ok {
		return env
	}
	return ""
}
func (s *Serverslice) getSettings() map[string]string {
	fi, err := os.Open("settings.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	json.Unmarshal(fd, &s)
	envMap := make(map[string]string)
	for _, v := range s.Servers {
		envMap[v.MachineName] = v.Environment
	}
	return envMap
}
