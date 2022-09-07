package main

import (
	"embed"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/xbt573/gotask/cmd/daemon/api"
)

var (
	//go:embed usage
	Usage embed.FS

	baseUrl = flag.String("baseUrl", "/var/gotask", "")
)

func init() {
	flag.Usage = func() {
		data, _ := Usage.ReadFile("usage/usage.txt")
		fmt.Print(string(data))
	}
}

func main() {
	flag.Parse()

	err := api.Database.Init(*baseUrl)
	if err != nil {
		panic("Failed to create baseurl.")
	}

	route := new(api.Api)
	rpc.Register(route)

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":6543")
	if err != nil {
		panic("Failed to bind on TCP :6543")
	}

	fmt.Println("Serving on :6543!")
	err = http.Serve(l, nil)
	if err != nil {
		panic(err)
	}
}
