package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Contestant struct {
	Id      string
	Entries int64
}

func main() {
	contestants := []Contestant{
		{"A", 1},
		{"B", 3},
		{"C", 1},
		{"D", 10},
		{"E", 2},
		{"F", 5},
		{"G", 2},
	}

	// count the tickets
	var tickets int64
	for _, contestant := range contestants {
		tickets += contestant.Entries
	}
	fmt.Printf("Counted %d tickets among %d contestants\n",
		tickets, len(contestants))

	// choose a ticket number
	r, err := rand.Int(rand.Reader, big.NewInt(tickets))
	if err != nil {
		fmt.Println(err)
		return
	}

	// whose ticket is it?
	remaining := r.Int64()
	for _, contestant := range contestants {
		if contestant.Entries == 0 {
			continue
		}
		if remaining >= contestant.Entries {
			remaining -= contestant.Entries
			continue
		}
		fmt.Printf("winner: ticket no. %d, %s\n", r, contestant.Id)
		return
	}
}
