package echoserver

import (
	"context"
	"echoserver/internal/config"
	"echoserver/internal/server"
	"echoserver/pkg/logger"
)

func Run(ctx context.Context) error {

	if err := config.Init(); err != nil {
		return err
	}

	l, err := logger.New(config.Conf.LogFilePath)
	if err != nil {
		return err
	}
	l.Debug = false

	s := server.New(config.Conf.Addr, ctx, l)

	return s.Run()
}
