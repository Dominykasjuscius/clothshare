package web

import (
	db "clothshare/db"
	"clothshare/web/webconfig"
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	ImageFilePath string = "../images"
)

type Application struct {
	Cfg webconfig.Config
	log *logrus.Logger

	db   *db.MongoRepository
	http *http.Server

	ctx     context.Context
	ctxStop context.CancelFunc
}

func (app *Application) init() {
	var err error
	app.ctx, app.ctxStop = context.WithCancel(context.Background())

	app.http = &http.Server{
		Addr:    app.Cfg.Http.Adress,
		Handler: app.Routes(),
	}

	app.log = logrus.New()

	app.log.Println("starting clothshare...")

	app.db = &db.MongoRepository{
		URI:      app.Cfg.Mongodb.URI,
		Database: app.Cfg.Mongodb.Database,
		PoolSize: app.Cfg.Mongodb.Poolsize,
	}

	if err = app.db.Connect(app.ctx); err != nil {
		app.log.WithError(err).Error("could not connect to mongodb")
	}
}

func (app *Application) Start() {
	app.init()

	app.log.Infof("running http server on %s", app.Cfg.Http.Adress)
	go func() {
		if err := app.http.ListenAndServe(); err != nil {
			app.log.WithError(err).Fatal("http server error")
		}
	}()
}

func (app *Application) Stop() {
	if err := app.http.Close(); err != nil {
		app.log.WithError(err).Error("could not stop http server")
	}
}
