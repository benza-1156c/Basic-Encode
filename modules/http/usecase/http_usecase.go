package usecase

import (
	"encode/pkg/utils"
	"os"
)

func Encode(input string) (string, string) {
	key := os.Getenv("KEY")
	encode := utils.XorEncryptDecrypt(input, []byte(key))

	return encode, key
}
