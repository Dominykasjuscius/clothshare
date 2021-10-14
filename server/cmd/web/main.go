package main

import (
	web "clothshare/web"
	"clothshare/web/webconfig"
	"os"
	"os/signal"
	"syscall"
)

var (
	kills = []os.Signal{
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	}
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, kills...)

	app := web.Application{
		Cfg: webconfig.Config{
			Mongodb: webconfig.MongoDatabase{
				URI:      "mongodb://127.0.0.1:27017",
				Database: "clothshare",
				Poolsize: 200,
			},
			Http: webconfig.HTTPServer{
				Adress: "127.0.0.1:8088",
			},
		},
	}

	app.Start()
	<-stop
}
