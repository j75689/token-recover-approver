package main

import "github.com/bnb-chain/token-recover-app/cmd"

//go:generate wire ./...

func main() {
	cmd.Execute()
}
