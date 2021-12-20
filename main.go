package main

import (
	"flag"
	"fmt"
	"github.com/signmem/mempool/cache"
	"github.com/signmem/mempool/client"
	"github.com/signmem/mempool/g"
	"github.com/signmem/mempool/http"
	"os"
	"runtime"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		version := g.Version
		fmt.Printf("%s", version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.InitLog()

	runtime.GOMAXPROCS(20)

	if g.Config().Role == "server" {
		go cache.ResetMetric()
		go http.Start()
		go http.RpcStart()
	}

	if g.Config().Role == "client" {
		go client.Bench()
	}

	select{}
}