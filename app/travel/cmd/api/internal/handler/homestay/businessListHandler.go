package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study-zero/app/travel/cmd/api/internal/logic/homestay"
	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
)

// boss all homestay room
func BusinessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BusinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestay.NewBusinessListLogic(r.Context(), svcCtx)
		resp, err := l.BusinessList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
