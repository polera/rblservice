package search

import (
	"encoding/json"
	"github.com/polera/gorbl"
	"log"
	"sync"
)

func search(rblList string, targetHost string, results chan<- *gorbl.RBLResults) {
	res := gorbl.Lookup(rblList, targetHost)
	results <- &res
}

func Run(hostAddress string) (res []byte) {

	lists, err := GetLists()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *gorbl.RBLResults)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(lists))

	for _, list := range lists {
		go func(rblList string, targetHost string) {
			search(rblList, targetHost, results)
			waitGroup.Done()
		}(list.HostName, hostAddress)
	}

	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()
		close(results)
	}()

	return toJSON(results)
}

func toJSON(results chan *gorbl.RBLResults) (res []byte) {
	var tmpRes []gorbl.RBLResults
	for result := range results {
		tmpRes = append(tmpRes, *result)
	}
	res, err := json.Marshal(tmpRes)
	if err != nil {
		log.Fatal(err)
	}
	return
}
