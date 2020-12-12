package main

import (
	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"math/rand"
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

	rand.Seed(time.Now().UnixNano())
	name := uuid.New()
	author := uuid.New()

	entity := &Entity{
		BookId:    bookId.String(),
		Name:      name.String()[:5],
		Author:    author.String()[:5],
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

	err = read(ctx, spannerClient)
	if err != nil {
		fmt.Println(err)
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

// Simple mutation
func writeWithMutation(ctx context.Context, client *spanner.Client) error {
	// Insert (table name, column name, values)
	m1 := spanner.Insert("Users",
		[]string{"name", "email"},
		[]interface{}{"alice", "a@example.com"})

	// Insert (table name, map{column name: value})
	m2 := spanner.InsertMap("Users", map[string]interface{}{
		"name":  "alice",
		"email": "a@example.com",
	})

	// Use Struct
	type User struct{ Name, Email string }
	u := User{Name: "alice", Email: "a@example.com"}
	m3, err := spanner.InsertStruct("Users", u)

	if err != nil {
		return err
	}

	//Apply list of mutation
	_, err = client.Apply(ctx, []*spanner.Mutation{m1, m2, m3})

	if err != nil {
		return err
	}

	return nil
}

// Use ReadWriteTransaction()
// The record is got before write
func writeWithMutationByReadWriteTransaction(ctx context.Context, client *spanner.Client, txn *spanner.ReadWriteTransaction) error {
	var balance int64
	row, err := txn.ReadRow(ctx, "Accounts", spanner.Key{"alice"}, []string{"balance"})
	if err != nil {
		// This function will be called again if this is an IsAborted error.
		return err
	}
	if err := row.Column(0, &balance); err != nil {
		return err
	}

	if balance <= 10 {
		return errors.New("insufficient funds in account")
	}
	balance -= 10
	m := spanner.Update("Accounts", []string{"user", "balance"}, []interface{}{"alice", balance})
	err = txn.BufferWrite([]*spanner.Mutation{m})
	if err != nil {
		return nil
	}

	// The buffered mutation will be committed.  If the commit
	// fails with an IsAborted error, this function will be called
	// again.
	return nil
}

func read(ctx context.Context, client *spanner.Client) error {
	//row, err := client.Single().ReadRow(ctx, "Accounts", spanner.Key{"alice"}, []string{"balance"})

	//iter := client.Single().Read(ctx, "Accounts", keyset1, columns)
	//key := spanner.AllKeys()
	//fmt.Println(key)

	//---------
	// About KeySet
	//---------
	// In this case, all records will be returned.
	//keySet := spanner.AllKeys()

	// In this case, only record whose uuid is same as following value will be returned.
	//keySet := spanner.Key{"66b09000-4071-46a9-a0e8-166ed2c341bb"}

	// In this case, records whose uuid begins with any character from "1" to "5" (not "6") will be returned.
	keySet := spanner.KeyRange{
		Start: spanner.Key{"1"},
		End:   spanner.Key{"6"},
		Kind:  spanner.ClosedClosed,
	}

	// - []string{a, b, c}
	//   -> value of "a", "b", "c" will be returned.
	iter := client.Single().Read(ctx, "Books", keySet,
		[]string{"BookId", "Name", "Author"})
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return err
		}
		var BookId, Name, Author string
		if err := row.Columns(&BookId, &Name, &Author); err != nil {
			return err
		}
		fmt.Println(fmt.Sprintf("%s %s %s\n", BookId, Name, Author))
	}
}

func closeSpanner() {
	fmt.Println("End spanner...")
}
