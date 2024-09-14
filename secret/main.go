package main

import (
	"aidanwoods.dev/go-paseto"
	"fmt"
)

func main() {
	secret := paseto.NewV4AsymmetricSecretKey()
	public := secret.Public()
	fmt.Println(public.ExportHex())
	fmt.Println(secret.ExportHex())
}
