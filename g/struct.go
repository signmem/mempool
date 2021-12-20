package g

type GlobalConfig struct {
	Debug         bool              `json:"debug"`
	LogLevel   		string			`json:"loglevel"`
	LogFile 		string			`json:"logfile"`
	Database 		string			`json:"database"`
	MaxIdle 		int				`json:"maxidle"`
	Listen 			string 			`json:"listen"`
	Trustable 		[]string 		`json:"trustable"`
	Http 			*HttpConfig 	`json:"http"`
	TestLine 		int 			`json:"testline"`
	ServerAddr 		string			`json:"serveraddr"`
	Role 			string 			`json:"role"`
	RpcAddr 		string 			`json:"rpcaddr"`
	RpcPort 		string			`json:"rpcport"`
}

type  HttpConfig  struct {
	Enabled 		bool 		`json:"enabled"`
	Listen 			string 		`json:"listen"`
}