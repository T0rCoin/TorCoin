package basic

import (
	types "TorCoin/type"
	RandX "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"math"
	"math/big"
	"math/rand"
	mathrand "math/rand"
	"os"
)

const (
	saltMinLen = 8
	saltMaxLen = 32
	iter       = 1000
	keyLen     = 32
)

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := RandX.Int(RandX.Reader, big.NewInt(max+1+i64Min))
		return result.Int64() - i64Min
	} else {
		result, _ := RandX.Int(RandX.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func RandLow(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 31
		b[i] = types.BasicRandomChars[arc]
	}
	return b
}

func RandomFactor(min, max int64) string {

	var slice []string

	for i := 0; i < 10; i++ {
		slice = append(slice,string(RandLow(int(RangeRand(min, max)))))
	}

	return slice[RangeRand(0, int64(len(slice) - 1))]
}

func LoadMnemonic(Path string) ([]string,error) {
	f, err :=  os.Open(Path)
	if err != nil{
		return []string{},err
	}
	s, err := io.ReadAll(f)
	if err != nil {
		return []string{},err
	}
	var Mnemonic []string
	err = json.Unmarshal(s, &Mnemonic)
	if err != nil {
		return []string{},err
	}
	return Mnemonic,nil
}

func RandMnemonic(Max int, language string) ([]string,error) {

	Mnemonic, err := LoadMnemonic(fmt.Sprintf("bips/bip39/%v.json", language))
	if err != nil {
		return []string{},err
	}
	var Slice []string
	for i := 0; i < Max + 1; i++ {
		Slice = append(Slice,Mnemonic[RangeRand(1, int64(len(Mnemonic)) - 1 )])
	}
	return Slice, nil
}

func WeHash(str string) string {

	if str == "" {

		return ""

	}

	Salt := fmt.Sprintf("%x", MD5(str))

	Password := fmt.Sprintf("%x", SHA512(str))

	Password = fmt.Sprintf("%x",SHA256(fmt.Sprintf("%v%v",Password,Salt)))

	return fmt.Sprintf("%x",MD5(Password))

}

func NewWalletID() string {
	return fmt.Sprintf("%x",SHA256(RandomFactor(0x4141,0x4141)))
}

func PassWordHash(Password string) string {
	return fmt.Sprintf("%x",SHA256(Password))
}

func WalletSecret(PasswordHash string) string {
	return SHA512(fmt.Sprintf("%v%v%v",MD5(RandomFactor(0,10)), PasswordHash, RandomFactor(0,40)))
}

func WalletKeyPairsIndex(WalletID, PasswordHash, Secret string) string {

	D1 := MD5(fmt.Sprintf("%v%v",WalletID,PasswordHash))
	D2 := MD5(fmt.Sprintf("%v%v",WalletID,Secret))
	D3 := MD5(fmt.Sprintf("%v%v",PasswordHash,Secret))

	return SHA256(fmt.Sprintf("%v%v%v",D1,D2,D3))
	
}

func EncryptPwd(pwd string) (encrypt string, err error) {
	salt, err := randSalt()
	if err != nil {
		return
	}
	en := encryptPwdWithSalt([]byte(pwd), salt)
	en = append(en, salt...)
	encrypt = base64.StdEncoding.EncodeToString(en)
	return
}

func randSalt() ([]byte, error) {
	salt := make([]byte, mathrand.Intn(saltMaxLen-saltMinLen)+saltMinLen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func encryptPwdWithSalt(pwd, salt []byte) (pwdEn []byte) {
	pwd = append(pwd, salt...)
	pwdEn = pbkdf2.Key(pwd, salt, iter, keyLen, sha256.New)
	return
}
