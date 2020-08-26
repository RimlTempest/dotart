package utils

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns hashed password.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), errors.Wrap(err, "failed to generate from password")
}

// CheckHashOfPassword ハッシュ値とパスワード文字列を比較
func CheckHashOfPassword(hashed, password string) bool {
	fmt.Println("HASHD : " + hashed + " PASSWORD : " + password)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	fmt.Println(err)
	if err != nil {
		return false
	}
	return true
}
