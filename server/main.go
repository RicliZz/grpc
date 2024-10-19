package main

import (
	"context"
	pb "github.com/RicliZz/grpc/protos"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"strings"
)

var min_val int64 = 1
var max_val int64 = 100

type Server struct {
	pb.UnimplementedRandomServer
}

func main() {
	lis, err := net.Listen("tcp", ":54404")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterRandomServer(server, &Server{})
	log.Println("SUCCESS")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *Server) GetRandNum(ctx context.Context, req *pb.RandNumRequest) (*pb.RandNumResponse, error) {
	rand.New(rand.NewSource(req.GetSeed()))
	var temp int
	temp = random(int(min_val), int(max_val))
	return &pb.RandNumResponse{
		Num: int64(temp),
	}, nil
}

func (s *Server) GetRandPass(ctx context.Context, req *pb.RandPassRequest) (*pb.RandPassResponse, error) {
	length := req.GetLength()
	res := strings.Builder{}
	startChar := '!'
	res.Grow(int(length))
	for i := 0; i < int(length); i++ {
		temp := random(int(startChar), 126)
		res.WriteByte(byte(temp))
	}
	return &pb.RandPassResponse{Pass: res.String()}, nil
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
