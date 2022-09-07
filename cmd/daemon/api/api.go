package api

import (
	"github.com/xbt573/gotask/pkg/types"
)

type Api struct{}

func (this *Api) List(empty *string, reply *types.ListResponse) error {
	tasks, err := Database.List()
	if err != nil {
		*reply = types.ListResponse{
			Ok:      false,
			Message: err.Error(),
		}
		return err
	}

	*reply = types.ListResponse{
		Ok:    true,
		Tasks: tasks,
	}

	return nil
}

func (this *Api) Info(req *types.InfoRequest, reply *types.InfoResponse) error {
	task, err := Database.Info(req.Id)
	if err != nil {
		*reply = types.InfoResponse{
			Ok:      false,
			Message: err.Error(),
		}

		return err
	}

	*reply = types.InfoResponse{
		Ok:   true,
		Task: task,
	}
	return nil
}

func (this *Api) Delete(req *types.DeleteRequest, reply *types.DeleteResponse) error {
	task, err := Database.Delete(req.Id)
	if err != nil {
		*reply = types.DeleteResponse{
			Ok:      false,
			Message: err.Error(),
		}

		return err
	}

	*reply = types.DeleteResponse{
		Ok:   true,
		Task: task,
	}

	return nil
}

func (this *Api) New(req *types.NewRequest, reply *types.NewResponse) error {
	task, err := Database.New(req.Name, req.Description, req.Priority)
	if err != nil {
		*reply = types.NewResponse{
			Ok:      false,
			Message: err.Error(),
		}

		return err
	}

	*reply = types.NewResponse{
		Ok:   true,
		Task: task,
	}

	return nil
}
