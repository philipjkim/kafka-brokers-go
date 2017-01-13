kafka-brokers-go
================

[![GoDoc](https://godoc.org/github.com/philipjkim/kafka-brokers-go?status.svg)](https://godoc.org/github.com/philipjkim/kafka-brokers-go) [![Build Status](https://travis-ci.org/philipjkim/kafka-brokers-go.svg)](https://travis-ci.org/philipjkim/kafka-brokers-go)

Go library to access Kafka broker list in Zookeeper. It's useful when working with [sarama](https://github.com/Shopify/sarama).

## Usage

```go
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
```

## References

- https://github.com/wvanbergen/kazoo-go
- https://github.com/Shopify/sarama
