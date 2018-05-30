package main

import (
	"flag"
	"net"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/discv5"
)

type bootnodes []*discv5.Node

func (f *bootnodes) String() string {
	return "discv5 nodes"
}

// Set unmarshals enode into discv5.Node.
func (f *bootnodes) Set(value string) error {
	n, err := discv5.ParseNode(value)
	if err != nil {
		return err
	}
	*f = append(*f, n)
	return nil
}

func runBootnode(listenAddr string, nursery bootnodes) *discv5.Network {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Crit("Failed to generate ecdsa key from", "error", err)
	}
	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		log.Crit("Unable to resolve UDP", "address", listenAddr, "error", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Crit("Unable to listen on udp", "address", addr, "error", err)
	}

	realaddr := conn.LocalAddr().(*net.UDPAddr)
	tab, err := discv5.ListenUDP(key, conn, realaddr, "", nil)
	if err != nil {
		log.Crit("Failed to create discovery v5 table:", "error", err)
	}
	if err := tab.SetFallbackNodes(nursery); err != nil {
		log.Crit("Failed to set fallback", "nodes", nursery, "error", err)
	}
	return tab
}

func main() {
	var (
		listenAddr = flag.String("addr", ":0", "listen address")
		verbosity  = flag.Int("verbosity", int(log.LvlInfo), "log verbosity (0-9)")
		vmodule    = flag.String("vmodule", "", "log verbosity pattern")
		nursery    = bootnodes{}
		topic      = flag.String("topic", "whisper", "topic to search for")
		limit      = flag.Int("limit", 5, "search no more than for 5 peers")
		timer      = flag.Duration("timer", 2*time.Minute, "controls how often healthcheck runs")
		period     = flag.Duration("period", 3*time.Second, "controls how often topic is searched")
		statsPort  = flag.String("stats", ":8080", "listen addr for stats")
	)
	flag.Var(&nursery, "n", "These nodes are used to connect to the network if the table is empty and there are no known nodes in the database.")
	flag.Parse()

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(*verbosity))
	if err := glogger.Vmodule(*vmodule); err != nil {
		log.Crit("Failed to set glog verbosity", "value", *vmodule, "err", err)
	}
	log.Root().SetHandler(glogger)
	stats := NewStats(*statsPort)

	healthcheck := func() {
		tab := runBootnode(*listenAddr, nursery)
		defer tab.Close()
		setPeriod := make(chan time.Duration, 1)
		setPeriod <- *period
		found := make(chan *discv5.Node, *limit)
		lookup := make(chan bool, 10)
		log.Info("Started search.", "topic", *topic, "limit", *limit)
		go tab.SearchTopic(discv5.Topic(*topic), setPeriod, found, lookup)
		current := 0
		failTimer := time.After(2 * time.Minute)
		testStart := time.Now()
		stats.Started()
		for {
			select {
			case <-failTimer:
				stats.Failed()
				return
			case <-lookup:
			case n := <-found:
				current++
				latency := time.Since(testStart)
				log.Info("Discovered node", "total", current, "node", n, "latency", latency)
				stats.Discovered(current, latency)
				if current == *limit {
					return
				}
			}
		}
	}
	for {
		healthcheck()
		time.Sleep(*timer)
	}
}
