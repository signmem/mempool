# mempool  

# 目的  

>  1 利用 golang 创建 rcp 客户端, 服务器  
>  2 通过调用 gitlab.com/gorilla/rpc 库进行 json 格式 rcp 调用  
>  3 对 golang 服务器进行压测   
>  4 客户端把生成的主机名记录到服务器变量中   


# cfg.json 使用  

```
{
    "debug":false,
    "loglevel": "debug",
    "logfile": "/path_to/app.log",
    "role": "client",    ( 支持 'client', 'server' )
    "serveraddr": "ipaddr:6061",  // 客户端使用：指定 (服务器端的 http 端口), 可选择使用 http 进行压测 （目前废弃)
    "rpcaddr": "ipaddr",   // 服务器端 rpc 地址  ( client, server 都需要指定)
    "rpcport": "port",     // 服务器 rpc port  ( client, server 都需要指定)
    "testline": 60000,     // 客户端指定压测行数量  
    "http": {
       "enabled": true,                  //  server 都需要指定 启用 http
       "listen": "0.0.0.0:6061"          //  server 端 http 端口 
    }
}

```
