package tokens

import (
	"errors"
	"net/http"
	"net/textproto"
	"github.com/connext-cs/pub/log"
	"github.com/connext-cs/pub/response"

	jwt "github.com/dgrijalva/jwt-go"
	api "github.com/micro/go-api/proto"
)

func getHeaderValue(r *api.Request, key string) string {
	if r.Header[key] != nil {
		if len(r.Header[key].Values) > 0 {
			return r.Header[key].Values[0]
		}
	}
	return ""
}

func GetTokenValueNew(r *api.Request) (*TokenValue, error) {
	token, err := GetTokenFromRequest(r)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("token.Raw:", token.Raw)
	tokenvalue, err := ReadToken(token.Raw)
	if err != nil {
		log.Error(err.Error())
	}
	if tokenvalue == nil {
		return nil, nil
	}

	log.Info("tokenvalue:", tokenvalue)

	userid := tokenvalue.UserID
	usertype := tokenvalue.UserType
	username := tokenvalue.UserName

	if userid <= 0 {
		if userid != -10000 {
			err = errors.New("userid is error.")
		}
	}

	if usertype == User_None {
		err = errors.New("usertype is out of range.")
	}

	return &TokenValue{
		UserID:   userid,
		UserType: usertype,
		UserName: username,
	}, err
}

func GetTokenFromRequest(req *api.Request) (token *jwt.Token, err error) {
	pair := &api.Pair{}
	pair = req.Header["Authorization"]
	if req.Header["Authorization"] == nil {
		log.Error("no Authorization in request")
		return &jwt.Token{}, errors.New("no Authorization in request")
	} else {
		pair = req.Header["Authorization"]
	}
	values := make([]string, 0)
	if pair.Values == nil {
		log.Error("pair.Values == nil")
		return &jwt.Token{}, errors.New("pair.Values == nil")
	}
	values = pair.Values
	tokenStr := values[0]
	token, err = jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, err := ReadToken(token.Raw)
		if err != nil {
			return []byte(SecretKey), err
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		log.Error("GetTokenFromRequest err:%s", err.Error())
	}
	return token, err
}

func HeaderAddKey(req *api.Request, key, value string) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	strings := make([]string, 0)
	req.Header[key] = &api.Pair{Key: "", Values: strings}
	req.Header[key].Values = append(req.Header[key].Values, value)
}

func ValidateToken(req *api.Request, rsp *api.Response) error {
	//AllowOrigin(rsp)
	if req.Method == "OPTIONS" {
		rsp.StatusCode = http.StatusOK

	} else {
		token, err := GetTokenFromRequest(req)
		if err != nil || !token.Valid {
			rsp.StatusCode = http.StatusUnauthorized
			var outdata []byte
			outdata, _ = response.NewResponse(1, "", "token is invalid.").Pack()
			rsp.Body = string(outdata)
			if err != nil {
				return err
			} else {
				return errors.New("token is invalid.")
			}
		}
	}
	return nil
}
