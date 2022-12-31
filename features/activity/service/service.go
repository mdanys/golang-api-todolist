package service

import "golang/features/activity"

type activityService struct {
	qry activity.Repository
}

func New(repo activity.Repository) activity.Service {
	return &activityService{qry: repo}
}

func (as *activityService) GetAll() ([]activity.Core, error) {
	res, err := as.qry.ShowAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *activityService) GetOne(id uint) (activity.Core, error) {
	res, err := as.qry.ShowOne(id)
	if err != nil {
		return activity.Core{}, err
	}

	return res, nil
}

func (as *activityService) Create(data activity.Core) (activity.Core, error) {
	res, err := as.qry.Insert(data)
	if err != nil {
		return activity.Core{}, err
	}

	return res, nil
}

func (as *activityService) Update(data activity.Core, id uint) (activity.Core, error) {
	res, err := as.qry.Edit(data, id)
	if err != nil {
		return activity.Core{}, err
	}

	return res, nil
}

func (as *activityService) Delete(id uint) (activity.Core, error) {
	res, err := as.qry.Remove(id)
	if err != nil {
		return activity.Core{}, err
	}

	return res, nil
}
