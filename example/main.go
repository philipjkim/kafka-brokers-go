package main

import (
	"fmt"
	"strings"

	"github.com/philipjkim/kafka-brokers-go"
)

func main() {
	zkServers := []string{"192.168.99.100:2181"}
	c, err := kb.NewConn(zkServers)
	if err != nil {
		panic(err)
	}

	for {
		brokers, ch, err := c.GetW()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%q\n", strings.Join(brokers, ","))
		e := <-ch
		fmt.Printf("%v\n", e.Type)
	}
}