package crypto

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"unsafe"
)

func PasswordHash(password string) (string, error) {
	if (len([]rune(password)) >= 73) {
		return "", errors.New("73文字以上のパスワードは指定できません")
	}

	pp := (*[]byte)(unsafe.Pointer(&password))
	hash, err := bcrypt.GenerateFromPassword(*pp, bcrypt.DefaultCost)

	if (err != nil) {
		return "", err
	} else {
		return string(hash), err
	}
}

func PasswordVerify(hashedPass, password string) error {
	php := (*[]byte)(unsafe.Pointer(&hashedPass))
	pp := (*[]byte)(unsafe.Pointer(&password))

    return bcrypt.CompareHashAndPassword(*php, *pp)
}

