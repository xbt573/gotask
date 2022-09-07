package handlers

import (
	"flag"
	"fmt"
	"net/rpc"
	"os"
	"sort"
	"strings"

	"github.com/xbt573/gotask/pkg/types"
)

type ListOptions struct {
	Flags *flag.FlagSet
}

func ListHandler(options ListOptions) {
	if options.Flags.NArg()+options.Flags.NFlag() > 0 {
		fmt.Println("Too many arguments!")
		options.Flags.Usage()

		os.Exit(1)
	}

	client, err := rpc.DialHTTP("tcp", "localhost:6543")
	if err != nil {
		fmt.Printf("Failed to dial api: %v\n", err)
		os.Exit(1)
	}

	res := types.ListResponse{}
	err = client.Call("Api.List", "", &res)
	if err != nil {
		fmt.Printf("Failed to get list: %v\n", err)
		os.Exit(1)
	}

	sort.Slice(res.Tasks, func(i, j int) bool {
		return res.Tasks[i].Priority > res.Tasks[j].Priority
	})

	out := []string{"ID\tName\tDescription\tPriority\t"}
	for _, task := range res.Tasks {
		out = append(out, fmt.Sprintf("%v\t\"%v\"\t\"%v\"\t%v\t", task.Id, task.Name, task.Description, task.Priority))
	}

	fmt.Println(strings.Join(out, "\n"))
}
