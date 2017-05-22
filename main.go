package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"fmt"
	"os"
)

var sugar *zap.SugaredLogger
var (
	path  string
	debug bool
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()

	app := cli.NewApp()
	app.Name = "ffserver"
	app.Usage = "ffserver --dir app"
	app.Version = "1.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path",
			Usage:       "发布的资源目录",
			Value:       "../../mj_h5\bin-release\native",
			Destination: &path,
		},
		cli.BoolTFlag{
			Name:        "debug, d",
			Usage:       "0:正式版 1:测试版",
			Destination: &debug,
		},
	}

	app.Action = func(c *cli.Context) error {
		sugar.Info(path)
		// Echo instance
		e := echo.New()

		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.Static(path))
		// Start server
		e.Logger.Fatal(e.Start(":80"))

		return nil
	}

	app.Run(os.Args)
}
