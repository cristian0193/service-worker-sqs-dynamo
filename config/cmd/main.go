package main

import (
	"os"
	"os/signal"
	"service-worker-sqs-dynamo/config/cmd/builder"
	cases "service-worker-sqs-dynamo/core/usecases/events"
	repository "service-worker-sqs-dynamo/dataproviders/repository/events"
	"service-worker-sqs-dynamo/dataproviders/server"
	"service-worker-sqs-dynamo/entrypoints/controllers/events"
	"syscall"
)

func main() {

	// logger is initialized
	logger := builder.NewLogger()
	logger.Info("Starting service-worker-sqs-dynamo ...")
	defer builder.Sync(logger)

	// config is initialized
	config, err := builder.LoadConfig()
	if err != nil {
		logger.Fatalf("error in LoadConfig : %v", err)
	}

	// session aws is initialized
	session, err := builder.NewSession(config)
	if err != nil {
		logger.Fatalf("error in Session : %v", err)
	}

	// db is initialized
	db, err := builder.NewDB(config)
	if err != nil {
		logger.Fatalf("error in RDS : %v", err)
	}

	// repositories are initialized
	eventsRepository := repository.NewEventsRepository(db)

	// usecases are initialized
	eventsUseCases := cases.NewEventsUseCases(eventsRepository)

	// controllers are initialized
	eventsController := events.NewEventsController(eventsUseCases)

	// sqs is initialized
	sqs, err := builder.NewSQS(logger, config, session, eventsRepository)
	if err != nil {
		logger.Fatalf("error in SQS : %v", err)
	}

	// processor is initialized
	processor, err := builder.NewProcessor(logger, sqs)
	if err != nil {
		logger.Fatalf("error in Processor : %v", err)
	}
	go processor.Start()

	// server is initialized
	srv := server.NewServer(config.Port, eventsController)
	if err = srv.Start(); err != nil {
		logger.Fatalf("error Starting Server: %v", err)
	}

	// Graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	sig := <-sigQuit

	logger.Infof("Shutting down server with signal [%s] ...", sig.String())
	if err = sqs.Close(); err != nil {
		logger.Error("error Closing Consumer SQS: %v", err)
	}

	if err = srv.Stop(); err != nil {
		logger.Error("error Stopping Server: %v", err)
	}

	logger.Info("service-worker-sqs-dynamo ended")

}