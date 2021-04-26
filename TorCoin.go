package main

import (
	"TorCoin/src/basic"
	"encoding/json"
	"fmt"
)

func main()  {

	/*Init.HelloToRc()

	route := router.SetupRouter()

	if err := route.Run(conf.Get("Server","BindAddress")); err != nil {

		fmt.Println("CANT RUN SERVER")

	}*/

	for i := 0; i < 10; i++ {
		src, err := basic.AddAddressToWallet("f720e4808fca10cde6d531b2444fe59d2dcd6de88c5fe6f2900adb61790421f9","b632f767cbd492747382ee5f197b4d7b6f2b8707485617340f404908ba49851350de48de57c5c661bb12163e2916998c1ad33a9cbf62bdb934705ebbf13a313a")
		if err != nil{
			fmt.Println(err)
			return
		}
		e, _ := json.Marshal(src)
		fmt.Println(fmt.Sprintf("%s",e))
	}
	//var result []map[string]interface{}

}
