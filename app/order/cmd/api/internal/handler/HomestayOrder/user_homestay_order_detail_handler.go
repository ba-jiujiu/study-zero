package HomestayOrder

import (
	"net/http"

	"study-zero/app/order/cmd/api/internal/logic/HomestayOrder"
	"study-zero/app/order/cmd/api/internal/svc"
	"study-zero/app/order/cmd/api/internal/types"
	"study-zero/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// user homestay order detail
func UserHomestayOrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHomestayOrderDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := HomestayOrder.NewUserHomestayOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserHomestayOrderDetail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
