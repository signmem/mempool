package http

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/signmem/mempool/g"
	"github.com/signmem/mempool/cache"
	"net/http"
)

type HostResult bool

type Arith int


func (t *Arith) FalconPing(r *http.Request, args *Host, result *HostResult) error {
	hostname := args.HostName
	cache.HostMap.PutHost(hostname)
	*result = true
	return nil
}


func RpcStart() {
	svr := rpc.NewServer()
	svr.RegisterCodec(json.NewCodec(), "application/json")
	svr.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	addr := g.Config().RpcAddr
	port := g.Config().RpcPort
	listen := addr + ":" + port

	registry := new(Arith)
	svr.RegisterService(registry, "")
	rec := mux.NewRouter()
	rec.Handle("/v1/api/host", svr)
	http.ListenAndServe(listen, rec)
}