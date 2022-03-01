package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/simplesteph/grpc-golang/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello i'm Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't Connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	//fmt.Printf("Create Client: %f", c)

	//doUnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting to do in Streaming RPC ...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Abid",
			LastName:  "Sajid",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While Calling GreetManyTimes RPC:%v ", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {

			//we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error While Reading Stream: %v", err)
		}
		log.Printf("Response From GreeetManyTimes: %v", msg.GetResult())
	}

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting to do in Unary RPC ...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Muhammad",
			LastName:  "Abid",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While Calling Greet RPC:%v ", err)
	}
	log.Fatalf("Response From Greeet: %v", res.Result)

}
