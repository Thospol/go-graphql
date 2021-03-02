package main

import (
	"flag"
	"strconv"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
	"github.com/thospol/go-graphql/core/config"
	"github.com/thospol/go-graphql/core/jwt"
	"github.com/thospol/go-graphql/core/postgres"
	"github.com/thospol/go-graphql/router"
	"github.com/thospol/go-graphql/servers"
)

func main() {
	environment := flag.String("environment", "local", "set working environment")
	configs := flag.String("config", "configs", "set configs path, default as: 'configs'")

	flag.Parse()

	// Init configuration
	if err := config.InitConfig(*configs, *environment); err != nil {
		panic(err)
	}
	//=======================================================

	// set logrus
	if config.CF.App.Release {
		logrus.SetFormatter(stackdriver.NewFormatter(
			stackdriver.WithService("api"),
			stackdriver.WithVersion("v1.0.0")))
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.Infof("Initial 'Configuration'. %+v", config.CF)
	//=======================================================

	// Init return result
	if err := config.InitReturnResult("configs"); err != nil {
		panic(err)
	}
	//=======================================================

	// Load key jwt
	jwt.LoadKey()
	// =======================================================

	// Create a new connection to our pg database
	conf := postgres.Configuration{
		Host:         config.CF.Postgresql.Host,
		Port:         config.CF.Postgresql.Port,
		Username:     config.CF.Postgresql.Username,
		Password:     config.CF.Postgresql.Password,
		DatabaseName: config.CF.Postgresql.DatabaseName,
	}
	err := postgres.NewConnection(conf)
	if err != nil {
		panic(err)
	}

	// Initialize our api and return a pointer to our router for http.ListenAndServe
	// and a pointer to our db to defer its closing when main() is finished
	r := router.New()
	srv := servers.NewServer(strconv.Itoa(config.CF.App.Port), r)
	srv.ListenAndServeWithGracefulShutdown()
	//=======================================================
}
