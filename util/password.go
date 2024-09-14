package util

import (
	"github.com/JPratama7/util/convert"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(convert.UnsafeBytes(password), 14)
	if err != nil {
		return "", err
	}
	return convert.UnsafeString(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword(convert.UnsafeBytes(hashedPassword), convert.UnsafeBytes(password))
	return err == nil
}
