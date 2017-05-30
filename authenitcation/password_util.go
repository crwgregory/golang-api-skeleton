package authentication

import (
	"encoding/base64"
	"golang.org/x/crypto/sha3"
)

type PasswordHasher struct {
	Key []byte
}

func (p *PasswordHasher) HashPassword(password string) (string, error) {

	if len(password) < 8 {
		return "", errors.PasswordToShort{}
	}

	if len(p.Key) == 0 {
		key, err := config.GetSecretKey()
		if err != nil {
			return "", err
		}
		p.Key = key
	}

	shakeHash := sha3.NewShake256()
	s := make([]byte, 48)
	shakeHash.Write(p.Key)
	shakeHash.Write([]byte(password))
	shakeHash.Read(s)

	hashed := base64.StdEncoding.EncodeToString(s)

	return hashed, nil
}
