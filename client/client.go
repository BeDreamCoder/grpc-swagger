package main

import (
	pb "grpc-swagger/protos/utlcore"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := pb.NewSmsClient(conn)
	context := context.Background()
	body := &pb.SMSRequest{}

	r, err := c.SendSMS(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Status)
}
