package main

import (
	"algs_datastructs/datastructs"
	"fmt"
)

func main() {
	hashTable := datastructs.InitHashTable()
	list := []string{
		"ERIC",
		"KYLE",
		"KENNY",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println(hashTable.Search("ERIC"))
}
