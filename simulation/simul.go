package main

import (
	// Service needs to be imported here to be instantiated.
	_ "github.com/csanti/cothority_template/service"
	"github.com/csanti/onet/simul"
)

func main() {
	simul.Start()
}
