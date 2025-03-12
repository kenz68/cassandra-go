package db

import (
	"github.com/gocql/gocql"
	"github.com/kenz68/cassandra-go/src/pkg/model"
)

type TodoRepository interface {
	Save(todo *model.Todo) (*model.Todo, error)
	GetByID(id string) (*model.Todo, error)
}

type todoRepository struct {
	session *gocql.Session
}

// GetByID implements TodoRepository.
func (t *todoRepository) GetByID(id string) (*model.Todo, error) {
	var todo model.Todo
	var query string = "SELECT id, title, content FROM todo WHERE id = ?"
	if err := t.session.Query(query, id).Scan(&todo.Id, &todo.Title, &todo.Content); err != nil {
		return nil, err
	}
	return &todo, nil
}

// Save implements TodoRepository.
func (t *todoRepository) Save(todo *model.Todo) (*model.Todo, error) {
	var query string = "INSERT INTO todo (id, title, content) VALUES (?, ?, ?)"
	if err := t.session.Query(query, todo.Id, todo.Title, todo.Content).Exec(); err != nil {
		return nil, err
	}
	return todo, nil
}

func NewTodoRepository(session *gocql.Session) TodoRepository {
	return &todoRepository{session: session}
}
