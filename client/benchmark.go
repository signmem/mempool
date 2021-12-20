package client

import (
	"bytes"
	"encoding/json"
	rpcjson "github.com/gorilla/rpc/json"
	"github.com/signmem/mempool/g"
	"github.com/signmem/mempool/http"
	"log"
	"strconv"
	nethttp "net/http"
)

func Bench() () {

	num := 1
	endLine := g.Config().TestLine
	for {
		if num >= endLine {
			num = 1
		}
		numStr := strconv.Itoa(num)
		hostname := "falcon-agent-test-" + numStr
		err := RpcPutHost(hostname)
		if err != nil {
			log.Printf("[ERROR] put host %s err:%s\n", hostname, err)
		}
		num += 1
	}
}

func PutHost(hostname string) (err error) {
	server := g.Config().ServerAddr
	url := "http://" + server + "/v1/api/host"

	var hostInfo http.Host
	hostInfo.HostName = hostname
	hostInfoBytes, err := json.Marshal(hostInfo)
	if err != nil {
		return err
	}

	resp, err := http.HttpApiPost(url,hostInfoBytes)
	if err != nil {
		return err
	}
	err = resp.Close()
	if err != nil {
		return err
	}
	return nil
}

func RpcPutHost(hostname string) (err error) {
	serveraddr := g.Config().RpcAddr
	serverport := g.Config().RpcPort
	serverurl := serveraddr + ":" + serverport

	url := "http://" + serverurl + "/v1/api/host"
	args := &http.Host {
		HostName: hostname,
	}

	message, err := rpcjson.EncodeClientRequest("Arith.FalconPing", args)
	if err != nil {
		return err
	}
	req, err := nethttp.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(nethttp.Client)
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var result http.HostResult
	err = rpcjson.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		return err
	}
	return nil
}
