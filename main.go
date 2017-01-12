package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Cepave/open-falcon-backend/common/logruslog"
	"github.com/Cepave/open-falcon-backend/common/vipercfg"

	"github.com/masato25/open_lambda_query/conf"
	"github.com/masato25/open_lambda_query/database"
	"github.com/masato25/open_lambda_query/g"
	ginHttp "github.com/masato25/open_lambda_query/gin_http"
	"github.com/masato25/open_lambda_query/graph"
)

func main() {
	vipercfg.Parse()
	vipercfg.Bind()

	if vipercfg.Config().GetBool("version") {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	// config
	vipercfg.Load()
	g.ParseConfig(vipercfg.Config().GetString("config"))
	logruslog.Init()

	// graph
	go graph.Start()

	//lambdaSetup
	database.Init()
	conf.ReadConf()
	go ginHttp.StartWeb()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	select {
	case sig := <-c:
		if sig.String() == "^C" {
			os.Exit(3)
		}
	}
}
