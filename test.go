package main

import (
	"dfa/smartcontract"
	"fmt"
)

func main() {
	contract := smartcontract.NewConract()
	contract.DeployContract()
	priv := "0xf1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"
	fmt.Println("")
	ok, tx, err := contract.Topup(priv, 20)
	fmt.Println(ok, tx, err)

	fmt.Println("")
	user, err := contract.Login(priv)
	fmt.Println(user)

	fmt.Println("")
	ok, tx, err = contract.Topup(priv, 20)
	fmt.Println(ok, tx, err)

	fmt.Println("")
	user, err = contract.Login(priv)
	fmt.Println(user)

	fmt.Println("")
	ok, tx, err = contract.WithDraw(priv, 5)
	fmt.Println(ok, tx, err)

	fmt.Println("")
	user, err = contract.Login(priv)
	fmt.Println(user)

}
