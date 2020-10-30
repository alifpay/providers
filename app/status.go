package app

import (
	"context"

	"github.com/alifpay/providers/models"
)

func (s Server) status(ctx context.Context, req models.ReqAPI) (resp models.RespAPI) {

	//here will be logic of status request
	//check status of payment by alif request ID

	/*
		//check payment in your database by request ID, if does not exists return 104
		if doesNotExists(req.ID){
			resp.Code = 104
			return
		}
	*/

	//if payment exists and final successful status return 200 code
	resp.Code = 200
	//if payment is not final status return 201
	//resp.Code = 201
	//if payment is not final status
	//alif processing system will send status request to check status of payment

	//this is your payment id
	resp.ResponseID = "4522145"

	return
}
