package app

import (
	"log"
	"net/http"

	"github.com/alifpay/providers/models"
)

//Ping handler
func (s Server) api(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		req  models.ReqAPI
		resp models.RespAPI
	)

	if !parseBody(r, &req) {
		resp.Code = 400
		s.reply(w, resp)
		return
	}

	switch req.Action {
	case "check":
		resp = s.check(ctx, req)
	case "pay":
		resp = s.pay(ctx, req)
	case "status":
		resp = s.status(ctx, req)
	default:
		log.Println("req.Action", req.Action)
		resp.Code = 400
	}

	s.reply(w, resp)
}
