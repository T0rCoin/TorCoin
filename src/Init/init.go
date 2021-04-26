package Init

import "fmt"

const (

	Version              = 0x00
	addreddChechsumLen   = 4
	PrivKeyBytesLen      = 32

)

func HelloToRc()  {
	fmt.Println("  ______              ______      _     ")
	fmt.Println(" /_  __/___  _____   / ____/___  (_)___ ")
	fmt.Println("  / / / __ \\/ ___/  / /   / __ \\/ / __ \\")
	fmt.Println(" / / / /_/ / /     / /___/ /_/ / / / / /")
	fmt.Println("/_/  \\____/_/      \\____/\\____/_/_/ /_/ ")
}
