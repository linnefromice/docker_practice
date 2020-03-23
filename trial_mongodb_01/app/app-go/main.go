package main

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    defer c.Disconnect(ctx)
    err = c.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println("connection error:", err)
    } else {
        fmt.Println("connection success:")
    }
}