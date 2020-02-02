package main

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"sandbox/blockchain/pkg"
	"time"
)

func main() {
	go func() {
		t := time.Now()
		genesisBlock := pkg.Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		pkg.Blockchain = append(pkg.Blockchain, genesisBlock)
	}()
	log.Fatal(pkg.Run())
}
