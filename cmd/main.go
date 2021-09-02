package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"uploader/cmd/router"
	"uploader/pkg/config"
	mylog "uploader/pkg/log"

	"github.com/gorilla/handlers"
)

const (
	logFileName = "./log/uploder_%s.log"
)

var log = mylog.Log

func main() {

	//Initialize config
	{
		config.Init("./configs/config.ini")
	}

	//Initalize logger
	{
		logFileName := strings.Replace(logFileName, "%s", time.Now().Format("2006_01_02___15_04_05"), 1)
		mylog.Init(logFileName, config.GetLogLevel())
	}

	PORT := config.LoadHTTPConfigs()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	var router = router.Init()

	headers := handlers.AllowedHeaders([]string{"*"})
	methods := handlers.AllowedMethods([]string{"*"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Info("Listining on port:", PORT)

	srv := &http.Server{
		Addr: ":" + PORT,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handlers.CORS(headers, methods, origins)(router),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.ErrorDepth("Error on running server", 1, "7001", "Error starting server ", err, srv)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	log.Info("\n\n\n <---------\tshutting down\t---------------->\n\n\n->")
	os.Exit(0)
}
