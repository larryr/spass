package main

import (
	"flag"
	"fmt"
	"spass"
)

/*
var combineCmd = &cobra.Command{
	Use:   "combine",
	Short: "Reconstruct a secret from the parts read from stdin",
	Run:   runCombine,
}
*/

var parts = flag.Int("p", 3, "Number of shares")
var threshold = flag.Int("t", 2, "threshold of shares required to reconstruct secret")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Printf("need a command\n")
		return
	}

	switch args[0] {
	case "split":
		spass.Split(*parts, *threshold)
	case "combine":
		spass.Combine()
	default:
		{
			fmt.Printf("unknown command:%s\n", args[0])
			return
		}
	}
}
