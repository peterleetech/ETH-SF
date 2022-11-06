package main

import (
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"veric-backend-mvp/logic/config"
	"veric-backend-mvp/logic/db"
	"veric-backend-mvp/logic/events"
	"veric-backend-mvp/logic/http"
	"veric-backend-mvp/logic/http/open_api"
	"veric-backend-mvp/logic/log"
)

func main() {
	db.InitDB()

	eventsListen, err := events.NewEvents(config.Get().Rpc)
	if err != nil {
		panic(err)
	}
	defer eventsListen.Close()

	open_api.SetEvent(eventsListen)

	go func() {
		err := http.StartAndServe()
		if err != nil {
			if err == netHttp.ErrServerClosed {
				log.GetLogger().Info("http stopped")
				return
			} else {
				panic(err)
			}
		}
	}()
	defer http.Stop()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.GetLogger().Info("gracefully shutdown...")
}
