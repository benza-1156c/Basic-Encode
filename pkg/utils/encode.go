package utils

func XorEncryptDecrypt(input string, key []byte) string {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ key[i%len(key)]
	}
	return string(result)
}
