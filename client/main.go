package main

import (
	"context"
	"flag"
	pb "github.com/RicliZz/grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	addr  = flag.String("addr", "localhost:54404", "the address to connect to")
	seed  = flag.Int64("seed", 0, "the seed for random number generator")
	place = flag.Int64("place", 0, "the place to use")
)

func main() {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewRandomClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetRandNum(ctx, &pb.RandNumRequest{Seed: *seed, Place: *place})
	if err != nil {
		log.Fatalf("could not get random number: %v", err)
	}
	log.Printf("random number: %v", r.GetNum())
	r2, err := c.GetRandPass(ctx, &pb.RandPassRequest{Seed: *seed, Length: 8})
	if err != nil {
		log.Fatalf("could not get random password: %v", err)
	}
	log.Println("random password:", r2.GetPass())
}
