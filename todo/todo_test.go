package todo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/uhey22e/sqlboiler-tutor/models"
	"github.com/uhey22e/sqlboiler-tutor/types"
	"github.com/volatiletech/null"
)

var testApp App

func TestMain(m *testing.M) {
	// connect to db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "sample_db",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// validate whether or not the connection string was correct
	if err := db.Ping(); err != nil {
		panic(err)
	}

	testApp = App{
		db: db,
	}

	os.Exit(m.Run())
}

func TestStore(t *testing.T) {
	// prepare sample insert data
	now := time.Now().UTC()
	todos := []*models.Todo{
		&models.Todo{
			Title:    "Sample ToDo 1",
			DueDate:  types.NullTimeFrom(now.Add(3 * time.Hour)),
			Note:     null.StringFrom("note1..."),
			Finished: false,
		},
		&models.Todo{
			Title:    "Sample ToDo 2",
			DueDate:  types.NullTimeFrom(now.Add(6 * time.Hour)),
			Note:     null.StringFrom("note2..."),
			Finished: false,
		},
		&models.Todo{
			Title:    "Sample ToDo 3",
			Note:     null.StringFrom("note3..."),
			Finished: false,
		},
	}

	// store
	if err := testApp.Store(todos); err != nil {
		t.Fatal(err)
	}

	for _, v := range todos {
		t.Log(fmt.Sprintf("%+v", v))
	}
}

func TestFetchUnfinished(t *testing.T) {
	todos, err := testApp.FetchUnfinished()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range todos {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(v); err != nil {
			t.Fatal(err)
		}
		t.Log(buf.String())
	}
}

func TestFinish(t *testing.T) {
	if err := testApp.Finish([]int64{1, 2}); err != nil {
		t.Fatal(err)
	}
}
