package basic

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"log"
	"math/big"
)

func ByteToString(b []byte) (s string) {
	s = ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%02X", b[i])
	}
	return s
}

func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

func b58encode(b []byte) (s string) {
	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	x := new(big.Int).SetBytes(b)
	// Initialize
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""
	for x.Cmp(zero) > 0 {
		x.QuoRem(x, m, r)
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}
	return s
}

func b58checkencode(ver uint8, b []byte) (s string) {
	bcpy := append([]byte{ver}, b...)
	sha256H := sha256.New()
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)
	bcpy = append(bcpy, hash2[0:4]...)
	s = b58encode(bcpy)
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}
	return s
}