package kb

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// NewConn creates a new connection instance.
func NewConn(servers []string) (*Conn, error) {
	c, _, err := zk.Connect(servers, time.Second)
	if err != nil {
		return nil, err
	}
	return &Conn{c}, nil
}

// Conn interacts with the Kafka metadata in Zookeeper.
type Conn struct {
	zkConn *zk.Conn
}

// GetW returns a list of all registered Kafka brokers, and watches that list for changes.
func (c *Conn) GetW() ([]string, <-chan zk.Event, error) {
	children, _, ch, err := c.zkConn.ChildrenW("/brokers/ids")
	if err != nil {
		return nil, nil, err
	}
	result := []string{}
	for _, child := range children {
		var data []byte
		data, _, err = c.zkConn.Get("/brokers/ids/" + child)
		if err != nil {
			return nil, nil, err
		}

		var broker kafkaBroker
		err = json.Unmarshal(data, &broker)
		if err != nil {
			return nil, nil, err
		}
		result = append(result, broker.String())
	}
	return result, ch, nil
}

type kafkaBroker struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (b kafkaBroker) String() string {
	return fmt.Sprintf("%v:%v", b.Host, b.Port)
}
