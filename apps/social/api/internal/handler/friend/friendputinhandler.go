package friend

import (
	"net/http"

	"github.com/hanzug/easy-chat/apps/social/api/internal/logic/friend"
	"github.com/hanzug/easy-chat/apps/social/api/internal/svc"
	"github.com/hanzug/easy-chat/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendPutInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendPutInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := friend.NewFriendPutInLogic(r.Context(), svcCtx)
		resp, err := l.FriendPutIn(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}