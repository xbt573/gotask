package handlers

import (
	"flag"
	"fmt"
	"net/rpc"
	"os"

	"github.com/google/uuid"
	"github.com/xbt573/gotask/pkg/types"
)

type InfoOptions struct {
	Flags *flag.FlagSet
}

func InfoHandler(options InfoOptions) {
	if options.Flags.NArg() > 1 {
		fmt.Println("Too many args! Id must to be the only argument")
		options.Flags.Usage()

		os.Exit(1)
	}

	if options.Flags.NArg() < 1 {
		fmt.Println("Not enough args! At least id should exists")
		options.Flags.Usage()

		os.Exit(1)
	}

	id, err := uuid.Parse(options.Flags.Arg(0))
	if err != nil {
		fmt.Println("Id is not valid!")
		os.Exit(1)
	}

	client, err := rpc.DialHTTP("tcp", "localhost:6543")
	if err != nil {
		fmt.Printf("Failed to dial api: %v\n", err)
		os.Exit(1)
	}

	res := types.InfoResponse{}
	err = client.Call("Api.Info", types.InfoRequest{
		Id: id,
	}, &res)

	if err != nil {
		fmt.Printf("Failed to fetch info: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("ID\tName\tDescription\tPriority\t\n%v\t\"%v\"\t\"%v\"\t%v\t\n", res.Task.Id,
		res.Task.Name, res.Task.Description, res.Task.Priority)
}
