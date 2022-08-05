package user

import (
	"context"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/svc"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddOneHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddOneHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddOneHandleLogic {
	return &UserAddOneHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddOneHandleLogic) UserAddOneHandle(req *types.UserAddOneReq) (resp *types.UserAddOneResp, err error) {
	// todo: add your logic here and delete this line

	return
}
