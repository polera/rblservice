package search

import (
	"encoding/json"
	"os"
)

const dataFile = "lists/lists.json"

type RBL struct {
	Name     string `json:"name"`
	HostName string `json:"host"`
	Enabled  bool   `json:"enabled"`
}

func GetLists() ([]*RBL, error) {

	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lists []*RBL
	err = json.NewDecoder(file).Decode(&lists)

	return lists, err
}
