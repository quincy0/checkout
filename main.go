package main

import (
	"github.com/quincy0/checkout/backend"
	"github.com/quincy0/checkout/frontend"
)

func main() {
	token := frontend.GetToken()
	backend.CreateInstrument(token, "wq111")
}
