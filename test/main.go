package main

import (
	"fmt"
	"os"

	"github.com/mdlayher/wifi"
)

func main() {
	cli, err := wifi.New()
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		os.Exit(1)
	}

	if err := cli.WiPhys(); err != nil {
		fmt.Printf("Failed to get WiPhys: %v\n", err)
		os.Exit(1)
	}
}

func getInterface(cli *wifi.Client, name string) (*wifi.Interface, error) {
	ifis, err := cli.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, ifi := range ifis {
		if ifi.Name == name {
			return ifi, nil
		}
	}

	return nil, nil
}
