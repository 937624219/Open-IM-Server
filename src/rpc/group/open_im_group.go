package main

import (
	"Open_IM/src/common/db"
	"Open_IM/src/rpc/group/group"
	"flag"
)

func main() {
	db.Init()
	rpcPort := flag.Int("port", 10500, "get RpcGroupPort from cmd,default 16000 as port")
	flag.Parse()
	rpcServer := group.NewGroupServer(*rpcPort)
	rpcServer.Run()
}
