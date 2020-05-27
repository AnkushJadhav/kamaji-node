package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/AnkushJadhav/kamaji-node/logger"
	"github.com/AnkushJadhav/kamaji-node/pkg/server"
	"github.com/AnkushJadhav/kamaji-node/pkg/server/http"
)

type app struct {
	isProd     bool
	httpServer *http.Server
}

var mainApp *app

// Start starts the kamaji-node server
func Start(cfgFile string) error {
	if mainApp != nil {
		return fmt.Errorf("kamaji-node application has already started")
	}

	logger.Infoln("staring application")
	conf, err := getConfig(cfgFile)
	if err != nil {
		return err
	}
	logger.Infoln("config loaded successfully")

	var logFile = conf.Server.LogFile
	if fileLoggingEnabled(logFile) {
		startFileLogging(logFile)
	}

	logger.Infoln("starting http server")
	httpServer := &http.Server{}
	serverConfig := &server.Config{
		EnableTLS:     false,
		PopulatePool:  true,
		BindIP:        conf.Server.BindIP,
		Port:          conf.Server.Port,
		StorageDriver: nil,
	}
	mainApp = &app{
		httpServer: httpServer,
	}

	if err := startServer(httpServer, serverConfig); err != nil {
		return err
	}

	return nil
}

// Stop gracefully stops the kamaji-node application
func Stop() error {
	if err := mainApp.httpServer.Stop(); err != nil {
		return err
	}
	return nil
}

func fileLoggingEnabled(logFile string) bool {
	return strings.TrimSpace(logFile) != ""
}

func startFileLogging(logFile string) error {
	f, err := os.Open(logFile)
	if err != nil {
		return err
	}
	defer f.Close()

	logger.SetOutput(f)
	return nil
}

func startServer(srv server.Server, conf *server.Config) error {
	if err := srv.Bootstrap(conf); err != nil {
		return err
	}

	if err := srv.Start(); err != nil {
		return err
	}

	return nil
}
