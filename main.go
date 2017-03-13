package main

import (
	"fmt"

	"github.build.ge.com/aviation-intelligent-network/px-go-micro-svc-tmp/rest"
)

func main() {
	fmt.Println("-------------------------------------------")
	fmt.Println("Starting px-application-name ...")
	fmt.Println("-------------------------------------------")

	//  start rest interface here
	rest.StartServer()
}