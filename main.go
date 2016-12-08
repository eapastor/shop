package main

import (
	"github.com/eapastor/shop/cmd"
	"github.com/eapastor/shop/testData"
)

const (
	port int = 3000
)

func main() {

	testData.Init()

	cmd.Start()

}
