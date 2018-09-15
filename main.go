package main

//go:generate sqlboiler --wipe psql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/uhey22e/sqlboiler-tutor/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// App is settings
type App struct {
	db boil.ContextExecutor
}

func main() {
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

	app := &App{
		db: db,
	}

	// prepare insert data
	now := time.Now().UTC()
	todos := []*models.Todo{
		&models.Todo{
			Title:     "Sample ToDo 1",
			DueDate:   null.TimeFrom(now.Add(3 * time.Hour)),
			Note:      null.StringFrom("note..."),
			Finished:  false,
			CreatedAt: now,
		},
		&models.Todo{
			Title:     "Sample ToDo 2",
			DueDate:   null.TimeFrom(now.Add(3 * time.Hour)),
			Note:      null.StringFrom("note..."),
			Finished:  false,
			CreatedAt: now,
		},
	}

	// store
	_ = todos
	// app.Store(todos)

	//	fetched, err := app.FetchUnfinished()
	//	if err != nil {
	//		panic(err)
	//	}
	//	for _, v := range fetched {
	//		fmt.Println(v)
	//	}

	if err := app.Finish([]int64{12, 13}); err != nil {
		panic(err)
	}
}

// Store stores a one or more todos into DB
func (a *App) Store(obj []*models.Todo) error {
	// insert for each todo
	for _, v := range obj {
		if err := v.Insert(context.Background(), a.db, boil.Infer()); err != nil {
			return err
		}
	}
	return nil
}

// FetchUnfinished fetches unfinished ToDos
func (a *App) FetchUnfinished() ([]*models.Todo, error) {
	queries := []qm.QueryMod{
		qm.Where(models.TodoColumns.Finished+"=?", false),
	}

	counts, err := models.Todos(queries...).Count(context.Background(), a.db)
	if err != nil {
		return nil, err
	} else if counts == 0 {
		return make([]*models.Todo, 0), nil
	}

	// fetch!
	return models.Todos(queries...).All(context.Background(), a.db)
}

// Finish update one or more records selected by ID
func (a *App) Finish(ids []int64) error {
	// Set updating columns
	updCols := map[string]interface{}{
		models.TodoColumns.Finished: true,
	}

	// WhereIn method needs to pass a slice of interface{}
	targetIDs := make([]interface{}, len(ids))
	for i, v := range ids {
		targetIDs[i] = v
	}
	query := qm.WhereIn(models.TodoColumns.ID+" IN ?", targetIDs...)

	// update!
	_, err := models.Todos(query).UpdateAll(context.Background(), a.db, updCols)

	return err
}
