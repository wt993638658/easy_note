package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo/noteservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func main() {
	c, err := noteservice.NewClient("example",
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithMuxConnection(1),
	)
	if err != nil {
		log.Fatal(err)
	}

	//req := &notedemo.CreateNoteRequest{Title: "nihao", Content: "xxx", UserId: 1}
	//resp, err := c.CreateNote(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
	searchkey := ""
	req := &notedemo.QueryNoteRequest{UserId: 1, SearchKey: &searchkey, Offset: 0, Limit: 0}
	resp, err := c.QueryNote(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	//req := &notedemo.DeleteNoteRequest{UserId: 1, NoteId: 3}
	//resp, err := c.DeleteNote(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
	//updatetitle := "hhh"
	//updatecontent := "xx"
	//req := &notedemo.UpdateNoteRequest{NoteId: 2, UserId: 1, Title: &updatetitle, Content: &updatecontent}
	//resp, err := c.UpdateNote(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
}
