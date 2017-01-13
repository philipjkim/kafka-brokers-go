package main

import (
	"fmt"

	"github.com/philipjkim/kafka-brokers-go"
)

func main() {
	zkServers := []string{"192.168.99.100:2181"}
	c, err := kb.NewConn(zkServers)
	if err != nil {
		panic(err)
	}

	defer c.Close()

	for {
		brokers, ch, err := c.GetW()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%q\n", brokers)
		e := <-ch
		fmt.Printf("%v\n", e.Type)
	}
}
