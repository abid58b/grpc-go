package main

import (
	"context"
	"fmt"
	"log"

	"github.com/simplesteph/grpc-golang/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hi I am Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("SOmething Wrong With Client: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Unary RPC")

	var fnum int32
	var snum int32

	fmt.Print("Enter Your First Number: ")
	fmt.Scan(&fnum)
	fmt.Print("Enter Your Second Number: ")
	fmt.Scan(&snum)

	req := calculatorpb.SumRequest{
		FirstNumber:  fnum,
		SecondNumber: snum,
	}

	res, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error While Calling the Sum GRPC: %v", err)
	}
	log.Fatalf("Response From Server : Sum : %v ", res.SumResult)

}
