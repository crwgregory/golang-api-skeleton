package authentication

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
	"time"
)

type Payload struct {
	UserID     int    `json:"user_id"`
	UserGroups []int  `json:"user_groups"`
	EntityIDs  []int  `json:"entity_ids"`
	Expires    string `json:"exp"`
}

type header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

func ParseJWT(jwt string) (payload *Payload, err error) {
	parts := strings.Split(jwt, ".")

	headerBytes, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return
	}
	payloadBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return
	}
	signature := parts[2]

	header := new(header)
	payload = new(Payload)

	json.Unmarshal(headerBytes, &header)
	json.Unmarshal(payloadBytes, &payload)

	encodedHeader, encodedPayload := encodeHeaderAndPayload(header, payload)
	encodedSignature, err := getSignature(encodedHeader, encodedPayload)

	if signature != encodedSignature {
		return nil, new(errors.JWTSignatureMismatch)
	}
	i, err := strconv.ParseInt(payload.Expires, 10, 64)
	if err != nil {
		return nil, err
	}

	if i < time.Now().Unix() {
		return nil, new(errors.JWTTokenExpiredError)
	}
	return
}

func GetJWT(userID int, groupIDs, entityIDs []int) (jwt string, err error) {

	header := header{
		Algorithm: "SHAKE256",
		Type:      "JWT",
	}

	payload := Payload{
		UserID:     userID,
		UserGroups: groupIDs,
		EntityIDs:  entityIDs,
		Expires:    strconv.Itoa(int(time.Now().Add(config.JWTLifeTime()).Unix())),
	}

	encodedHeader, encodedPayload := encodeHeaderAndPayload(header, payload)

	encodedSignature, err := getSignature(encodedHeader, encodedPayload)
	if err != nil {
		return
	}
	jwt = encodedHeader + "." + encodedPayload + "." + encodedSignature
	return
}

func getSignature(encodedHeader, encodedPayload string) (signature string, err error) {
	key, err := config.GetSecretKey()
	if err != nil {
		return
	}
	shakeHash := sha3.NewShake256()
	s := make([]byte, 32)
	shakeHash.Write(key)
	shakeHash.Write([]byte(encodedHeader + "." + encodedPayload))
	shakeHash.Read(s)
	return base64.StdEncoding.EncodeToString(s), nil
}

func encodeHeaderAndPayload(header, payload interface{}) (encodedHeader, encodedPayload string) {
	h, _ := json.Marshal(header)
	p, _ := json.Marshal(payload)
	encodedHeader = base64.StdEncoding.EncodeToString(h)
	encodedPayload = base64.StdEncoding.EncodeToString(p)
	return
}
