package helper

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/haimait/go-mindoc/define"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

//Encrypt 加密 即使密码一样，每次加密的串也不一样
func Encrypt(password string) string {
	if password == "" {
		return ""
	}
	var err error
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return ""
	}
	return string(hash)
}

// CompareHashAndPassword 验证密码是否正确
func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}
func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

// RFC3339ToNormalTime RFC3339 日期格式标准化
func RFC3339ToNormalTime(rfc3339 string) string {
	if len(rfc3339) < 19 || rfc3339 == "" || !strings.Contains(rfc3339, "T") {
		return rfc3339
	}
	return strings.Split(rfc3339, "T")[0] + " " + strings.Split(rfc3339, "T")[1][:8]
}

func GenerateToken(id int, username string, second int) (string, error) {
	//claims := make(jwt.MapClaims)
	//JWT 规定了7个Registered claims
	//iss (issuer)：签发人
	//exp (expiration time)：过期时间
	//sub (subject)：主题
	//aud (audience)：受众 用户id
	//nbf (Not Before)：生效时间 在此之前不可用
	//iat (Issued At)：签发时间
	//jti (JWT ID)：编号 JWT ID用于标识该JWT

	//作者：pilisiyang
	//链接：https://www.jianshu.com/p/75b865e55568
	//来源：简书
	//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
	uc := define.UserClaim{
		Id:       id,
		Username: username,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second))), //过期时间
			Issuer:    define.JwtKey,                                                           // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// httpRequest .
func httpRequest(url, method string, data, header []byte) ([]byte, error) {
	var err error
	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 处理 header
	if len(header) > 0 {
		headerMap := new(map[string]interface{})
		err = json.Unmarshal(header, headerMap)
		if err != nil {
			return nil, err
		}
		for k, v := range *headerMap {
			if k == "" || v == "" {
				continue
			}
			request.Header.Set(k, v.(string))
		}
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func HttpDelete(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header)
}

func HttpPut(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header)
}

func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header)
}

func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header)
}
