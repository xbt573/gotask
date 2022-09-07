package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/xbt573/gotask/cmd/cli/handlers"
)

var (
	//go:embed usage
	Usage embed.FS

	help = flag.Bool("help", false, "")
	h    = flag.Bool("h", false, "")

	listCmd   = flag.NewFlagSet("list", flag.ExitOnError)
	newCmd    = flag.NewFlagSet("new", flag.ExitOnError)
	infoCmd   = flag.NewFlagSet("info", flag.ExitOnError)
	deleteCmd = flag.NewFlagSet("delete", flag.ExitOnError)
)

func init() {
	flag.Usage = func() {
		data, _ := Usage.ReadFile("usage/usage.txt")
		fmt.Print(string(data))
	}

	listCmd.Usage = func() {
		data, _ := Usage.ReadFile("usage/list.txt")
		fmt.Print(string(data))
	}

	newCmd.Usage = func() {
		data, _ := Usage.ReadFile("usage/new.txt")
		fmt.Print(string(data))
	}

	infoCmd.Usage = func() {
		data, _ := Usage.ReadFile("usage/info.txt")
		fmt.Print(string(data))
	}

	deleteCmd.Usage = func() {
		data, _ := Usage.ReadFile("usage/delete.txt")
		fmt.Print(string(data))
	}
}

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "list":
		listCmd.Parse(os.Args[2:])
		handlers.ListHandler(handlers.ListOptions{
			Flags: listCmd,
		})

	case "new":
		newCmd.Parse(os.Args[2:])
		handlers.NewHandler(handlers.NewOptions{
			Flags: newCmd,
		})

	case "info":
		infoCmd.Parse(os.Args[2:])
		handlers.InfoHandler(handlers.InfoOptions{
			Flags: infoCmd,
		})

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		handlers.DeleteHandler(handlers.DeleteOptions{
			Flags: deleteCmd,
		})

	default:
		flag.Usage()

		if !(*help || *h) {
			os.Exit(1)
		}
	}
}
