package main

import (
	"fmt"
	"github.com/shinytang6/openpaas-security/client"
)

func main() {
	clientset := client.InitClient()
	fmt.Println(*clientset)
}