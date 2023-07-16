package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	api "github.com/lvkeliang/kitexstu/kitex_gen/api"
	"github.com/lvkeliang/kitexstu/kitex_gen/api/echo"
	"log"
	"time"
)

func main() {
	cli, err := echo.NewClient("echo", client.WithHostPorts("localhost:8888"))
	if err != nil {
		log.Fatal(err)
	}

	req := &api.Request{Message: "Hello, Kitex!"}
	resp, err := cli.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Message)
}
