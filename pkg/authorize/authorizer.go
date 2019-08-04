package authorize

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	secret string
}

func NewClient(
	secret string,
) Client {
	return Client{
		secret: secret,
	}
}

func (c Client) IsAuthorized(ctx context.Context) (bool, error) {
	ctxMD, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errors.New("no metadata found in context")
	}

	authorization := ctxMD.Get("authorization")
	if len(authorization) == 0 {
		return false, errors.New("token not found in context")
	}

	tokenStr := strings.TrimPrefix(authorization[0], "Bearer ")

	// decode authorization
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(c.secret), nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["fvp.Do"]; !ok {
			return false, errors.New("missing fvp.Do claim")
		}
	}
	return true, nil
}
