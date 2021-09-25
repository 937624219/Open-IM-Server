package main

import (
	"Open_IM/src/common/db"
	rpcChat "Open_IM/src/rpc/chat/chat"
	"Open_IM/src/utils"
	"flag"
)

func main() {
	db.Init()
	rpcPort := flag.String("port", "", "rpc listening port")
	flag.Parse()
	rpcServer := rpcChat.NewRpcChatServer(utils.StringToInt(*rpcPort))
	rpcServer.Run()
}
