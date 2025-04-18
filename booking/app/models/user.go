package models

import (
	"fmt"
	"regexp"

	"github.com/wiselike/revel"
)

type User struct {
	UserID             int
	Name               string
	Username, Password string
	HashedPassword     []byte
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var userRegex = regexp.MustCompile(`^\w*$`)

func (u *User) Validate(v *revel.Validation) {
	v.Check(u.Username,
		revel.Required{},
		revel.MaxSize{Max: 15},
		revel.MinSize{Min: 4},
		revel.Match{Regexp: userRegex},
	)

	ValidatePassword(v, u.Password).
		Key("user.Password")

	v.Check(u.Name,
		revel.Required{},
		revel.MaxSize{Max: 100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{Max: 15},
		revel.MinSize{Min: 5},
	)
}
