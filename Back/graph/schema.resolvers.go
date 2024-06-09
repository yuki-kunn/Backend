package graph

import (
	"context"
	"log"

	"github.com/[username]/todoapp-graphql-go-react/app/graph/model"
	"github.com/[username]/todoapp-graphql-go-react/app/infrastructure"
	"github.com/google/uuid"
)

// 新しいTodoを作成するためのリゾルバ
func (r *mutationResolver) CreateTodo(ctx context.Context, todoInput model.CreateTodoInput) (*model.Todo, error) {
	// 新しいUUIDの生成
	newUUID, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error generating UUID: %v\n", err)
		return nil, err
	}
	// 新しいTodoの作成
	todo := &infrastructure.Todo{
		ID:   newUUID.String(),
		Text: todoInput.Text,
	}

	// DBに新しいTodoを挿入
	_, err = r.DB.Insert(todo)
	if err != nil {
		log.Printf("Error insert todo: %v\n", err)
		return nil, err
	}

	// 新しいTodoを返す
	return model.NewTodo(todo), nil
}

// Todoのステータスを更新するためのリゾルバ
func (r *mutationResolver) UpdateTodoStatus(ctx context.Context, todoID string, done bool) (bool, error) {
	todo := &infrastructure.Todo{
		Done: done,
	}
	// 指定されたIDのTodoのステータスを更新
	_, err := r.DB.ID(todoID).Cols("done").Update(todo)
	if err != nil {
		log.Printf("Error update todo status: %v\n", err)
		return false, err
	}
	// 更新成功を返す
	return true, nil
}

// Todoを削除するためのリゾルバ
func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID string) (bool, error) {
	// 指定されたIDのTodoを削除
	_, err := r.DB.ID(todoID).Delete(&infrastructure.Todo{})
	if err != nil {
		log.Printf("Error delete todo: %v\n", err)
		return false, err
	}
	// 削除成功を返す
	return true, nil
}

// 全てのTodoを取得するためのリゾルバ
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*infrastructure.Todo
	// DBから全てのTodoを取得
	r.DB.AllCols().Find(&todos)
	// 取得したTodoを返す
	return model.NewTodos(todos), nil
}
