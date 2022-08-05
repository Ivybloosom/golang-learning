package handler

import (
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/logic/user"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/svc"
	"github.com/Ivybloosom/go-zero-template/app/service/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserAddOneHandleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddOneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserAddOneHandleLogic(r.Context(), svcCtx)
		resp, err := l.UserAddOneHandle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
