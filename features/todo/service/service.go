package service

import "golang/features/todo"

type todoService struct {
	qry todo.Repository
}

func New(repo todo.Repository) todo.Service {
	return &todoService{qry: repo}
}

func (ts *todoService) GetAll(query string) ([]todo.Core, error) {
	res, err := ts.qry.ShowAll(query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ts *todoService) GetOne(id uint) (todo.Core, error) {
	res, err := ts.qry.ShowOne(id)
	if err != nil {
		return todo.Core{}, err
	}

	return res, nil
}

func (ts *todoService) Create(data todo.Core) (todo.Core, error) {
	data.IsActive = true
	data.Priority = "very-high"
	res, err := ts.qry.Insert(data)
	if err != nil {
		return todo.Core{}, err
	}

	return res, nil
}

func (ts *todoService) Update(data todo.Core, id uint) (todo.Core, error) {
	res, err := ts.qry.Edit(data, id)
	if err != nil {
		return todo.Core{}, err
	}

	return res, nil
}

func (ts *todoService) Delete(id uint) (todo.Core, error) {
	res, err := ts.qry.Remove(id)
	if err != nil {
		return todo.Core{}, err
	}

	return res, nil
}
