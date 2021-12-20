package cache

import (
	"sync"
	"time"
)

type SaftHostMap struct {
	sync.Mutex
	MemMap   map[string]int64
}

var (
	HostMap = &SaftHostMap{MemMap: make(map[string]int64)}
	READ = 0
	WRITE = 0
)


func (this *SaftHostMap)  GetHost(hostname string) (int64) {
	this.Lock()
	defer this.Unlock()
	// timeStamp, exists := this.MemMap[hostname]
	timeStamp, exists := HostMap.MemMap[hostname]
	if exists != true {
		timeStamp = 0
	}
	return timeStamp
}

/*
func (this *SaftHostMap) Init() {
	m, err := db.QueryHosts()
	if err != nil {
		return
	}

	this.Lock()
	defer this.Unlock()
	this.MemMap = m
}
*/

func (this *SaftHostMap) PutHost(hostname string) {
	this.Lock()
	defer this.Unlock()
	timeStamp := time.Now().Unix()
	HostMap.MemMap[hostname] = timeStamp

	WRITE += 1

	return
}


func (this *SaftHostMap) DelHost(hostname string) {
	this.Lock()
	defer this.Unlock()
	delete(this.MemMap, hostname)
	return
}


func (this *SaftHostMap) GetAllHostname() ( hostnames []string) {
	this.Lock()
	defer this.Unlock()
	for host, _ := range this.MemMap {
		hostnames = append(hostnames, host)
	}
	return
}