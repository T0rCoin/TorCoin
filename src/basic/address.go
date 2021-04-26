package basic

import (
	"TorCoin/src/Init"
	types "TorCoin/type"
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

type GKey struct {
	privateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

func (k GKey) GetPrivKey() []byte {
	d := k.privateKey.D.Bytes()
	b := make([]byte, 0, Init.PrivKeyBytesLen)
	priKey := paddedAppend(Init.PrivKeyBytesLen, b, d) // []bytes type
	return priKey
}

func (k GKey) GetPubKey() []byte {
	pubKey := append(k.PublicKey.X.Bytes(), k.privateKey.Y.Bytes()...) // []bytes type
	return pubKey
}

func (k GKey) Sign(text []byte) (string, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.privateKey, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

func (k GKey) GetAddress() (address string) {
	pub_bytes := k.GetPubKey()
	sha256_h := sha256.New()
	sha256_h.Reset()
	sha256_h.Write(pub_bytes)
	pub_hash_1 := sha256_h.Sum(nil)
	ripemd160_h := ripemd160.New()
	ripemd160_h.Reset()
	ripemd160_h.Write(pub_hash_1)
	pub_hash_2 := ripemd160_h.Sum(nil)
	address = b58checkencode(0x00, pub_hash_2)
	return address
}

func getSign(signature string) (rint, sint big.Int, err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error," + err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode = ", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return
}

func Verify(text []byte, signature string, pubKey *ecdsa.PublicKey) (bool, error) {
	rint, sint, err := getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(pubKey, text, &rint, &sint)
	return result, nil
}

func MakeNewKey(randKey string) (*GKey, error) {
	var err error
	var gkey GKey
	var curve elliptic.Curve
	lenth := len(randKey)
	if lenth < 224/8+8 {
		err = errors.New("RandKey is too short. It mast be longer than 36 bytes.")
		return &gkey, err
	} else if lenth > 521/8+8 {
		curve = elliptic.P521()
	} else if lenth > 384/8+8 {
		curve = elliptic.P384()
	} else if lenth > 256/8+8 {
		curve = elliptic.P256()
	} else if lenth > 224/8+8 {
		curve = elliptic.P224()
	}
	private, err := ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		log.Panic(err)
	}
	gkey = GKey{private, private.PublicKey}
	return &gkey, nil
}

func Address(Secret, WalletID string) (types.Address, error) {
	gkey, err := MakeNewKey(Secret)
	if err != nil {
		return types.Address{} ,err
	}
	privKey := gkey.GetPrivKey()
	pubKey := gkey.GetPubKey()
	address := "0x" + gkey.GetAddress()
	return types.Address{
		WalletID: 	WalletID,
		Index:		WeHash(address),
		PublicKey:  ByteToString(pubKey),
		PrivateKey: ByteToString(privKey),
		Address:    address,
		Amount:     0,
	}, err
}


func GetAddressStatus(Address string) (*types.Address, error) {

	file, err := ioutil.ReadFile("Database/address/" + Address + ".json")

	if err != nil{
		return nil, err
	}

	var JSON types.Address

	json.Unmarshal(file, &JSON)

	return &JSON, err

}