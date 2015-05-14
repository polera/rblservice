package search

import (
	"github.com/polera/gorbl"
	"log"
	"sync"
)

func search(rblList string, targetHost string, results chan<- *gorbl.RBLResults) {
	res := gorbl.Lookup(rblList, targetHost)

	results <- &res

}

func Run(hostAddress string) {

	lists, err := GetLists()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *gorbl.RBLResults)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(lists))

	for _, list := range lists {
		//	fmt.Printf("List: %s\n", list.HostName)
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

	Display(results)
}

func Display(results chan *gorbl.RBLResults) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for result := range results {
		log.Printf("%s\n\n", result.List)
		for _, listing := range result.Results {
			log.Printf("%s:\n%v\n%v\n\n", listing.Address,
				listing.Listed, listing.Text)
		}
	}
	return
}
