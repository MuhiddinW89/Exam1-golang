package models

type Branch struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type BranchPrimaryKey struct{
	Id string `json:"id"`
}

type CreateBranch struct {
	Name string `json:"name"`
}
