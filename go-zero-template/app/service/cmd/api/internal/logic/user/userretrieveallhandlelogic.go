package user

import (
	"context"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/svc"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRetrieveAllHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRetrieveAllHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRetrieveAllHandleLogic {
	return &UserRetrieveAllHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRetrieveAllHandleLogic) UserRetrieveAllHandle(req *types.UserRetrieveAllReq) (resp *types.UserRetrieveAllResp, err error) {
	// todo: add your logic here and delete this line

	return
}
