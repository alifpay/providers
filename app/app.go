package app

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/alifpay/providers/models"
	jsoniter "github.com/json-iterator/go"
)

func (s *Server) routers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(s.cfg.Version+"/api", s.mw(s.api))
	return mux
}

//http Response writes in json format
func (s Server) reply(w http.ResponseWriter, r models.RespAPI) {
	b, err := jsoniter.Marshal(r)
	if err != nil {
		log.Println("jsoniter.Marshal", err)
		http.Error(w, "внутренняя ошибка, попробуйте позже", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(b)
}

//authorization of requests
func (s *Server) mw(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			log.Println("warning: Authorization wrong value")
			s.reply(w, models.RespAPI{Code: 401})
			return
		}

		bt, err := base64.StdEncoding.DecodeString(r.Header.Get("Authorization"))
		if err != nil {
			log.Println("Authorization base64.StdEncoding.DecodeString", err)
			s.reply(w, models.RespAPI{Code: 401})
			return
		}

		//username and password must be changed and can be dynamic and stored securely
		u := models.User{UserName: "testUser", Password: "Test@PAss4"}

		if u.UserName+":"+u.Password != string(bt) {
			log.Println("Authorization failed", string(bt))
			s.reply(w, models.RespAPI{Code: 401})
			return
		}

		next.ServeHTTP(w, r)
	}
}

//parse body of http Request
func parseBody(r *http.Request, req interface{}) bool {
	if r.Body == nil {
		return false
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil.ReadAll", err)
		return false
	}
	err = jsoniter.Unmarshal(b, req)
	if err != nil {
		log.Println("jsoniter.Unmarshal", err)
		return false
	}
	return true
}
