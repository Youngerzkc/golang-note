package main

import (
	//"encoding/json"
	//"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	//"io/ioutil"
	"log"
	//"strings"

	"fmt"

	greeterProto "aicloud/srv/example/proto/greeter"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

var (
	cl greeterProto.GreeterClient
)

func main() {
	r := gin.Default()
	// Note: this is a must
	group := r.Group("/greeter")
	group.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		response, _ := cl.Hello(context.TODO(), &greeterProto.HelloRequest{
			Name: name,
		})
		c.JSON(200, gin.H{
			"message": response.Greeting,
		})
	})

	group.POST("/hellos", func(c *gin.Context) {
		req := greeterProto.PostRequest{}
		//bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		//log.Print(string(bodyBytes))
		//if err := proto.Unmarshal(bodyBytes, &req); err != nil {
		if err := c.BindJSON(&req); err != nil {
			log.Print(err)
			return
		}
		fmt.Printf("%+v", req)
		response, _ := cl.Post(context.TODO(), &req)
		c.JSON(200, response)
	})

	// setup Greeter Server Client
	cl = greeterProto.NewGreeterClient("com.workai.aicloud.srv.greeter", client.DefaultClient)

	service := web.NewService(
		web.Name("com.workai.aicloud.api.greeter"),
	)

	service.Init()

	// Register Handler
	service.Handle("/", r)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
