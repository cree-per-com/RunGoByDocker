package main

import (
	"encoding/base64"
	"fmt"
)

func xorWithKey(data []byte, key []byte) []byte {
	encrypted := make([]byte, len(data))
	for i := range data {
		encrypted[i] = data[i] ^ key[i%len(key)]
	}
	return encrypted
}

func main() {
	// 암호화된 문자열
	encryptedData := "OwwBHANFOBATGgdNEQIBSyYREgwFBw1E"

	// 암호화에 사용한 키와 동일한 키 사용
	fullKey := []byte("simplekey")

	// Base64 인코딩된 문자열 디코딩
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// 복호화
	decrypted := xorWithKey(data, fullKey)
	fmt.Println("해독 결과:", string(decrypted))
}
