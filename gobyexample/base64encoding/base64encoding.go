package main

// Rename base64 to base64
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64 encoding
	// Encoding with standard encoder, convert string to []byte
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding may return an error
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Encoding/decoding URL-compatible base64 format
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
