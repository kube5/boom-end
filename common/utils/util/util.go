package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	rand2 "math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func ReadFileAsString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ProjectDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(filename))))
}

func Retry(retryInterval time.Duration, retryTime int, work func() (notRetry bool, err error)) error {
	notRetry, err := work()
	if notRetry {
		return err
	}

	timer := time.NewTicker(retryInterval)
	defer timer.Stop()
	totalTime := 1
	for {
		<-timer.C
		notRetry, err = work()
		if notRetry {
			return err
		}

		totalTime++
		if totalTime >= retryTime {
			return err
		}
	}
}

func GenerateCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			result = Substr2(str, 0, 3) + "****" + Substr2(str, 7, 11)
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)

			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}

func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}

func HandleTokenDecimal(value string, decimal uint64, eth2Wei bool) (string, error) {
	if value == "" {
		return "0", nil
	}
	amount, ok := new(big.Float).SetString(value)
	if !ok {
		return "", fmt.Errorf("invalid big float value: %s", value)
	}

	bigDecimal := big.NewFloat(math.Pow10(int(decimal)))

	if eth2Wei {
		amount = amount.Mul(amount, bigDecimal)
	} else {
		amount = amount.Quo(amount, bigDecimal)
	}

	result, _ := amount.Float64()

	return strconv.FormatFloat(result, 'f', -1, 64), nil
}

func Contains(array []string, item string) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

func StringToBigInt(value string) (*big.Int, error) {
	if value == "" {
		return big.NewInt(0), nil
	}

	bigInt, ok := big.NewInt(0).SetString(value, 10)
	if !ok {
		return nil, fmt.Errorf("invalid big int value: %s", value)
	}

	return bigInt, nil
}

func Min(x *big.Int, y *big.Int) *big.Int {
	if x.Cmp(y) < 0 {
		return x
	}
	return y

}

func IsZero(a float64) bool {
	return math.Abs(a) < 1e-10
}

func GenerateRandomCode(length int) string {
	rng := rand2.New(rand2.NewSource(time.Now().UnixNano()))

	code := ""
	for i := 0; i < length; i++ {
		randomDigit := rng.Intn(10)
		code += fmt.Sprintf("%d", randomDigit)
	}

	return code
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}
