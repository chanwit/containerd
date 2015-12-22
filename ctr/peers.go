package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/docker/containerd/api/grpc/types"
	netcontext "golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

var PeersCommand = cli.Command{
	Name:  "peers",
	Usage: "list all peers",
	Subcommands: []cli.Command{
		ListPeerCommand,
		CreateCheckpointCommand,
	},
	Action: listPeers,
}

var ListPeerCommand = cli.Command{
	Name:   "list",
	Usage:  "list all daemon peers",
	Action: listPeers,
}

func listPeers(context *cli.Context) {
	var (
		c = getClient(context)
	)

	// os.MkdirAll(path.Join(os.Getenv("HOME"), ".ctr"), 0744)
	// knownPeers := path.Join(os.Getenv("HOME"), ".ctr", "known_peers")

	// TODO f := os.OpenFile(knownPeers, ..., ...)
	// defer f.Close()

	resp, err := c.ListPeers(netcontext.Background(), &types.PeersRequest{})
	if err != nil {
		fatal(err.Error(), 1)
	}
	w := tabwriter.NewWriter(os.Stdout, 20, 1, 3, ' ', 0)
	fmt.Fprint(w, "PEER\tSTATUS\n")
	for _, p := range resp.Peers {
		fmt.Fprintf(w, "%s\t%s\n", p.Address, p.Status)
	}
	if err := w.Flush(); err != nil {
		fatal(err.Error(), 1)
	}

	save(resp.Peers)
}

func save(peers []*types.Peer) {
	out, _ := yaml.Marshal(peers)
	ioutil.WriteFile("./known_peers", out, 0644)
}

func load() []*types.Peer {
	data, _ := ioutil.ReadFile("./known_peers")
	peers := []*types.Peer{}
	yaml.Unmarshal(data, &peers)
	return peers
}
