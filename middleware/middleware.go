package middleware

import (
	"fmt"
	"net/http"
	"github.com/connext-cs/pub/log"
	"github.com/connext-cs/pub/tokens"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//AllowOrigin... 解决跨域问题属性设置
func AllowOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type,Authorization")
	w.Header().Set("Content-Type", "application/json")
}

//FilterOptions 拦截OPTIONS方法
func FilterOptions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	AllowOrigin(w)
	if r.Method == "OPTIONS" {
		log.Info("0001 r.Method == OPTIONS ")
	} else {
		next(w, r)
	}
}

func GetTokenFromRequest(r *http.Request) (token *jwt.Token, err error) {
	token, err = request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			// log.Info("token.Raw %s", token.Raw)
			// log.Info("token.Method:%s", token.Method)
			// log.Info("token.Header:%s", token.Header)
			// log.Info("token.Claims:%s", token.Claims)
			// log.Info("token.Signature:%s", token.Signature)
			// log.Info("token.Raw:%s", token.Raw)
			// log.Info("token.Method:%s", token.Method)
			// log.Info("token.Header:%s", token.Header)
			// log.Info("token.Claims:%s", token.Claims)
			// log.Info("token.Signature:%s", token.Signature)
			tokenvalue, err := tokens.ReadToken(token.Raw)
			// log.Info("begin read token")
			if err != nil {
				return []byte(tokens.SecretKey), err
			}
			// log.Info("token read form Etcd success.")
			r.Header.Add(tokens.USERIDHEADER, fmt.Sprintf("%d", tokenvalue.UserID))
			r.Header.Add(tokens.USERTYPEHEADER, fmt.Sprintf("%d", tokenvalue.UserType))
			// log.Info(tokens.USERIDHEADER, fmt.Sprintf("%d", tokenvalue.UserID))
			// log.Info(tokens.USERTYPEHEADER, fmt.Sprintf("%d", tokenvalue.UserType))
			return []byte(tokens.SecretKey), nil
		})
	if err != nil {
		log.Info("GetTokenFromRequest err:%s", err.Error())
	}
	return token, err
}

//验证Token
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := GetTokenFromRequest(r)
	if err == nil {
		if token.Valid {
			log.Info("ValidateToken, token.Valid")
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
			log.Info("Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		log.Info("Unauthorized access to this resource, " + err.Error())
	}
}

// InValidateToken ... 删除token
func InValidateToken(r *http.Request) error {
	token, err := GetTokenFromRequest(r)
	if err != nil {
		return err
	}
	log.Info("token.Raw:", token.Raw)
	err = tokens.DeleteToken(token.Raw)
	if err != nil {
		return err
	}
	return nil
}
