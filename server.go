package main

import (
	"context"

	pb "github.com/Kpatoc452/cli_messanger/gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Connection struct {
	pb.User
	stream grpc.ServerStreamingServer[pb.MessageResponse]
}

type Server struct {
	pb.UnimplementedChatServer
	conns   map[int64]*Connection
	msgChan chan *pb.MessageRequest
}

func NewServer() *Server {
	return &Server{
		conns:   make(map[int64]*Connection),
		msgChan: make(chan *pb.MessageRequest),
	}
}

func (s *Server) CreateConnect(req *pb.ConnectRequest, stream grpc.ServerStreamingServer[pb.MessageResponse]) error {
	s.conns[req.User.Id] = &Connection{
		*req.User,
		stream,
	}
	return nil
}

func (s *Server) SendMsg(ctx context.Context, req *pb.MessageRequest) (*pb.None, error) {
	s.msgChan <- req
	return &pb.None{}, nil
}

func (s *Server) sendMsgs(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logrus.Info("Sender Stopped by ctx")
			return
		case msgReq := <-s.msgChan:
			for _, c := range s.conns {
				msgResp := &pb.MessageResponse{
					Id:        msgReq.Id,
					User:      msgReq.User,
					Text:      msgReq.Text,
					Timestamp: msgReq.Timestamp,
				}

				err := c.stream.SendMsg(msgResp)
				if err != nil {
					logrus.Error(err)
				}
			}
		}
	}
}

func (s *Server) Run(ctx context.Context) {
	go s.sendMsgs(ctx)
	logrus.Info("Starting messaging")
}
