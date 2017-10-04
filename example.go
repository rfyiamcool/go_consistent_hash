package main

import (
	"fmt"

	"github.com/rfyiamcool/go_consistent_hash"
)

func main(){

		nodeWeight := make(map[string]int)
		nodeWeight["node1"] = 1
		nodeWeight["node2"] = 1
		nodeWeight["node3"] = 2
		vitualSpots := 100
		hash := hashnode.NewHashRing(vitualSpots)
		//add nodes
		hash.AddNodes(nodeWeight)
		//remove node
		hash.RemoveNode("node3")

		//add node
		hash.AddNode("node3", 3)

		//get key's node
		node := hash.GetNode("key")
		fmt.Println(node)

}
