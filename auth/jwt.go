package auth

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/twinj/uuid"
// )

// type tokenservice struct{}

// func NewToken() *tokenservice {
// 	return &tokenservice{}
// }

// type TokenInterface interface {
// 	CreateToken(username string) (*TokenDetails, error)
// 	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
// }

// //Token implements the TokenInterface
// var _ TokenInterface = &tokenservice{}

// func (t *tokenservice) CreateToken(username string) (*TokenDetails, error) {
// 	td := &TokenDetails{}
// 	td.AtExpires = time.Now().Add(time.Minute * 30).Unix() //expires after 30 min
// 	td.TokenUuid = uuid.NewV4().String()

// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = td.TokenUuid + "++" + username

// 	var err error
// 	//Creating Access Token
// 	atClaims := jwt.MapClaims{}
// 	atClaims["access_uuid"] = td.TokenUuid
// 	atClaims["username"] = username
// 	atClaims["exp"] = td.AtExpires
// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
// 	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}

// 	//Creating Refresh Token
// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = td.TokenUuid + "++" + username

// 	rtClaims := jwt.MapClaims{}
// 	rtClaims["refresh_uuid"] = td.RefreshUuid
// 	rtClaims["username"] = username
// 	rtClaims["exp"] = td.RtExpires
// 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

// 	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return td, nil
// }

// func TokenValid(r *http.Request) error {
// 	token, err := verifyToken(r)
// 	if err != nil {
// 		return err
// 	}
// 	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
// 		return err
// 	}
// 	return nil
// }

// func verifyToken(r *http.Request) (*jwt.Token, error) {
// 	tokenString := extractToken(r)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("ACCESS_SECRET")), nil
// 	})
// 	if err != nil {
// 		return nil, err

// 	}
// 	return token, nil
// }

// //get the token from the request body
// func extractToken(r *http.Request) string {
// 	bearToken := r.Header.Get("Authorization")
// 	strArr := strings.Split(bearToken, " ")
// 	if len(strArr) == 2 {
// 		return strArr[1]
// 	}

// 	return ""
// }

// func extract(token *jwt.Token) (*AccessDetails, error) {

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		accessUuid, ok := claims["access_uuid"].(string)
// 		username, userOk := claims["username"].(string)
// 		if ok == false || userOk == false {
// 			return nil, errors.New("unauthorized")
// 		} else {
// 			return &AccessDetails{
// 				TokenUuid: accessUuid,
// 				Username:  username,
// 			}, nil
// 		}
// 	}
// 	return nil, errors.New("something went wrong")
// }

// func (t *tokenservice) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
// 	token, err := verifyToken(r)

// 	if err != nil {
// 		return nil, err
// 	}
// 	acc, err := extract(token)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return acc, nil
// }
