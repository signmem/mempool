package http

import (
	"github.com/signmem/mempool/g"
	"github.com/signmem/mempool/cache"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


func putHost() {
	http.HandleFunc("/v1/api/put", func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadGateway)
		}

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			http.Error(w, "body not json format", http.StatusBadGateway)
		}

		var hostInfo Host

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("[ERROR] io read err:%s\n", err)
		}

		err = json.Unmarshal(body, &hostInfo)

		if err != nil {
			http.Error(w, "body json unmarshar error", http.StatusBadGateway)
		}

		hostname := hostInfo.HostName
		if g.Config().Debug {
			log.Printf("[DEBUG] hostname is %s\n", hostname)
			log.Printf("[DEBUG]  info: %v\n", hostInfo)
		}

		cache.HostMap.PutHost(hostname)

		w.Write([]byte("success"))
	})
}

func getHost() {
	http.HandleFunc("/v1/api/get", func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadGateway)
		}

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			http.Error(w, "body not json format", http.StatusBadGateway)
		}

		var hostInfo Host

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("[ERROR] io read err:%s\n", err)
		}
		log.Printf("[DEBUG] info body is :%s\n", string(body))

		err = json.Unmarshal(body, &hostInfo)

		if err != nil {
			http.Error(w, "body json unmarshar error", http.StatusBadGateway)
		}

		hostname := hostInfo.HostName
		value := cache.HostMap.GetHost(hostname)
		RenderDataJson(w, map[string]interface{}{
			"hostname": hostname,
			"timestamp": value,
		})

	})
}

func delHost() {
	http.HandleFunc("/v1/api/del", func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadGateway)
		}

		headerContentTtype := r.Header.Get("Content-Type")
		if headerContentTtype != "application/json" {
			http.Error(w, "body not json format", http.StatusBadGateway)
		}

		var hostInfo Host

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("[ERROR] io read err:%s\n", err)
		}
		log.Printf("[DEBUG] info body is :%s\n", string(body))

		err = json.Unmarshal(body, &hostInfo)

		if err != nil {
			http.Error(w, "body json unmarshar error", http.StatusBadGateway)
		}

		hostname := hostInfo.HostName
		cache.HostMap.DelHost(hostname)

		w.Write([]byte("success"))
	})
}