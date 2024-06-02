# Firebase Firestore Operations in Go

This project demonstrates basic CRUD (Create, Read, Update, Delete) operations using Firebase Firestore in a Go application.

## Prerequisites

- Go installed on your machine.
- A Firebase project set up with Firestore enabled.
- A Firebase service account key file downloaded from your Firebase project settings.

## Project Structure

```
.
├── main.go
└── path/to/serviceAccountKey.json
```

## Steps to Run the Application

1. **Clone the repository or create a new Go project and add `main.go` file:**

    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. **Install the Firebase Admin SDK:**

    ```sh
    go get firebase.google.com/go
    go get google.golang.org/api/option
    ```

3. **Add your service account key file:**

    Download your service account key file from your Firebase project settings and place it in the specified path (e.g., `path/to/serviceAccountKey.json`).

4. **Update the path to the service account key file in `main.go`:**

    ```go
    sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")
    ```

5. **Run the application:**

    ```sh
    go run main.go
    ```

## Logic and Flow Explanation

### 1. Initializing Firebase

The `initFirebase` function initializes the Firebase app using the provided service account key file.

```go
func initFirebase() {
    sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")
    app, err := firebase.NewApp(context.Background(), nil, sa)
    if err != nil {
        log.Fatalf("error initializing app: %v\n", err)
    }
    client = app
    log.Println("Firebase successfully initialized!")
}
```

### 2. Creating a Document

The `createDocument` function creates a new document in the specified collection with the given data.

```go
func createDocument(collection string, id string, data map[string]interface{}) {
    ctx := context.Background()
    ref := client.NewRef(collection).Child(id)
    if err := ref.Set(ctx, data); err != nil {
        log.Fatalf("error creating document: %v\n", err)
    }
    log.Println("Document successfully created!")
}
```

### 3. Reading a Document

The `readDocument` function reads a document from the specified collection and returns the data.

```go
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
```

### 4. Updating a Document

The `updateDocument` function updates a document in the specified collection with the given data.

```go
func updateDocument(collection string, id string, data map[string]interface{}) {
    ctx := context.Background()
    ref := client.NewRef(collection).Child(id)
    if err := ref.Set(ctx, data); err != nil {
        log.Fatalf("error updating document: %v\n", err)
    }
    log.Println("Document successfully updated!")
}
```

### 5. Deleting a Document

The `deleteDocument` function deletes a document from the specified collection.

```go
func deleteDocument(collection string, id string) {
    ctx := context.Background()
    ref := client.NewRef(collection).Child(id)
    if err := ref.Delete(ctx); err != nil {
        log.Fatalf("error deleting document: %v\n", err)
    }
    log.Println("Document successfully deleted!")
}
```

### Main Function

The `main` function demonstrates the usage of all CRUD operations.

```go
func main() {
    initFirebase()
    createDocument("collectionName", "documentID", map[string]interface{}{"field": "value"})
    data := readDocument("collectionName", "documentID")
    log.Println(data)
    updateDocument("collectionName", "documentID", map[string]interface{}{"field": "new value"})
    deleteDocument("collectionName", "documentID")
}
```

This script initializes Firebase, creates a document, reads it, updates it, and then deletes it, demonstrating the basic operations in a structured flow.