package project

import (
	"context"
	"github.com/kfchen81/beego/vanilla"
	"teamdo/business/user"
	m_project "teamdo/models/project"
	"time"
)

type Project struct {
	vanilla.EntityBase
	Id				int
	Name    		string
	CreateAt		time.Time
	ManagerId		int

	Members			[]*user.Member
	Tunnel			[]*Tunnel
}


func NewProjectForModel(ctx context.Context, dbModel *m_project.Project) *Project {
	instance := new(Project)
	instance.Ctx = ctx
	instance.Id = dbModel.Id
	instance.Name = dbModel.Name
	instance.ManagerId = dbModel.ManagerId
	return instance
}