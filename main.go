package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var client *firebase.App

func initFirebase() {
	// Path ke file kunci layanan
	sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")

	// Inisialisasi aplikasi Firebase
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client = app
	log.Println("Firebase successfully initialized!")
}

func createDocument(collection string, id string, data map[string]interface{}) {
	ctx := context.Background()
	ref := client.NewRef(collection).Child(id)
	if err := ref.Set(ctx, data); err != nil {
		log.Fatalf("error creating document: %v\n", err)
	}
	log.Println("Document successfully created!")
}

func readDocument(collection string, id string) map[string]interface{} {
	ctx := context.Background()
	ref := client.NewRef(collection).Child(id)
	var data map[string]interface{}
	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalf("error reading document: %v\n", err)
	}
	log.Println("Document successfully read!")
	return data
}

func updateDocument(collection string, id string, data map[string]interface{}) {
	ctx := context.Background()
	ref := client.NewRef(collection).Child(id)
	if err := ref.Set(ctx, data); err != nil {
		log.Fatalf("error updating document: %v\n", err)
	}
	log.Println("Document successfully updated!")
}

func deleteDocument(collection string, id string) {
	ctx := context.Background()
	ref := client.NewRef(collection).Child(id)
	if err := ref.Delete(ctx); err != nil {
		log.Fatalf("error deleting document: %v\n", err)
	}
	log.Println("Document successfully deleted!")
}

func main() {
	initFirebase()
	// Example usage
	createDocument("collectionName", "documentID", map[string]interface{}{"field": "value"})
	data := readDocument("collectionName", "documentID")
	log.Println(data)
	updateDocument("collectionName", "documentID", map[string]interface{}{"field": "value"})
	deleteDocument("collectionName", "documentID")
}