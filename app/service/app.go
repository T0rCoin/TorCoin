package service

import (
	"TorCoin/src/basic"
	types "TorCoin/type"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Root(c * gin.Context)  {
	c.JSON(201, gin.H{
		"code": 201,
		"message": "TORc Node",
	})
}

func NewAddress(c * gin.Context)  {

	var Password = c.Request.Header.Get("password")
	var WalletID = c.Param("id")

	//Header->password PasswordHash

	if Password == "" || WalletID == "" {
		c.JSON(503, gin.H{
			"msg": "Please check",
		})
		return
	}

	address, err := basic.Address(Password, WalletID)

	if err != nil{
		c.JSON(503, gin.H{
			"msg": err,
		})
		return
	}

	NiggerAddress := types.Address{
		Index:      address.Index,
		PublicKey:  address.PublicKey,
		PrivateKey: address.PrivateKey,
		Address:    address.Address,
		Amount:     0,
	}

	c.JSON(200, gin.H{
		"address": NiggerAddress.Address,
	})

	e,_ := json.Marshal(address)
	basic.OutputFile([]byte(fmt.Sprintf("%s", e)), fmt.Sprintf("Database/address/%s.json",address.Address))
	return
}

func GetBalance(c * gin.Context)  {
	address := c.Param("address")
	if address == ""{
		c.JSON(401, gin.H{
			"msg": "Unable to get address",
		})
		return
	}
	status, err := basic.GetAddressStatus(address)
	if err != nil{
		c.JSON(503, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"address": status.Address,
		"amount": status.Amount,
	})
	return
}