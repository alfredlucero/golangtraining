package main

import (
	"crypto/sha1"
	"fmt"
)

// SHA1 hashes are frequently used to compute short identities for binary or text blobs
// i.e. git uses SHA1s to identify versioned files and directories

func main() {
	s := "sha1 this string"

	// sha1.New(), sha1.Write(bytes), and sha1.Sum([]byte{})
	h := sha1.New()

	// Expects bytes so we must coerce string
	h.Write([]byte(s))

	// Get finalized hash result as a byte slice
	// Argument can be used to append to an existing byte slice and not usually needed
	bs := h.Sum(nil)

	// SHA1 values often printed in hex like in git commits
	// Use %x format verb to convert a hash result to a hex string
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	// To compute other hashes with similar pattern, can compute MD5 hash with crypto/md5 and md5.New9)
}
