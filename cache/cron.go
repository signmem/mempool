package cache

import (
	"time"
)


func ResetMetric() {
	for {
		READ = WRITE
		WRITE = 0
		time.Sleep(60 * time.Second    )
	}
}