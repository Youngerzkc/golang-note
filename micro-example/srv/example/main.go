package main

import (
	"log"
	"time"

	greeterProto "aicloud/srv/example/proto/greeter"
	"github.com/micro/go-micro"

	"golang.org/x/net/context"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *greeterProto.HelloRequest, rsp *greeterProto.HelloResponse) error {
	log.Print("Received Say.Hello request")
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func (s *Say) Post(ctx context.Context, req *greeterProto.PostRequest, rsp *greeterProto.PostResponse) error {
	log.Print("hello")
	log.Print(req.RewardRule)
	rsp.Result = "success"
	rsp.RewardRule = req.RewardRule
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("com.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	greeterProto.RegisterGreeterHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
