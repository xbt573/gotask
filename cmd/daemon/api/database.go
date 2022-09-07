package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/xbt573/gotask/pkg/types"
)

type database struct {
	baseUrl string
}

var Database = database{}
var lock = sync.Mutex{}

func (this *database) Init(baseUrl string) error {
	err := os.Mkdir(baseUrl, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	_, err = os.Stat(fmt.Sprintf("%v/tasks.json", baseUrl))
	if os.IsNotExist(err) {
		file, err := os.Create(fmt.Sprintf("%v/tasks.json", baseUrl))
		if err != nil {
			return err
		}

		_, err = file.Write([]byte("[]"))
		if err != nil {
			return err
		}
	}

	this.baseUrl = baseUrl
	return nil
}

func (this *database) List() ([]types.Task, error) {
	tasks, err := this.read()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (this *database) New(name string, description string, priority int) (types.Task, error) {
	task := types.Task{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Priority:    priority,
	}

	tasks, err := this.read()
	if err != nil {
		return task, err
	}

	tasks = append(tasks, task)
	err = this.write(tasks)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (this *database) Delete(id uuid.UUID) (types.Task, error) {
	tasks, err := this.read()
	if err != nil {
		return types.Task{}, err
	}

	for key, task := range tasks {
		if task.Id != id {
			continue
		}

		tasks = append(tasks[:key], tasks[key+1:]...)
		err := this.write(tasks)
		if err != nil {
			return types.Task{}, err
		}

		return task, nil
	}

	return types.Task{}, errors.New("Not found")
}

func (this *database) Info(id uuid.UUID) (types.Task, error) {
	tasks, err := this.read()
	if err != nil {
		return types.Task{}, err
	}

	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}

	return types.Task{}, errors.New("Not found")
}

func (this *database) read() ([]types.Task, error) {
	tasks := []types.Task{}

	data, err := os.ReadFile(fmt.Sprintf("%v/tasks.json", this.baseUrl))
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (this *database) write(tasks []types.Task) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Create(fmt.Sprintf("%v/tasks.json", this.baseUrl))
	if err != nil {
		return err
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return err
}
