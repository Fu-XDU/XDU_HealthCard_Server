package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/urfave/cli.v1"
	"os"
	"xdu-health-card/cron"
	"xdu-health-card/database"
	"xdu-health-card/routes"
	"xdu-health-card/utils/flags"
)

const (
	clientIdentifier = "ServerApp" // Client identifier to advertise over the network
	clientVersion    = "1.0.0"
	clientUsage      = "A golang web server scaffolding using Gin framework"
)

var (
	app       = cli.NewApp()
	baseFlags = []cli.Flag{
		flags.PortFlag,
		flags.AppidFlag,
		flags.SecretFlag,
	}
	mysqlFlags = []cli.Flag{
		flags.MysqlHostFlag,
		flags.MysqlPortFlag,
		flags.MysqlUserFlag,
		flags.MysqlPasswdFlag,
		flags.MysqlDBFlag,
	}
)

func init() {
	app.Action = ServerApp
	app.Name = clientIdentifier
	app.Version = clientVersion
	app.Usage = clientUsage
	app.Commands = []cli.Command{}
	app.Flags = append(app.Flags, baseFlags...)
	app.Flags = append(app.Flags, mysqlFlags...)
}

func ServerApp(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}
	err := prepare(ctx)
	if err != nil {
		log.Error(err)
	}
	return err
}

func prepare(ctx *cli.Context) (err error) {
	if err = database.Setup(); err != nil {
		return
	}
	p := ctx.String("port")
	cron.Start()
	routes.Run(p)
	return
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
