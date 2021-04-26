package basic

import (
	types "TorCoin/type"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)


//func b2s(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

func InitWallet (Language string) (types.Wallet, error) {

	MnemonicSlice, err := RandMnemonic(26, Language)

	if err != nil{
		return types.Wallet{}, err
	}

	var Mnemonic string
	for _,v := range MnemonicSlice{
		Mnemonic = Mnemonic + " " + v
	}

	Mnemonic = strings.TrimSpace(Mnemonic)

	var wallet types.Wallet
	
	wallet.Id = SHA256(RandomFactor(0,99))

	wallet.PasswordHash = SHA256(Mnemonic)

	Secret, err := EncryptPwd(wallet.PasswordHash)

	wallet.Secret = SHA512(Secret)

	wallet.Mnemonic = Mnemonic

	address, err := Address(WeHash(RandomFactor(1,30)) + wallet.Secret, wallet.Id)

	if err != nil{
		return types.Wallet{}, err
	}

	wallet.KeyPairs = append(wallet.KeyPairs, struct {
		Index     int    `json:"index"`
		Address string `json:"address"`
		SecretKey string `json:"secretKey"`
		PublicKey string `json:"publicKey"`
	}{Index: 0, Address: address.Address,SecretKey: address.PrivateKey, PublicKey: address.PublicKey})

	add,_ := json.Marshal(address)
	OutputFile([]byte(fmt.Sprintf("%s", add)), fmt.Sprintf("Database/address/%s.json",address.Address))

	wal,_ := json.Marshal(wallet)

	OutputFile([]byte(fmt.Sprintf("%s", wal)), fmt.Sprintf("Database/wallet/%s.json",wallet.Id))

	return wallet, nil

}

func GetWallet(WalletID string) (*types.Wallet, error) {

	file, err := ioutil.ReadFile("Database/wallet/" + WalletID + ".json")

	if err != nil{
		return nil, err
	}

	var JSON types.Wallet

	json.Unmarshal(file, &JSON)

	return &JSON, nil

}

func AddAddressToWallet(WalletID, Secret string) (*types.Wallet, error) {

	Wallet, err := GetWallet(WalletID)

	if err != nil{
		return nil, err
	}

	if Secret != Wallet.Secret {
		return nil, errors.New("secret mismatch")
	}

	address, err := Address(WeHash(RandomFactor(1,30)) + Wallet.Secret, Wallet.Id)

	if err != nil{
		return nil, err
	}

	Wallet.KeyPairs = append(Wallet.KeyPairs, struct {
		Index     int    `json:"index"`
		Address string `json:"address"`
		SecretKey string `json:"secretKey"`
		PublicKey string `json:"publicKey"`
	}{Index: len(Wallet.KeyPairs), Address: address.Address,SecretKey: address.PrivateKey, PublicKey: address.PublicKey})

	e,err := json.Marshal(address)
	OutputFile([]byte(fmt.Sprintf("%s", e)), fmt.Sprintf("Database/address/%s.json",address.Address))
	e,err = json.Marshal(Wallet)
	err = WriteToFile(fmt.Sprintf("%s",e), fmt.Sprintf("Database/wallet/%s.json",Wallet.Id))
	if err != nil {
		return nil, err
	}
	return Wallet, nil

}