package main

import (
	"fmt"

	"goWeb/web9/cipher"
	"goWeb/web9/lzw"
)

var sentData string
var recvData string

//Component is ...
type Component interface {
	Operator(string)
}

//SendComponent is ...
type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	sentData = data
}

//SendComponent is ...
type ReceiveComponent struct{}

func (self *ReceiveComponent) Operator(data string) {
	recvData = data
}

//ZipComponent is ...
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))

	if err != nil {
		panic(err)
	}

	self.com.Operator(string(zipData))
}

//UnzipComponent is ...
type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))

	if err != nil {
		panic(err)
	}

	self.com.Operator(string(unzipData))
}

//EncryptComponent is ...
type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(encryptData))
}

//DecryptComponent is ...
type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(decryptData))
}

func main() {
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello World")

	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReceiveComponent{},
		},
	}

	receiver.Operator(sentData)
	fmt.Println(recvData)

}
