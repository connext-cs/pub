package micro

import (
	"context"
	"errors"
	"fmt"
	"github.com/connext-cs/pub/tokens"
	"strings"

	"github.com/astaxie/beego/logs"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/micro/go-micro/metadata"
)

// Strips 'Bearer ' prefix from bearer token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

//GetTokenFromRequest ...
func GetTokenFromRequest(mda metadata.Metadata) (token *jwt.Token, err error) {
	if tokenStr, ok := mda["Authorization"]; ok {
		var err error
		tokenStr, err = stripBearerPrefixFromTokenString(tokenStr)
		if err != nil {
			return nil, err
		}
		return jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			fmt.Println("token.Raw:", token.Raw)
			tokenvalue, err := tokens.ReadToken(token.Raw)
			if err != nil {
				return []byte(tokens.SecretKey), err
			}
			mda[tokens.USERIDHEADER] = fmt.Sprintf("%d", tokenvalue.UserID)
			// mda[tokens.USERTYPEHEADER] = &api.Pair{Key: tokens.USERTYPEHEADER, Values: []string{fmt.Sprintf("%d", tokenvalue.UserType)}}
			return []byte(tokens.SecretKey), nil
		})
	} else {
		return nil, errors.New("UnAuthorization")
	}
}

//ValidateToken ... 验证Token
func ValidateToken(ctx context.Context) (bool, error) {
	if mda, ok := metadata.FromContext(ctx); ok {
		logs.Info(mda)
		logs.Info(mda["Authorization"])
		_, err := GetTokenFromRequest(mda)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	} else {
		err := errors.New("metadata is error:")
		return false, err
	}
	return true, nil
}
