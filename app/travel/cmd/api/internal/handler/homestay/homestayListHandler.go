package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study-zero/app/travel/cmd/api/internal/logic/homestay"
	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
)

// homestay room list
func HomestayListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestay.NewHomestayListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
