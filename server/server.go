package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type MailtoGeneratorServer struct {
	firestoreDb *firestore.Client
	router      http.Handler
	config      Config
}

func NewMailtoGeneratorServer(ctx context.Context, config Config) (*MailtoGeneratorServer, error) {
	server := &MailtoGeneratorServer{
		config: config,
	}

	//TODO think this has to be done here to actually link an active db connection to its required service
	//Loading the db here initializes the service with an empty db connection
	//TODO figure out if theres a better way to do this?
	if value := os.Getenv("replace with firestore emulator host"); value != "" {
		fmt.Println("using firestore emulator: ", value)
	}

	fmt.Println(config.GCPProjectId)
	conf := &firebase.Config{ProjectID: config.GCPProjectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		fmt.Println("error making new firebase app: ", err)
		return nil, err
	}

	fmt.Println("making firestore connection...")
	database, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("error making firestore connection: ", err)
		return nil, err
	}
	server.firestoreDb = database

	server.loadRoutes()

	return server, nil
}

func (u *MailtoGeneratorServer) StartMailtoGeneratorServer(ctx context.Context) error {
	fmt.Println("starting mailto server...")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", u.config.ServerPort),
		Handler: u.router,
	}

	defer func() {
		if err := u.firestoreDb.Close(); err != nil {
			fmt.Println("failed to close firestore", err)
		}
	}()

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
