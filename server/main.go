package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	pb "github.com/deepthawtz/percrpc/percentage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (s *server) Calculate(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	n, d := float32(in.Dividend), float32(in.Divisor)
	if d == 0 {
		return nil, status.Error(codes.InvalidArgument, "division by zero")
	}
	perc := d * 100.0 / n
	return &pb.Response{
		Percentage: perc,
	}, nil
}

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen on %d: %v", *port, err)
	}
	s := grpc.NewServer()
	pb.RegisterPercentageServer(s, &server{})
	logrus.Infof("listening on %d", *port)
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("failed to start grpc server: %v", err)
	}
}
