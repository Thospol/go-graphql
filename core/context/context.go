package context

import (
	"net/http"

	con "github.com/gorilla/context"
	"github.com/thospol/go-graphql/model"
	"gorm.io/gorm"
)

const (
	userKey       = "user"
	langKey       = "lang"
	parametersKey = "parameters"
	databaseKey   = "sql"
	errMsgKey     = "errMsg"
)

// SetUser set user
func SetUser(r *http.Request, user *model.UserSession) {
	con.Set(r, userKey, user)
}

// GetUser get user
func GetUser(r *http.Request) (*model.UserSession, bool) {
	userSession := con.Get(r, userKey)
	if userSession == nil {
		return nil, false
	}
	return userSession.(*model.UserSession), true
}

// SetLanguage set message language response
func SetLanguage(r *http.Request, language string) {
	con.Set(r, langKey, language)
}

// GetLanguage get message language response
func GetLanguage(r *http.Request) string {
	language := con.Get(r, langKey)
	return language.(string)
}

// SetParameters set parameters
func SetParameters(r *http.Request, parameters interface{}) {
	con.Set(r, parametersKey, parameters)
}

// GetParameters get parameter
func GetParameters(r *http.Request) (interface{}, bool) {
	parameters := con.Get(r, parametersKey)
	if parameters == nil {
		return nil, false
	}
	return parameters, true
}

// SetDatabase set database
func SetDatabase(r *http.Request, database *gorm.DB) {
	con.Set(r, databaseKey, database)
}

// GetDatabase get database
func GetDatabase(r *http.Request) (*gorm.DB, bool) {
	databaseConnection := con.Get(r, databaseKey)
	if databaseConnection == nil {
		return nil, false
	}
	return databaseConnection.(*gorm.DB), true
}

// SetErrMsg set err message
func SetErrMsg(r *http.Request, errMsg string) {
	con.Set(r, errMsgKey, errMsg)
}

// GetErrMsg get err message
func GetErrMsg(r *http.Request) string {
	errMsg := con.Get(r, errMsgKey)
	if errMsg == nil {
		errMsg = ""
	}

	return errMsg.(string)
}
