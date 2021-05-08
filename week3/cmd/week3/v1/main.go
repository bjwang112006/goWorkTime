package main

import (
	"github.com/gorilla/mux"
	"os"
	http "week3/pkg/server"
	"time"
	"github.com/go-kratos/kratos/v2/log"
	appserver "week3"
	handler "week3/app"
)

func main() {
	logger := log.NewStdLogger(os.Stdout)
	logger = log.With(logger, "caller", log.DefaultCaller, "ts", log.DefaultTimestamp)


	hs := http.NewServer(
		http.Address("127.0.0.1:8000"),
		http.Logger(log.Warn(logger)),
		http.Timeout(1*time.Second),
		http.Network("tcp"),
		)
	router1 := mux.NewRouter()
	router1.HandleFunc("/ws", handler.WsHandler)
	hs.HandlePrefix("/", router1)

	profileHs := http.NewServer(
		http.Address("127.0.0.1:8001"),
		http.Logger(log.Debug(logger)),
		http.Timeout(1*time.Second),
		http.Network("tcp"),
		)
	router2 := mux.NewRouter()
	router2.HandleFunc("/ws", handler.WsDebugHandler)
	profileHs.HandlePrefix("/", router2)

	healthHs := http.NewServer(
		http.Address("127.0.0.1:8002"),
		http.Logger(log.Info(logger)),
		http.Timeout(1*time.Second),
		http.Network("tcp"),
	)
	router3 := mux.NewRouter()
	router3.HandleFunc("/ws", handler.WsHealthHandler)
	healthHs.HandlePrefix("/", router3)


	app := appserver.New(
		appserver.ID("1"),
		appserver.Name("test1"),
		appserver.Version("v1.0.0"),
		appserver.Server(
			*hs,
			*profileHs,
			*healthHs,
			),
	)
	//time.AfterFunc(1000*time.Second, func() {
	//	app.Stop()
	//})
	if err := app.Run(); err != nil {
		logger.Log("msg",err)
	}
}