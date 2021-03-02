package middlewares

import (
	"encoding/json"
	"net/http"

	"strings"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"github.com/thospol/go-graphql/core/config"
	"github.com/thospol/go-graphql/core/context"
	"github.com/thospol/go-graphql/core/jwt"
	"github.com/thospol/go-graphql/model"
)

const (
	authHeader = "Authorization"
	bearer     = "Bearer "
)

// RequireAuthentication require authentication
func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get(authHeader)
		tokenString := strings.Replace(accessToken, bearer, "", 1)
		if tokenString == "" {
			logrus.Error("[RequireAuthentication] token error: ", config.RR.Internal.Unauthorized.Error())
			render.Status(r, config.RR.Internal.Unauthorized.HTTPStatusCode())
			render.JSON(w, r, config.RR.Internal.Unauthorized.WithLocale(r))
			return
		}

		userSession, err := generateUserSessionFromJwtToken(tokenString, true)
		if err != nil {
			logrus.Error("[RequireAuthentication] generate user session from JWT Token: ", err)
			render.Status(r, config.RR.Internal.Unauthorized.HTTPStatusCode())
			render.JSON(w, r, config.RR.Internal.Unauthorized.WithLocale(r))
			return
		}

		context.SetUser(r, userSession)
		next.ServeHTTP(w, r)
	})
}

func generateUserSessionFromJwtToken(token string, onlyValid bool) (*model.UserSession, error) {
	claims, err := jwt.Parsed(token, onlyValid)
	if err != nil {
		return nil, err
	}

	userIDInterface := claims["sub"]
	var userID uint
	if userIDInterface != nil {
		err = json.Unmarshal([]byte(userIDInterface.(string)), &userID)
		if err != nil {
			return nil, err
		}
	}

	refreshTokenIDInterface := claims["refreshTokenID"]
	var rfTokenID uint
	if refreshTokenIDInterface != nil {
		err = json.Unmarshal([]byte(refreshTokenIDInterface.(string)), &rfTokenID)
		if err != nil {
			return nil, err
		}
	}

	userSession := &model.UserSession{
		ID:             userID,
		RefreshTokenID: rfTokenID,
	}

	return userSession, nil
}
