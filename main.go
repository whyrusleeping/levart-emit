package main

import (
	api "github.com/ipfs/go-ipfs-api"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("need an argument")
	}

	sh_local := api.NewShell("localhost:5001")
	sh_gateway := api.NewShell("v04x.ipfs.io")

	finished := make(map[string]struct{})
	tofetch := []string{os.Args[1]}
	for len(tofetch) > 0 {
		next := tofetch[0]
		tofetch = tofetch[1:]

		if _, has := finished[next]; has {
			continue
		}

		obj, err := sh_gateway.ObjectGet(next)
		if err != nil {
			log.Fatal(err)
		}

		outhash, err := sh_local.ObjectPut(obj)
		if err != nil {
			log.Fatal(err)
		}
		if outhash != next {
			log.Fatalf("hash mismatch! %q != %q", next, outhash)
		}

		finished[outhash] = struct{}{}
		for _, lnk := range obj.Links {
			tofetch = append(tofetch, lnk.Hash)
		}
	}
}
