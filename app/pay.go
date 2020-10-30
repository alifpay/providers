package app

import (
	"context"

	"github.com/alifpay/providers/models"
)

func (s Server) pay(ctx context.Context, req models.ReqAPI) (resp models.RespAPI) {

	resp.ID = req.ID

	//here will be logic of payment request
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

		//check payment in your database by request ID, if exists return 107, duplicate payments
		if isExists(req.ID){
			resp.Code = 107
			return
		}
	*/

	//if payment successfully processed and final status return 200 code
	resp.Code = 200
	//if payment is not final status return 202
	//resp.Code = 202
	//if payment is not final status
	//alif processing system will send status request to check status of payment

	//this is your payment id
	resp.ResponseID = "4522145"

	return
}
