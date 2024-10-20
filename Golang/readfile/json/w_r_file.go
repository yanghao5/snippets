package main

import (
	"encoding/json"
	"log"
	"os"
)

type Monster struct {
	Name  string `json:name`
	Age   int    `json:age`
	Skill string `json:skill`
}

func (m *Monster) Store() error {
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatal("error Marshal struct")
		return err
	}
	filePath := "./data.json"
	outPutFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("error open output file")
		return err
	}
	defer outPutFile.Close()
	encoder := json.NewEncoder(outPutFile)
	err = encoder.Encode(data)
	if err != nil {
		log.Fatalf("error encode data %v", err)
		return err
	}
	return nil

}
// 第一种写入
func (m *Monster) ReStore() error {
	filePath := "./data.json"
	data, err := os.ReadFile(filePath)
  if err != nil {
		log.Fatalf("error open data file")
		return err
	}
  
	err = json.Unmarshal(data, m)
	if err != nil {
		log.Fatalf("error decode json data")
		return err
	}
	return nil

}

// 第二种写入
func (m *Monster) ReStore() error {
	filePath := "./data.json"
	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	defer inputFile.Close()
	if err != nil {
		log.Fatalf("error open data file")
		return err
	}
	decoder := json.NewDecoder(inputFile)
	err = decoder.Decode(m)
	if err != nil {
		log.Fatalf("error decode json data")
		return err
	}
	return nil

}

func main() {

}
