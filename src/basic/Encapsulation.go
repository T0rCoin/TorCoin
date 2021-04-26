package basic

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"io"
	"os"
)

func MD5(str string) string {
	return fmt.Sprintf("%x",md5.Sum([]byte(str)))
}

func SHA256(str string) string {
	return fmt.Sprintf("%x",sha256.Sum256([]byte(str)))
}

func SHA512(str string) string {
	return fmt.Sprintf("%x",sha512.Sum512([]byte(str)))
}

func ECCEncrypt(pt []byte, puk ecies.PublicKey) ([]byte, error) {
	ct, err := ecies.Encrypt(rand.Reader, &puk, pt, nil, nil)
	return ct, err
}

func ECCDecrypt(ct []byte, prk ecies.PrivateKey) ([]byte, error) {
	pt, err := prk.Decrypt(ct, nil, nil)
	return pt, err
}

func getKey() (*ecdsa.PrivateKey, error) {
	prk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return prk, err
	}
	return prk, nil
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func OutputFile(fileBody []byte, fileName string) interface{} {
	var f *os.File
	var err error
	if !checkFileIsExist(fileName) {
		f, err = os.Create(fileName)
		if err != nil {
			return err
		}
	}else {
		return "Can't Open File"
	}
	_, err = io.WriteString(f, string(fileBody))
	if err != nil{
		return err
	}
	return 0
}

func WriteToFile(content, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		defer f.Close()
	}
	return err
}

func TransactionGenerationHash(Initiator, Receiver string, Value, TransactionFee int64)  {
	
}

