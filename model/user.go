package model

import (
	"encoding/json"
	"time"
)

// UserSession user session
type UserSession struct {
	ID             uint `json:"id"`
	RefreshTokenID uint `json:"refreshTokenID"`
}

// VerifyEmailType verify email type
type VerifyEmailType int

const (
	// VerifyEmailTypeNotVerify not verify
	VerifyEmailTypeNotVerify VerifyEmailType = iota + 1
	// VerifyEmailTypeVerified verified
	VerifyEmailTypeVerified
	// VerifyEmailTypeVerifyButChange but change email
	VerifyEmailTypeVerifyButChange
)

// AccountRole account role
type AccountRole int

const (
	// AccountRoleMember account role member
	AccountRoleMember AccountRole = iota + 1

	// AccountRoleAdmin account role member
	AccountRoleAdmin AccountRole = 100
)

// Member member role
func (ar AccountRole) Member() bool {
	return ar == AccountRoleMember
}

// Admin admin role
func (ar AccountRole) Admin() bool {
	return ar == AccountRoleAdmin
}

// User user model
type User struct {
	Model
	Name                        string          `json:"name"`
	InitialsName                string          `json:"initialsName"`
	Email                       string          `json:"email,omitempty"`
	ChangeToEmail               string          `json:"changeToEmail,omitempty" gorm:"-"`
	PasswordHash                string          `json:"-"`
	ProfilePicture              string          `json:"profilePicture"`
	Bio                         string          `json:"bio,omitempty"`
	FacebookID                  string          `json:"facebookId,omitempty"`
	GoogleID                    string          `json:"googleId,omitempty"`
	AppleID                     string          `json:"apple_id,omitempty"`
	ColorHex                    string          `json:"colorHex"`
	VerifyEmailType             VerifyEmailType `json:"verifyEmailType" gorm:"default:1"`
	NumberOfUnreadNotifications int             `json:"numberOfUnreadNotifications"`
	IsNewAccount                bool            `gorm:"-" json:"-"`
	IsPassword                  bool            `gorm:"-" json:"isPassword"`
	DeactivatedAt               *time.Time      `json:"deactivatedAt,omitempty"`
	DeactivatedReason           string          `json:"deactivatedReason,omitempty"`
	Suggestion                  string          `json:"suggestion,omitempty"`
	Language                    string          `json:"language"`
	Role                        AccountRole     `json:"role" gorm:"default:1"`
}

// MarshalJSON marshall custom json
func (u *User) MarshalJSON() ([]byte, error) {

	if isPassword := (u.PasswordHash != ""); isPassword {
		u.IsPassword = true
	}

	type Alias User
	user := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	return json.Marshal(user)
}
