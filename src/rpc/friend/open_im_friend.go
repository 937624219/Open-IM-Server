package main

import (
	"Open_IM/src/common/db"
	"Open_IM/src/rpc/friend/friend"
	"flag"
)

func main() {
	db.Init()
	rpcPort := flag.Int("port", 10200, "get RpcFriendPort from cmd,default 12000 as port")
	flag.Parse()
	rpcServer := friend.NewFriendServer(*rpcPort)
	rpcServer.Run()
}
