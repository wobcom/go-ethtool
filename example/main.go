package main

import (
	"gitlab.com/wobcom/golang-ethtool"

	"encoding/json"
	"flag"
	"log"
	"os"
)

func main() {
	ifname := flag.String("interface", "", "Interface name")
	flag.Parse()

	if *ifname == "" {
		log.Fatal("Specify interface with --interface")
	}

	tool, err := ethtool.NewEthtool()
	if err != nil {
		panic(err.Error)
	}

	iface, err := tool.NewInterface(*ifname, false)
	if err != nil {
		panic(err.Error())
	}

	b, err := json.Marshal(iface)
	if err != nil {
		panic(err.Error())
	}
	os.Stdout.Write(b)
}
