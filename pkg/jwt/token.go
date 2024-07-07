package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type Token struct {
	AccessToken  string
	AccessExpire int64
}

const TokenExpireDuration = time.Hour * 24 * 7

// CreateToken 生成Token
func CreateToken(Secret string, userId int) (Token, error) {
	var RespToken Token
	claims := jwt.MapClaims{
		"userId": userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return RespToken, err
	}
	return Token{
		AccessToken:  accessToken,
		AccessExpire: time.Now().Add(TokenExpireDuration).Unix(),
	}, nil
}

//// ParseToken 解析token
//func ParseToken(tokenString string) (*MyClaims, error) {
//	//tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
//	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return Secret, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}
//
//// RefreshToken 更新token的有效时间
//func RefreshToken(tokenString string) error {
//	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return Secret, nil
//	})
//	if err != nil {
//		return err
//	}
//
//	// 更新token的有效时间
//	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
//		jwt.TimeFunc = time.Now
//		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(TokenExpireDuration))
//		claims.RegisteredClaims.NotBefore = jwt.NewNumericDate(time.Now())
//	}
//
//	return nil
//}

// GenToken ⽣生成access token 和 refresh token
//func GenToken(userID int) (aToken, rToken string, err error) {
//	// 创建⼀一个我们⾃自⼰己的声明
//	c := MyClaims{
//		userID, // ⾃自定义字段
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
//			Issuer:    "bluebell",                                 // 签发⼈人
//		},
//	}
//	// 加密并获得完整的编码后的字符串串token
//	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256,
//		c).SignedString(Secret)
//	// refresh token 不不需要存任何⾃自定义数据
//	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
//		ExpiresAt: time.Now().Add(time.Second * 30).Unix(), // 过期时间
//		Issuer:    "bluebell",                              // 签发⼈人
//	}).SignedString(Secret)
//	// 使⽤用指定的secret签名并获得完整的编码后的字符串串token
//	return
//}
