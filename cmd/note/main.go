package main

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/note/dal"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/note/rpc"
	notedemo "github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo/noteservice"
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
	tracer2.InitJaeger(constants.NoteServiceName)
	rpc.InitRPC()
	dal.Init()
}
func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	Init()
	svr := notedemo.NewServer(new(NoteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.NoteServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
