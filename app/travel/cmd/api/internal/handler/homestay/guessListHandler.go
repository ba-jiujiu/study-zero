package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"study-zero/app/travel/cmd/api/internal/logic/homestay"
	"study-zero/app/travel/cmd/api/internal/svc"
	"study-zero/app/travel/cmd/api/internal/types"
)

// guess homestay room
func GuessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GuessListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestay.NewGuessListLogic(r.Context(), svcCtx)
		resp, err := l.GuessList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
