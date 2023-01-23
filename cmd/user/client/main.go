package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func main() {
	c, err := userservice.NewClient("example",
		client.WithHostPorts("0.0.0.0:8889"),
		client.WithMuxConnection(1),
	)
	if err != nil {
		log.Fatal(err)
	}

	//req := &userdemo.CreateUserRequest{UserName: "nihao", Password: "123"}
	//resp, err := c.CreateUser(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
	//user_ids := make([]int64, 0)
	//user_ids = append(user_ids, 1)
	//req := &userdemo.MGetUserRequest{UserIds: user_ids}
	//resp, err := c.MGetUser(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)

	req := &userdemo.CheckUserRequest{UserName: "nihao", Password: "123"}
	resp, err := c.CheckUser(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
