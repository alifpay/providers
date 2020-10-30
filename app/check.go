package app

import (
	"context"

	"github.com/alifpay/providers/models"
)

func (s Server) check(ctx context.Context, req models.ReqAPI) (resp models.RespAPI) {

	resp.ID = req.ID

	//here will be logic of check request
	//if account is not exist return failed code

	/*
		//validation failed return 400 code
		if validationFailed(req.Account) {
			resp.Code = 400
			return
		}
		//does not exists return 404 code
		if doesNotExists(req.Account) {
			resp.Code = 404
			return
		}
	*/
	//if account exists return 302 code
	resp.Code = 302

	//this is optional
	//if you can return the account information your client should see
	//this is format of information \\n new line
	resp.InfoForClient = "Баланс: 50.30 смн\\nСрок оплаты: 25.10.2020"

	return
}
