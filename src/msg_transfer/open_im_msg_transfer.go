package main

import (
	"Open_IM/src/common/db"
	"Open_IM/src/msg_transfer/logic"
	"sync"
)

func main() {
	db.Init()
	var wg sync.WaitGroup
	wg.Add(1)
	logic.Init()
	logic.Run()
	wg.Wait()
}
