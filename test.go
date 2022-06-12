package main

import (
	"dfa/entity"
	"dfa/general"
	"dfa/smartcontract"
	"fmt"
)

func main() {
	contract := smartcontract.NewConract()
	contract.DeployContract()
	priv1 := "0xf1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5"
	priv2 := "0x3b24a4fdf2e6e1375008c387c5456ce00cb0772435ae1938c2fe833103393b9a"
	// fmt.Println("")
	// ok, tx, err := contract.Topup(priv, 20)
	// fmt.Println(ok, tx, err)

	// fmt.Println("")
	// user, err := contract.Login(priv)
	// fmt.Println(user)

	// fmt.Println("")
	// ok, tx, err = contract.Topup(priv, 20)
	// fmt.Println(ok, tx, err)

	// fmt.Println("")
	// user, err = contract.Login(priv)
	// fmt.Println(user)

	// fmt.Println("")
	// ok, tx, err = contract.WithDraw(priv, 5)
	// fmt.Println(ok, tx, err)

	// fmt.Println("")
	// user, err = contract.Login(priv)
	// fmt.Println(user)

	data := entity.Data{
		ID:              "file_0",
		PermissionLevel: uint8(entity.L_2),
		Tradable:        true,
		Price:           22,
		MeteData: entity.MeteData{
			FileName:   "name",
			HashDigest: "digest",
			Signer:     "me",
			Size:       22,
			TimeStamp:  general.GenerateTimeStamp(),
		},
	}
	fmt.Println("priv1 Create a file")
	tx, err := contract.CreateFile(priv1, data)
	fmt.Println(tx, err)

	fmt.Println("priv1 write a file")
	contract.WriteFile(priv1, "file_0", "hash", data.MeteData)
	fmt.Println(tx, err)

	fmt.Println("priv2 download a file")
	tx, d, err := contract.ReadFile(priv2, "file_0")
	fmt.Println(d)

	fmt.Println("priv2 topup 50")
	ok, tx, err := contract.Topup(priv2, 50)
	fmt.Println(ok, tx, err)

	fmt.Println("priv2 approve 30")
	ok, tx, err = contract.Approve(priv2, 30)
	fmt.Println(ok, tx, err)

	fmt.Println("priv2 approved")
	amount, err := contract.GetApproved(priv2)
	fmt.Println(amount, err)

	fmt.Println("priv2 purchase a file")
	tx, ok, err = contract.PurchaseFile(priv2, "file_0")
	fmt.Println(ok, tx, err)

	fmt.Println("priv2 query status")
	user, err := contract.Login(priv2)
	fmt.Println(user)

	fmt.Println("priv1 query allowance")
	amount, err = contract.GetAllowance(priv1)
	fmt.Println(amount)

	fmt.Println("priv1 transfer from contract")
	tx, err = contract.TransferFrom(priv1, 22)
	fmt.Println(tx, err)

	fmt.Println("priv1 login")
	user, err = contract.Login(priv1)
	fmt.Println(user, err)

	fmt.Println("priv1 query allowance")
	amount, err = contract.GetAllowance(priv1)
	fmt.Println(amount)

}
