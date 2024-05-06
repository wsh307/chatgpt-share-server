package service

import (
	"backend/modules/chatgpt/model"

	"github.com/cool-team-official/cool-admin-go/cool"
)

type ChatgptUserService struct {
	*cool.Service
}

func NewChatgptUserService() *ChatgptUserService {
	return &ChatgptUserService{
		&cool.Service{
			Model: model.NewChatgptUser(),
			PageQueryOp: &cool.QueryOp{
				FieldEQ:      []string{"usertoken"},
				KeyWordField: []string{"usertoken"},
			},
		},
	}
}
