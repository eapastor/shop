package main

import (
	"github.com/eapastor/shop/cmd"
	"github.com/eapastor/shop/testData"
)

func main() {

	testData.Init()

	cmd.Start()

}
