package user

import (
	"context"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/svc"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/types"
	"github.com/Ivybloosom/go-zero-template/app/service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRetrieveOneHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRetrieveOneHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRetrieveOneHandleLogic {
	return &UserRetrieveOneHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRetrieveOneHandleLogic) UserRetrieveOneHandle(req *types.UserRetrieveOneReq) (resp *types.UserRetrieveOneResp, err error) {
	data := make(map[string]string, 0)

	// 连接表 user
	userModel := model.NewUserModel(l.svcCtx.Config.Mongodb.Uri, l.svcCtx.Config.Mongodb.Db, "user")

	// 查找
	user, _ := userModel.FindOne(l.ctx, req.Id)

	// 创建返回数据
	data["id"] = user.ID.String()
	data["name"] = user.Name

	// 返回
	return &types.UserRetrieveOneResp{
		Code: 200,
		Msg:  "success",
		User: data,
	}, nil
}
