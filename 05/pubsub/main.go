package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

var PROJECT_ID = os.Getenv("PROJECT_ID")
var PUBSUB_TOPIC = os.Getenv("PUBSUB_TOPIC")

func main() {
	if PROJECT_ID == "" || PUBSUB_TOPIC == "" {
		log.Fatal("Please set PROJECT_ID or/and PUBSUB_TOPIC")
	}
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	client, err := pubsub.NewClient(ctx, PROJECT_ID)
	if err != nil {
		log.Fatal(err)
	}

	topic := client.Topic(PUBSUB_TOPIC)
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("{message: \"Hello World\"}"),
	})

	msgID, err := res.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("messageID: %v", msgID)

	_, err = fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
