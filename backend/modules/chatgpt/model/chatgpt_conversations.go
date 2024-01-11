package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameChatgptConversations = "chatgpt_conversations"

// ChatgptConversations mapped from table <chatgpt_conversations>
type ChatgptConversations struct {
	*cool.Model
	UserToken string `gorm:"column:usertoken;type:varchar(255);not null;index;comment:用户token" json:"userToken"` // 用户token
	ConvId    string `gorm:"column:convid;type:varchar(255);not null;index;comment:会话id" json:"convId"`          // 会话id
	Title     string `gorm:"column:title;type:text;index,priority:1,not null;comment:会话标题" json:"title"`         // 会话标题
	Email     string `gorm:"column:email;type:varchar(255);not null;index;comment:官网账号邮箱" json:"email"`          // 邮箱
	// ChatGPT-Account-ID
	ChatGPTAccountID string `gorm:"column:chatgptaccountid;type:varchar(255);comment:ChatGPT-Account-ID" json:"chatGPTAccountID"`
}

// TableName ChatgptConversations's table name
func (*ChatgptConversations) TableName() string {
	return TableNameChatgptConversations
}

// GroupName ChatgptConversations's table group
func (*ChatgptConversations) GroupName() string {
	return "default"
}

// NewChatgptConversations create a new ChatgptConversations
func NewChatgptConversations() *ChatgptConversations {
	return &ChatgptConversations{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&ChatgptConversations{})
}
