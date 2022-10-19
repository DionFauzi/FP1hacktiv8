package models

import "errors"

type Todo struct {
	ID          int    `json:"id"`
	ToList      string `json:"tolist"`
	Description string `json:"description"`
}

var dataTodos = []Todo{
	{ID: 1, ToList: "Belanja", Description: "Belanja Bulanan"},
	{ID: 2, ToList: "Seminar", Description: "Seminar nasional di kampus"},
	{ID: 3, ToList: "dokter hewan", Description: "Memeriksa kucing"},
}

func GetAllTodos() []Todo {
	return dataTodos
}

func GetTodoById(id int) (*Todo, error) {
	for i, data := range dataTodos {
		if data.ID == id {
			return &dataTodos[i], nil
		}
	}

	return nil, errors.New("Data not found")
}

func AddNewTodo(newTodo *Todo) {
	data := &dataTodos
	newId := (*data)[len(*data)-1].ID + 1
	newTodo.ID = newId
	dataTodos = append(dataTodos, *newTodo)
}

func DeleteTodoById(id int) (*Todo, error) {
	for i, data := range dataTodos {
		if data.ID == id {
			newData := &dataTodos

			deletedData := (*newData)[i]

			(*newData)[i] = (*newData)[len(*newData)-1]
			*newData = (*newData)[:len(*newData)-1]

			return &deletedData, nil
		}
	}

	return nil, errors.New("Data cannot be deleted")
}

func UpdateTodoById(id int, updateTodo *Todo) (*Todo, error) {
	for i, data := range dataTodos {
		if data.ID == id {
			newData := &dataTodos

			oldData := (*newData)[i]
			updateTodo.ID = oldData.ID
			(*newData)[i] = *updateTodo

			return &oldData, nil
		}
	}

	return nil, errors.New("Data not found")
}
