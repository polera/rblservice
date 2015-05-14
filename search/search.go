package search

import (
	"fmt"
	_ "github.com/polera/gorbl"
	"log"
	"sync"
)

func Run(hostAddress string) {

	lists, err := GetLists()
	if err != nil {
		log.Fatal(err)
	}

	//results := make(chan *gorbl.RBLResults)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(lists))

	for _, list := range lists {
		fmt.Printf("List: %s\n", list.HostName)
	}
}
