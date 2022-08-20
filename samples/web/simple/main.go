package main

import (
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/server"
	"github.com/silverswords/sand/server/web"
)

func main() {
	sand.Instance().Load(
		func() server.Server {
			return web.CreateBuilder("path").Build(
				web.WithHost("host"),
				web.WithAddr("0"),
			).Run()
		}(),
	).Run()
}
