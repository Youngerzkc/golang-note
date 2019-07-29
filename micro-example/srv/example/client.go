package main

import (
	"fmt"

	greeterProto "aicloud/srv/example/proto/greeter"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.client.greeter"))

	// Create new greeter client
	greeter := greeterProto.NewGreeterClient("go.micro.srv.greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &greeterProto.HelloRequest{Name: "Johnw2"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
