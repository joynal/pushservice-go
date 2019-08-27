package main

import (
  "context"
  "fmt"
  "github.com/mongodb/mongo-go-driver/bson"
  "github.com/mongodb/mongo-go-driver/bson/primitive"
  "log"
  "os"
  "pushservice-go/utils"
  "time"
)

func main() {
  utils.LoadConfigs()

  dbUrl := os.Getenv("MONGODB_URL")
  dbName := os.Getenv("DB_NAME")

  // Db connection stuff
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  ctx = context.WithValue(ctx, utils.DbURL, dbUrl)
  db, err := utils.ConfigDB(ctx, dbName)
  if err != nil {
    log.Fatalf("database configuration failed: %v", err)
  }

  fmt.Println("Connected to MongoDB!")

  update := bson.M{
    "status":    "done",
    "updatedAt": time.Now(),
    "totalSent": 222,
  }

  pushId, err := primitive.ObjectIDFromHex("5d4aaf166a1d0b53687cde16")

  res, err := db.Collection("pushes").UpdateOne(ctx, bson.M{"_id": pushId}, bson.M{"$set": update})

  if err != nil {
    log.Fatalln("update err:", err)
  }

  fmt.Println("res: ", res)
}
