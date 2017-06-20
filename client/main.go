package main

import (
	"flag"
	"fmt"

	"github.com/Sirupsen/logrus"
	pb "github.com/deepthawtz/percrpc/percentage"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8080", "grpc server to connect to")
	outof := flag.Int64("o", 0, "total number out of")
	n := flag.Int64("n", 0, "number to divide by")
	flag.Parse()

	if *n <= 0 || *outof <= 0 {
		logrus.Fatalf("can't calculate percentage of numbers less than zero")
	}

	if *n > *outof {
		logrus.Fatalf("can't calculate percentages that would be less than zero")
	}

	conn, err := grpc.Dial(*host, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not dial host %s: %v", *host, err)
	}
	defer conn.Close()
	c := pb.NewPercentageClient(conn)
	r, err := c.Calculate(context.Background(), &pb.Request{Dividend: *outof, Divisor: *n})
	if err != nil {
		logrus.Fatalf("could not calculate percentage: %v", err)
	}
	fmt.Printf("%.2f%%\n", r.Percentage)
}
