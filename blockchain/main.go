package main

import (
	"fmt"
	"log"
	"sandbox/blockchain/pkg"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	go func() {
		t := time.Now()
		fmt.Println(t.String())
		genesisBlock := pkg.Block{
			Index:     0,
			TimeStamp: t.String(),
			BPM:       0,
			Hash:      "",
			PrevHash:  "",
		}
		spew.Dump(genesisBlock)
		pkg.Blockchain = append(pkg.Blockchain, genesisBlock)
	}()
	log.Fatal(pkg.Run())
}
