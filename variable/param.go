package variable

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

type Subscribe struct {
	LocalIP      string
	LocalPort    int
	PingInterval int
	Clients      map[string][]*InnerClient `json:"clients"`
}

func (s *Subscribe) ToJson() string {
	d, _ := json.Marshal(s)
	return string(d)
}

func (s *Subscribe) AddClient(cli *InnerClient) {
	clients := s.Clients[cli.Topic]
	if clients == nil {
		s.Clients = make(map[string][]*InnerClient)
		clients = make([]*InnerClient, 0)
		s.Clients[cli.Topic] = clients
	}
	if _, ok := IsArray(clients, cli); !ok {
		clients = append(clients, cli)
	}
	s.Clients[cli.Topic] = clients
}

func (s *Subscribe) DelClient(cli *InnerClient) {
	clients := s.Clients[cli.Topic]
	if clients != nil {
		index, ok := IsArray(clients, cli)
		if ok {
			tmp := append(clients[:index], clients[index+1:]...)
			s.Clients[cli.Topic] = tmp
		}
	}
}

func (s *Subscribe) SendMsg(topic string, msg []byte) {
	clients := s.Clients[topic]
	for _, cli := range clients {
		cli.MsgCh <- msg
	}
}

func (s *Subscribe) HealthCheck(l *sync.WaitGroup) {
	defer l.Done()
	for {
		lost := make([]*InnerClient, 0)
		for _, cliList := range s.Clients {
			for _, cli := range cliList {
				if cli.Running == false {
					lost = append(lost, cli)
				}
			}
		}
		for _, cli := range lost {
			s.DelClient(cli)
		}
		time.Sleep(time.Second * 2)
	}
}

func (s *Subscribe) UpdateClient(t string, c string, msg []map[string]interface{}) {
	for _, client := range msg {
		cid := client["client_id"].(string)
		group := client["group"].(string)
		host := client["host"].(string)
		port := int(client["port"].(float64))
		topic := client["topic"].(string)
		client := InnerClient{ClientID: cid,
			Topic: topic, Group: group, Host: host, Port: port, PingTimer: time.Second * time.Duration(s.PingInterval)}
		if client.Port == s.LocalPort && client.Host == s.LocalIP {
			continue
		}
		_, ok := IsArray(s.Clients[topic], &client)
		if !ok {
			err := client.NewConn()
			if err != nil {
				log.Print(err)
			}
			if client.Topic == t && client.ClientID == c {
				continue
			}
			s.AddClient(&client)
		}
	}
}

func (s *Subscribe) Close() {
	for _, clients := range s.Clients {
		for _, cli := range clients {
			cli.CloseCh <- true
		}
	}
}
