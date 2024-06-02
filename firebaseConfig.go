package main

import (
	"context"
	"log"
	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var client *db.Client

func initFirebase() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err = app.DatabaseWithURL(ctx, "https://<DATABASE_NAME>.firebaseio.com")
	if err != nil {
		log.Fatalf("error initializing database client: %v\n", err)
	}
}