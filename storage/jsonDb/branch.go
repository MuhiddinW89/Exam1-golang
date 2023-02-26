package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
}

func NewbranchRepo(fileName string) *branchRepo {
	return &branchRepo{
		fileName: fileName,
	}
}


func (b *branchRepo) Read() ([]models.Branch, error) {
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return []models.Branch{}, err
	}

	var branches []models.Branch
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return []models.Branch{}, err
	}
	return branches, nil
}


func (b *branchRepo) CreateBranch(req models.CreateBranch) error{
	branches, err := b.Read()
	if err != nil {
		return err
	}

	uuid := uuid.New().String()
	branches = append(branches, models.Branch{
		Id:      uuid,
		Name:    req.Name,
	})

	body, err := json.MarshalIndent(branches, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return  err
	}
	return nil
}



func (b *branchRepo) DeleteBranch(req *models.BranchPrimaryKey) error {
	branches, err := b.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range branches {
		if v.Id == req.Id {
			branches = append(branches[:i], branches[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no user with this id")
	}

	body, err := json.MarshalIndent(branches, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (b *branchRepo) UpdateBranch(req *models.BranchPrimaryKey, name string) error {
	branches, err := b.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range branches {
		if v.Id == req.Id  {
			branches[i].Name = name
			flag = false
		}
	}

	if flag {
		return errors.New("There is no branch with this id")
	}

	body, err := json.MarshalIndent(branches, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (b *branchRepo) GetByID(req *models.UserPrimaryKey) (models.Branch, error) {
	branches, err := b.Read()
	if err != nil {
		return models.Branch{}, err
	}		
	for _, v := range branches {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.Branch{}, errors.New("There is no branch with this id")
}