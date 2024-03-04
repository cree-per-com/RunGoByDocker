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
	// 이 문자열이 무엇일까요?
	encryptedData := "OwwBHANFOBATGgdNEQIBSyYREgwFBw1E"

	fullKey := []byte("simplekey")

	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// 해독 결과 출력
	decrypted := xorWithKey(data, fullKey)
	fmt.Println("해독 결과:", string(decrypted))
}

//docker build -t my-go-app .
//docker run -d my-go-app
