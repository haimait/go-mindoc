package helper

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	secRand "crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/astaxie/beego/validation"
	"github.com/shopspring/decimal"
)

const (
	space = " "
)

func init() {
	// 增加可跳过的验证
	validation.CanSkipFuncs["Match"] = struct{}{}
	validation.CanSkipFuncs["Min"] = struct{}{}
	validation.CanSkipFuncs["Max"] = struct{}{}
}
func VersionString2Int64(data string, align int) int64 {
	versionList := strings.Split(data, ".")
	for k, v := range versionList {
		if len(v) >= align {
			continue
		}
		versionList[k] = fmt.Sprintf("%0*d%s", align-len(v), 0, v)
	}
	versionStr := strings.Join(versionList, "")
	return String2Int64(versionStr)
}

func Byte2String(data []byte) string {
	return string(data[:])
}

func String2Byte(data string) []byte {
	return []byte(data)
}

func Int2String(data int) string {
	return strconv.Itoa(data)
}

func Int322String(data int32) string {
	return strconv.FormatInt(int64(data), 10)
}

func Int642String(data int64) string {
	return strconv.FormatInt(data, 10)
}

func Float642String(data float64) string {
	return strconv.FormatFloat(data, 'E', -1, 64)
}

func String2Int(data string) int {
	k, _ := strconv.Atoi(data)
	return k
}

func String2Int64(data string) int64 {
	k, _ := strconv.ParseInt(data, 10, 64)
	return k
}

func String2Float64(data string) float64 {
	k, _ := strconv.ParseFloat(data, 64)
	return k
}

func String2Float32(data string) float64 {
	k, _ := strconv.ParseFloat(data, 32)
	return k
}

func String2Decimal(data string) decimal.Decimal {
	d, _ := decimal.NewFromString(data)
	return d
}

func DecimalAdd(base string, amount float64) string {
	return String2Decimal(base).Add(decimal.NewFromFloat(amount)).String()
}

func DecimalSub(base string, amount float64) string {
	return String2Decimal(base).Sub(decimal.NewFromFloat(amount)).String()
}

func DecimalCmp(base string, amount float64) int {
	return String2Decimal(base).Cmp(decimal.NewFromFloat(amount))
}

func DecimalPrecisionCeil(base decimal.Decimal, precision int) decimal.Decimal {
	p := int32(precision)
	return base.Shift(p).Ceil().Shift(p * -1)
}

func DecimalPrecisionFloor(base decimal.Decimal, precision int) decimal.Decimal {
	p := int32(precision)
	return base.Shift(p).Floor().Shift(p * -1)
}

