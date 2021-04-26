package types

var(
	BasicRandomChars = []byte("abcdefghjkmnpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

const (
	Proportion int = 1000000
)

type TransactionHash struct {
	Previous       string  `json:"Previous"`
	BlockID        string  `json:"BlockID"`
	Initiator      string  `json:"Initiator"`
	Receiver       string  `json:"Receiver"`
	Value          int `json:"Value"`
	TransactionFee int `json:"TransactionFee"`
	Timestamp      int64     `json:"Timestamp"`
}

type Blockchain struct {
	Index        int    `json:"index"`
	PreviousHash string `json:"previousHash"`
	Timestamp    int    `json:"timestamp"`
	Nonce        int    `json:"nonce"`
	Transactions []struct {
		Id   string `json:"id"`
		Hash string `json:"hash"`
		Type string `json:"type"`
		Data struct {
			Inputs  []interface{} `json:"inputs"`
			Outputs []interface{} `json:"outputs"`
		} `json:"data"`
	} `json:"transactions"`
	Hash string `json:"hash"`
}

type Address struct {
	WalletID   string `json:"WalletID"`
	Index 	   string `json:"Index"`
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
	Address    string `json:"Address"`
	Amount     int64  `json:"amount"`
}

type Wallet struct {
	Id           string `json:"id"`
	PasswordHash string `json:"passwordHash"`
	Secret       string `json:"secret"`
	Mnemonic       string `json:"mnemonic"`
	KeyPairs     []struct {
		Index     int    `json:"index"`
		Address string `json:"address"`
		SecretKey string `json:"secretKey"`
		PublicKey string `json:"publicKey"`
	} `json:"keyPairs"`
}


type NewAddressRequest struct {
	Language string `json:"language"`
}
