package handlers

import (
	"flag"
	"fmt"
	"net/rpc"
	"os"
	"strconv"

	"github.com/xbt573/gotask/pkg/types"
)

type NewOptions struct {
	Flags *flag.FlagSet
}

func NewHandler(options NewOptions) {
	if options.Flags.NArg() > 3 {
		fmt.Println("Too many args!")
		options.Flags.Usage()

		os.Exit(1)
	}

	if options.Flags.NArg() < 1 {
		fmt.Println("Not enough arguments! At least name should exists")
		options.Flags.Usage()

		os.Exit(1)
	}

	name := options.Flags.Arg(0)
	description := options.Flags.Arg(1)
	priority, err := strconv.ParseInt(options.Flags.Arg(2), 10, 0)
	if err != nil {
		priority = 0
	}

	client, err := rpc.DialHTTP("tcp", "localhost:6543")
	if err != nil {
		fmt.Printf("Failed to dial api: %v\n", err)
		os.Exit(1)
	}

	res := types.NewResponse{}
	err = client.Call("Api.New", types.NewRequest{
		Name:        name,
		Description: description,
		Priority:    int(priority),
	}, &res)

	if err != nil {
		fmt.Printf("Failed to create task: %v\n", err)
		os.Exit(1)
	}
}
