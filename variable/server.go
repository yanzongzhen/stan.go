package variable

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/yanzongzhen/stan.go/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type WatchServer struct {
	Host    string
	Port    int
	MsgChan chan []byte
	Running bool
	grMu    sync.Mutex
	grWG    sync.WaitGroup
	Serve   *grpc.Server
}

type MessageHandler func(msg []byte)

func (server *WatchServer) Address() string {
	return fmt.Sprintf("%v:%v", server.Host, server.Port)
}

func (server *WatchServer) ServerSyncServer() {
}

func (server *WatchServer) MsgSync(ctx context.Context, in *pb.MsgRequest) (*pb.CommonReply, error) {
	resp := make(map[string]string)
	resp["content"] = in.GetContent()
	resp["group"] = in.GetGroup()
	resp["topic"] = in.GetTopic()
	data, _ := json.Marshal(resp)
	server.MsgChan <- data
	return &pb.CommonReply{Message: "Ok", Code: "0000"}, nil
}

func (server *WatchServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PongResponse, error) {
	return &pb.PongResponse{}, nil
}

func (server *WatchServer) StartGoRoutine(f func()) bool {
	var started bool
	server.grMu.Lock()
	if server.Running {
		server.grWG.Add(1)
		go f()
		started = true
	}
	server.grMu.Unlock()
	return started
}

func (server *WatchServer) StartGoRoutineWithParam(f func(cb MessageHandler), cb MessageHandler) bool {
	var started bool
	server.grMu.Lock()
	if server.Running {
		server.grWG.Add(1)
		go f(cb)
		started = true
	}
	server.grMu.Unlock()
	return started
}

func (server *WatchServer) Run(cb MessageHandler) {
	server.grMu.Lock()
	server.Running = true
	server.grMu.Unlock()
	server.StartGoRoutine(server.Server)
	server.StartGoRoutineWithParam(server.ProxyMsg, cb)
}

func (server *WatchServer) Server() {
	defer server.grWG.Done()
	lis, err := net.Listen("tcp", server.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server.Serve = s
	pb.RegisterServerSyncServer(s, server)
	_ = s.Serve(lis)
}

func (server *WatchServer) ProxyMsg(cb MessageHandler) {
	defer server.grWG.Done()
	for {
		select {
		case i := <-server.MsgChan:
			if cb != nil {
				cb(i)
			}
		}
	}
}

func (server *WatchServer) Close() {
	server.Serve.Stop()
}
