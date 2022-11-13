package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
)

func main() {
	secretKey := "yzbqklnj"
	fmt.Println(findAnswer(secretKey))
}

func findAnswer(secretKey string) int {
	currInt := 7406921
	for true {
		currIntStr := strconv.Itoa(currInt)
		currSecretKey := secretKey+currIntStr
		fmt.Println(currSecretKey)
		if checkZeros(grabMD5(currSecretKey)){
			return currInt
		}
		currInt++
	}
	return 0
}

func grabMD5(input string) [16]byte {
	inputBytes := []byte(input)
	md5bytes := md5.Sum(inputBytes)
	return md5bytes
}

func checkZeros(md5input [16]byte) bool {
	if md5input[0]==0 && md5input[1]==0 && md5input[2]==0 {
		return true
	} else {
		return false
	}
}




func convToHexString(md5input [16]byte) string {
	outputStr := ""
	for i:=0;i<len(md5input);i++{
		byteAsInt := int64(md5input[i])
		hexByte := strconv.FormatInt(byteAsInt, 16)
		if byteAsInt < 10 {
			hexByte = "0" + hexByte
		}
		outputStr = outputStr + hexByte
	}
	return outputStr
}