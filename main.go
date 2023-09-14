package main

import "github.com/bnb-chain/airdrop-service/cmd"

//go:generate wire ./...

func main() {
	cmd.Execute()
}
