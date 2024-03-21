package nft

import (
	"net/http"

	"Nft-Go/gateway/internal/logic/nft"
	"Nft-Go/gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllActivityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := nft.NewGetAllActivityLogic(r.Context(), svcCtx)
		resp, err := l.GetAllActivity()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
