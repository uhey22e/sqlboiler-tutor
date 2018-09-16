package todo

import (
	"context"

	"github.com/uhey22e/sqlboiler-tutor/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// App is settings
type App struct {
	db boil.ContextExecutor
}

// Store stores a one or more todos into DB
func (a *App) Store(todos models.TodoSlice) error {
	// insert for each todo
	return todos.Insert(context.Background(), a.db, boil.Infer())
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
