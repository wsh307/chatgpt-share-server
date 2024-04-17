package service

import (
	"backend/modules/chatgpt/model"

	"github.com/cool-team-official/cool-admin-go/cool"

	"strings"

	"github.com/google/uuid"
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

// 批量添加卡密
func (s *ChatgptUserService) ServiceAdd(ctx g.Ctx, req *cool.AddReq) (data interface{}, err error) {
	m := g.DB().Model(s.Model)
	expireTime := g.RequestFromCtx(ctx).Get("expireTime").Time()
	isPlus := g.RequestFromCtx(ctx).Get("isPlus").Bool()
	batchCount := g.RequestFromCtx(ctx).Get("batch").Int()
	remark := g.RequestFromCtx(ctx).Get("remark").String()

	g.Log().Info(ctx, "Request Info", g.Map{
		"expireTime": expireTime,
		"isPlus":     isPlus,
		"batchCount": batchCount,
		"remark":     remark,
	})

	tokens := make([]string, 0, batchCount) // 用于收集userToken
	for i := 0; i < batchCount; i++ {
		userToken := uuid.New().String() // 生成新的 UUID
		userMap := map[string]interface{}{
			"userToken":  userToken,
			"expireTime": expireTime,
			"isPlus":     isPlus,
			"remark":     remark,
		}

		// 插入数据
		_, err := m.Data(userMap).Insert()
		if err != nil {
			return nil, err // 如果插入过程中发生错误，立即返回错误
		}

		tokens = append(tokens, userToken) // 收集生成的userToken
	}

	// 将tokens数组中的元素通过换行符连接成一个字符串
	data = strings.Join(tokens, "\n")
	return data, nil
}
