package main

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "ordersystem/internal/logic"
	_ "ordersystem/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"ordersystem/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
