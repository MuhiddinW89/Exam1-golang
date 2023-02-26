package controller

import "app/models"


func (c *Controller) CreateBranch(req models.CreateBranch) error{
  err := c.store.Branch().CreateBranch(req)
  if err != nil {
	return  err
}
	return nil
}

func (c *Controller) UpdateBranch(req *models.BranchPrimaryKey, name string) error {
	err := c.store.Branch().UpdateBranch(req, name)
	if err != nil {
		return err
	}
	return err
}


func (c *Controller) DleteBranch(req *models.BranchPrimaryKey) error {
	err := c.store.Branch().DeleteBranch(req)
	if err != nil {
		return err
	}
	return err
}

func (c *Controller) GetByIdBranch(req *models.UserPrimaryKey) (models.Branch, error) {
	branches, err := c.store.Branch().GetByID(req)
	if err != nil {
		return models.Branch{}, err
	}
	return branches, nil
}