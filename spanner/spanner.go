package main

import (
	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

// Entity for Spanner record
type Entity struct {
	BookId    string
	Name      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	initialize()

	bookId := uuid.New()

	entity := &Entity{
		BookId:    bookId.String(),
		Name:      "test name",
		Author:    "test author",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	spannerClient, err := createClient(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err2 := write(ctx, spannerClient, entity)
	if err2 != nil {
		fmt.Println(err2)
	}

	closeSpanner()
}

func initialize() {
	fmt.Println("Start spanner...")
}

func createClient(ctx context.Context) (*spanner.Client, error) {
	gcpProject := os.Getenv("GCP_PROJECT_ID")
	spannerInstance := os.Getenv("SPANNER_INSTANCE_ID")
	spannerDb := os.Getenv("SPANNER_DB_ID")

	if gcpProject == "" {
		err := errors.New("failed get env gcpProject")
		return nil, err
	}
	if spannerInstance == "" {
		err := errors.New("failed get env spannerInstance")
		return nil, err
	}
	if spannerDb == "" {
		err := errors.New("failed get env spannerDb")
		return nil, err
	}

	myDB := fmt.Sprintf("projects/%s/instances/%s/databases/%s", gcpProject, spannerInstance, spannerDb)

	client, err := spanner.NewClient(ctx, myDB)
	if err != nil {
		fmt.Println("Failed Spanner Client")
		return nil, err
	}
	fmt.Printf("Mydb is %s\n", myDB)
	return client, nil
}

func write(ctx context.Context, client *spanner.Client, entity *Entity) error {
	fmt.Println(entity.BookId)

	BooksColumn := []string{"BookId", "Name", "Author", "CreatedAt", "UpdatedAt"}
	m := []*spanner.Mutation{
		spanner.InsertOrUpdate("Books", BooksColumn, []interface{}{entity.BookId, entity.Name, entity.Author, entity.CreatedAt, entity.UpdatedAt}),
	}
	_, err := client.Apply(ctx, m)

	return err
}

func closeSpanner() {
	fmt.Println("End spanner...")
}
