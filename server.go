package main

import (
	"net"
	"sync"

	chat "github.com/Kpatoc452/cli_messanger/gen"
	pb "github.com/Kpatoc452/cli_messanger/gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatServiceServer
	mu      sync.Mutex
	conns   map[pb.ChatService_ConnectServer]bool
	msgChan chan *pb.Message
}

func NewServer() *Server {
	return &Server{
		conns:   make(map[pb.ChatService_ConnectServer]bool),
		msgChan: make(chan *pb.Message),
	}
}

func (s *Server) Connect(stream pb.ChatService_ConnectServer) error {
	s.mu.Lock()
	s.conns[stream] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.conns, stream)
		s.mu.Unlock()
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				logrus.Error(err)
				return
			}

			s.msgChan <- msg
		}
	}()

	for msg := range s.msgChan {
		s.mu.Lock()
		for client := range s.conns {
			if err := client.Send(msg); err != nil {
				logrus.Error(err)
				return err 
			}
		}
		s.mu.Unlock()
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil { 
		logrus.Fatal(err)
	}

	gsrv := grpc.NewServer()
	chat.RegisterChatServiceServer(gsrv, NewServer())
	err = gsrv.Serve(lis)
	if err != nil {
		logrus.Fatal(err)
	}

}
