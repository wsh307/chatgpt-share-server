package admin

import (
	"backend/modules/chatgpt/service"

	"github.com/cool-team-official/cool-admin-go/cool"
)

type ChatgptConversationsController struct {
	*cool.Controller
}

func init() {
	var chatgpt_conversations_controller = &ChatgptConversationsController{
		&cool.Controller{
			Prefix:  "/admin/chatgpt/conversations",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewChatgptConversationsService(),
		},
	}
	// 注册路由
	cool.RegisterController(chatgpt_conversations_controller)
}
