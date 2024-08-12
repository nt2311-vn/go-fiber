package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	pb := pocketbase.New()

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		log.Printf("PocketBase is serving on %s\n", ":8080")
		return nil
	})

	if err := pb.Start(); err != nil {
		log.Fatalf("Failed to start PocketBase: %v", err)
	}
}
