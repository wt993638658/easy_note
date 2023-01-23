package main

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/user/dal"
	userdemo "github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/userdemo/userservice"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/bound"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"
	tracer2 "github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func Init() {
	tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}
	Init()
	svr := userdemo.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleware
		server.WithMiddleware(middleware.ServerMiddleware),                                             // middleware
		server.WithServiceAddr(addr),                                                                   // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),                             // limit
		server.WithMuxTransport(),                                                                      // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                                                // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),
		server.WithRegistry(r), // registry
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
