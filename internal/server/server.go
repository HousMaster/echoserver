package server

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net"
	"strings"
	"sync"

	"echoserver/pkg/logger"
)

var (
	ErrStartingServer = errors.New("Can't starting the server")
)

type Server struct {
	addr   string
	ctx    context.Context
	logger *logger.Logger
}

func New(addr string, ctx context.Context, logger *logger.Logger) *Server {
	return &Server{
		addr:   addr,
		ctx:    ctx,
		logger: logger,
	}
}
func (s *Server) Run() error {

	s.logger.Info("Starting server...")

	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		s.logger.Errorf("Starting server: %v\n", err)
		return ErrStartingServer
	}
	defer ln.Close()

	go func() {
		<-s.ctx.Done()
		ln.Close()
	}()

	var wg sync.WaitGroup

RUNLOOP:
	for {
		conn, err := ln.Accept()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Err == net.ErrClosed {
				break RUNLOOP
			}
			s.logger.Warnf("Accepting connection: %v\n", err)
			continue
		}

		wg.Add(1)
		go func() {
			s.handleConnection(conn)
			wg.Done()
		}()
	}

	wg.Wait()
	s.logger.Infof("Shutting down the server...\n")
	return nil
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		s.logger.Infof("Done serving client %v\n", conn.RemoteAddr())
		conn.Close()
	}()
	s.logger.Infof("Serving new conn %v\n", conn.RemoteAddr())

	stop := make(chan struct{})
	go func() {
		rdw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		for {
			message, err := rdw.ReadString('\n')
			if err != nil {
				if !errorIsEOFOrClosedNetworkConn(err) {
					s.logger.Errorf("Reading from connection: %v", err)
				}
				break
			}
			rdw.WriteString("Server: " + strings.ToUpper(message))

			if err := rdw.Flush(); err != nil {
				if !errorIsEOFOrClosedNetworkConn(err) {
					s.logger.Errorf("Reading from connection: %v", err)
				}
				break
			}
		}

		close(stop)
	}()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-stop:
			return
		}
	}

}

func errorIsEOFOrClosedNetworkConn(err error) bool {
	if err == io.EOF {
		return true
	}
	if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
		return true
	}
	return false
}
