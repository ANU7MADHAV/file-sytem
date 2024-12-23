package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"
)

type Config struct {
	IP   net.IP
	Port int
}

type Application struct {
	udpConn *net.UDPConn
	cfg     Config
	clients *sync.Map
	logger  *log.Logger
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewApplication(cfg Config) (*Application, error) {
	ctx, cancel := context.WithCancel(context.Background())
	log := log.New(os.Stdout, "[udpserver]", log.LstdFlags)

	addrs := &net.UDPAddr{
		IP:   cfg.IP,
		Port: cfg.Port,
		Zone: "",
	}

	conn, err := net.ListenUDP("udp", addrs)

	if err != nil {
		log.Fatal(err)
	}

	return &Application{
		udpConn: conn,
		cfg:     cfg,
		clients: &sync.Map{},
		logger:  log,
		ctx:     ctx,
		cancel:  cancel,
	}, nil
}

func main() {
	cfg := Config{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	}

	app, err := NewApplication(cfg)

	if err != nil {
		log.Fatal(err)
	}

	app.HandleUdp()
}
