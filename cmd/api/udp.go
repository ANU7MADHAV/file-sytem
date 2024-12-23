package main

import (
	"net"
)

func (app *Application) HandleUdp() {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-app.ctx.Done():
			return
		default:
			n, remoteAdrs, err := app.udpConn.ReadFrom(buffer)
			if err != nil {
				continue
			}
			app.clients.Store(remoteAdrs.String(), remoteAdrs)
			app.Broadcast(buffer[:n])
		}

	}
}

func (app *Application) Broadcast(message []byte) {
	app.clients.Range(func(key, value any) bool {
		addrs := value.(net.Addr)

		if _, err := app.udpConn.WriteTo(message, addrs); err != nil {
			app.logger.Fatalf("Error broadcasting to %s: %w", addrs.String(), addrs)
			app.clients.Delete(key)
		}
		return true
	})
}
