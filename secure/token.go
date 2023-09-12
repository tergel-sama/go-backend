package secure

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	TokenSeconds        = 1800
	RefreshTokenSeconds = 43200
)

var (
	ErrIncorrectJWTStructure = errors.New("incorrect JWT structure")
)

// JWT payload structure
type Payload struct {
	jwt.StandardClaims
	UUID         string                 `json:"uuid"`
	Data         map[string]interface{} `json:"data"`
	RefreshToken bool                   `json:"refreshtoken"`
}

// Types can have associated functions which are equivalent of C# extensions.
func (pl *Payload) FromJSON(s string) error {
	return json.Unmarshal([]byte(s), pl)
}

// Encapsulates the JSON to IO writer logic, read more at
// https://pkg.go.dev/encoding/json#Encoder.Encode
func (pl *Payload) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(pl)
}

// Generates a JWT using RS512 (RSA public key cryptography based) hashing algorithm.
// Includes custom claims in addition to the standard JWT claims.
func IssueToken(uuid string, refreshToken bool, custom []string, suuid string, data map[string]interface{}, kp *RsaKey) (string, error) {
	ep := make(map[string]int)
	for i, v := range custom {
		ep[v] = i
	}

	claims := Payload{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   "tanasoft-auth-service",
		},
		uuid,
		data,
		refreshToken,
	}

	if refreshToken {
		claims.StandardClaims.ExpiresAt = time.Now().Unix() + RefreshTokenSeconds
	} else {
		claims.StandardClaims.ExpiresAt = time.Now().Unix() + TokenSeconds
	}

	method := jwt.GetSigningMethod(kp.Algo)
	token := jwt.NewWithClaims(method, claims)

	stoken, err := token.SignedString(kp.GetK1())
	if err != nil {
		return "", err
	}

	return stoken, nil
}

// Computes signature, verifies the JWT authenticity and returns the JWT paylaod.
func VerifyToken(token string, kp *RsaKey) (string, error) {
	parts := strings.Split(token, ".")

	if len(parts) < 3 {
		return "", ErrIncorrectJWTStructure
	}
	method := jwt.GetSigningMethod(kp.Algo)

	err := method.Verify(strings.Join(parts[0:2], "."), parts[2], kp.GetK2())
	if err != nil {
		return "", err
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	return string(payload), nil
}