func ValidateEmail(email string) bool {
	pattern := `^[0-9a-zA-Z][_.0-9a-zA-Z]{0,31}@([0-9a-zA-Z][0-9a-zA-Z-]{0,30}[0-9a-zA-Z]\.){1,4}[a-zA-Z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func ValidatePhone(phone string) bool {
	reg := regexp.MustCompile(`^[0-9]{5,15}$`)
	if reg.FindString(phone) == "" {
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	if len(password) > 64 || len(password) < 6 {
		return false
	}

	reg := regexp.MustCompile("[a-z]+")
	if reg.FindString(password) == "" {
		return false
	}

	reg = regexp.MustCompile("[A-Z]+")
	if reg.FindString(password) == "" {
		return false
	}

	reg = regexp.MustCompile("[0-9]+")
	if reg.FindString(password) == "" {
		return false
	}

	reg = regexp.MustCompile("[^a-zA-Z0-9]+")
	if reg.FindString(password) == "" {
		return false
	}

	return true
}

func Sha512(data []byte) string {
	sOb := sha512.New()
	sOb.Write(data)
	r := sOb.Sum(nil)
	return hex.EncodeToString(r)
}

func Md5v2(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	md5Data := md5Ctx.Sum(nil)
	return hex.EncodeToString(md5Data)
}

func MakeSecRand512() []byte {
	return MakeSecRand(64)
}

func MakeSecRand(length int) []byte {
	k := make([]byte, length)
	if _, err := secRand.Read(k); err != nil {
		if _, err := secRand.Read(k); err != nil {
			panic("Make secure rand failed")
		}
	}
	return k
}

func DoubleMd5WithSalt(data []byte, salt string) string {
	if salt == "" {
		salt = RandomStr(32, "Aa0")
	}

	first := Md5v2(data)
	second := Md5v2([]byte(first[:8] + salt + first[8:]))
	result := second[:8] + salt + second[8:]
	return result
}

func VerifyDoubleMd5(data, hash string) bool {
	if len(hash) != 64 {
		return false
	}

	salt := hash[8 : 8+32]
	first := Md5v2([]byte(data))
	second := Md5v2([]byte(first[:8] + salt + first[8:]))
	result := second[:8] + salt + second[8:]
	return subtle.ConstantTimeCompare([]byte(result), []byte(hash)) == 1
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

func HashHmac(data, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)

	hash := mac.Sum(nil)

	// -- 二进制转为十六进制
	return fmt.Sprintf("%x", hash)
}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

func RandomStr(randLength int, randType string) (result string) {
	num := "0123456789"
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result = ""

	b := bytes.Buffer{}
	if strings.Contains(randType, "0") {
		b.WriteString(num)
	}
	if strings.Contains(randType, "a") {
		b.WriteString(lower)
	}
	if strings.Contains(randType, "A") {
		b.WriteString(upper)
	}

	str := b.String()
	strLen := len(str)
	if strLen == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < randLength; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	result = b.String()
	return
}

func RandomRangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}

// CurrrentTime获取系统当前时间，时间格式是：yyyy-mm-dd h24:mm:ss
func CurrrentTime() string {
	return time.Now().Format(DateTimeFormat)
}

// CurrentDate获取系统当前日前，日期格式是： yyyy-mm-dd
func CurrentDate() string {
	return time.Now().Format(DateFormat)
}

func ValidStructData(s interface{}) (err error) {
	valid := validation.Validation{RequiredFirst: true}
	b, _ := valid.Valid(s)
	if !b {
		for _, err := range valid.Errors {
			log.Print(err.Key, err.Value, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func ConvertStructData(src interface{}, dst ...interface{}) error {
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}

	for _, v := range dst {
		if err := json.Unmarshal(jsonStr, v); err != nil {
			return err
		}
	}

	return nil
}

func NumberSplitByComma(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}

// IsEmpty returns true if the string is empty
func IsEmpty(text string) bool {
	return len(text) == 0
}

// IsNotEmpty returns true if the string is not empty
func IsNotEmpty(text string) bool {
	return !IsEmpty(text)
}

// IsBlank returns true if the string is blank (all whitespace)
func IsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

// IsNotBlank returns true if the string is not blank
func IsNotBlank(text string) bool {
	return !IsBlank(text)
}

// Left justifies the text to the left
func Left(text string, size int) string {
	spaces := size - Length(text)
	if spaces <= 0 {
		return text
	}

	var buffer bytes.Buffer
	buffer.WriteString(text)

	for i := 0; i < spaces; i++ {
		buffer.WriteString(space)
	}
	return buffer.String()
}

// Right justifies the text to the right
func Right(text string, size int) string {
	spaces := size - Length(text)
	if spaces <= 0 {
		return text
	}

	var buffer bytes.Buffer
	for i := 0; i < spaces; i++ {
		buffer.WriteString(space)
	}

	buffer.WriteString(text)
	return buffer.String()
}

// Center justifies the text in the center
func Center(text string, size int) string {
	left := Right(text, (Length(text)+size)/2)
	return Left(left, size)
}

// IsMark determines whether the rune is a marker
func IsMark(r rune) bool {
	return unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Mc, r)
}

// Length counts the input while respecting UTF8 encoding and combined characters
func Length(text string) int {
	textRunes := []rune(text)
	textRunesLength := len(textRunes)

	sum, i, j := 0, 0, 0
	for i < textRunesLength && j < textRunesLength {
		j = i + 1
		for j < textRunesLength && IsMark(textRunes[j]) {
			j++
		}
		sum++
		i = j
	}
	return sum
}

// Reverse reverses the input while respecting UTF8 encoding and combined characters
func Reverse(text string) string {
	textRunes := []rune(text)
	textRunesLength := len(textRunes)
	if textRunesLength <= 1 {
		return text
	}

	i, j := 0, 0
	for i < textRunesLength && j < textRunesLength {
		j = i + 1
		for j < textRunesLength && IsMark(textRunes[j]) {
			j++
		}

		if IsMark(textRunes[j-1]) {
			// Reverses Combined Characters
			reverse(textRunes[i:j], j-i)
		}

		i = j
	}

	// Reverses the entire array
	reverse(textRunes, textRunesLength)

	return string(textRunes)
}

func reverse(runes []rune, length int) {
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

// ToString converts a value to string.
func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(value.(int64)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(value.(uint64)), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value.(float64)), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value.(bool))
	default:
		return fmt.Sprintf("%+v", value)
	}
}

func MtRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func Explode(str string, split string) []string {
	return strings.Split(str, split)
}

//Contain 判断obj是否在target中，target支持的类型array,slice,map
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}
