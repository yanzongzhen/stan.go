package variable

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/yanzongzhen/stan.go/proto"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type InnerClient struct {
	ClientID  string `json:"client_id,omitempty"`
	Topic     string `json:"topic,omitempty"`
	Group     string `json:"group,omitempty"`
	Host      string `json:"host,omitempty"`
	Port      int    `json:"port,omitempty"`
	Cli       *pb.ServerSyncClient
	Conn      *grpc.ClientConn
	Running   bool
	grMu      sync.Mutex
	grWG      sync.WaitGroup
	CloseCh   chan bool
	MsgCh     chan []byte
	PingTimer time.Duration
}

func IsArray(arr []*InnerClient, cli *InnerClient) (int, bool) {
	for idx, item := range arr {
		if item.ClientID == cli.ClientID && item.Topic == cli.Topic {
			return idx, true
		}
	}
	return -1, false
}

func (c *InnerClient) startGoRoutine(f func()) bool {
	var started bool
	c.grMu.Lock()
	if c.Running {
		c.grWG.Add(1)
		go f()
		started = true
	}
	c.grMu.Unlock()
	return started
}

func (c *InnerClient) String() string {
	d, _ := json.Marshal(c)
	return fmt.Sprintf("[InnerClient] %s", d)
}

func (c *InnerClient) Address() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

func (c *InnerClient) NewConn() error {
	con, err := grpc.Dial(c.Address(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewServerSyncClient(con)

	c.grMu.Lock()
	c.Conn = con
	c.Cli = &client
	c.Running = true
	c.CloseCh = make(chan bool)
	c.MsgCh = make(chan []byte)
	c.grMu.Unlock()

	c.startGoRoutine(func() {
		defer c.grWG.Done()
		for {
			_, err := (*c.Cli).Ping(context.Background(), &pb.PingRequest{})
			if err != nil {
				(*c).CloseCh <- true
				break
			}
			time.Sleep((*c).PingTimer)
		}
	})

	c.startGoRoutine(func() {
		defer c.grWG.Done()
		for {
			select {
			case i := <-(*c).MsgCh:
				if (*c).Running {
					data := &pb.MsgRequest{Content: string(i), Topic: c.Topic, Group: c.Group}
					_, _ = (*c.Cli).MsgSync(context.Background(), data)
					break
				}
			case <-(*c).CloseCh:
				c.grMu.Lock()
				(*c).Running = false
				_ = (*c).Conn.Close()
				c.grMu.Unlock()
				break
			}
		}
	})
	return nil
}
