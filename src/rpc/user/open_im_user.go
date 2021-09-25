package main

import (
	"Open_IM/src/common/db"
	"Open_IM/src/rpc/user/user"
	"flag"
)

func main() {
	db.Init()
	rpcPort := flag.Int("port", 10100, "rpc listening port")
	flag.Parse()
	rpcServer := user.NewUserServer(*rpcPort)
	rpcServer.Run()
}
