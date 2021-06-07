package commands

import (
	"fmt"
	"github.com/CyganFx/ArdanLabs-Service/business/auth"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

// GenToken generates a JWT for the specified user.
func GenToken() {
	privatePEM, err := ioutil.ReadFile("C:\\Users\\Думан\\go\\src\\github.com\\CyganFx\\se1903service\\private.pem")
	if err != nil {
		log.Fatalln(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		log.Fatalln(err)
	}

	// The call to retrieve a user requires an Admin role by the caller.
	claims := struct {
		jwt.StandardClaims
		Roles []string `json:"roles"`
	}{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "service project",
			Subject:   "123456789",
			ExpiresAt: time.Now().Add(8760 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Roles: []string{auth.RoleAdmin},
	}

	method := jwt.GetSigningMethod("RS256")
	tkn := jwt.NewWithClaims(method, claims)
	tkn.Header["kid"] = "54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"
	str, err := tkn.SignedString(privateKey)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("-----BEGIN TOKEN-----\n%s\n-----END TOKEN-----\n", str)
}
