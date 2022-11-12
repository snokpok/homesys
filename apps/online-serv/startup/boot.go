package startup

import (
	"fmt"
	"homesys/apps/online-serv/env"
	"homesys/apps/online-serv/mws"
	"homesys/apps/online-serv/routes"
)

// starts the server with given params
func BootServer() {
	e := routes.Setup()
	mws.AttachMiddlewares(e)
	evs := env.GetEnv()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", evs.Port)))
}
