package module

import "fmt"

type Host struct {
	TimeStamp		int64
	HostName 		string
}

func (this *Host)  String() string {
	return fmt.Sprintf(
		"<hostname:%s, timestamp:%d>",
		this.HostName,
		this.TimeStamp,
		)
}