package HomestayOrder

import (
	"net/http"

	"study-zero/pkg/result"

	"study-zero/app/order/cmd/api/internal/logic/HomestayOrder"
	"study-zero/app/order/cmd/api/internal/svc"
	"study-zero/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// create homestay order
func CreateHomestayOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateHomestayOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := HomestayOrder.NewCreateHomestayOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateHomestayOrder(&req)
		result.HttpResult(r, w, resp, err)
	}
}
