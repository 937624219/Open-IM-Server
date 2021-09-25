package main

import (
	"Open_IM/src/common/db"
	rpcAuth "Open_IM/src/rpc/auth/auth"
	"flag"
)

func main() {
	db.Init()
	rpcPort := flag.Int("port", 10600, "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := rpcAuth.NewRpcAuthServer(*rpcPort)
	rpcServer.Run()
}
