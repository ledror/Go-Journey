package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var encodingFormat = flag.Int("e", 256, "Hash Width (256, 384, 512)")

func main() {
	flag.Parse()
	var encoding func([]byte) []byte
	switch *encodingFormat {
	case 256:
		encoding = func(b []byte) []byte {
			enc := sha256.Sum256(b)
			return enc[:]
		}
	case 384:
		encoding = func(b []byte) []byte {
			enc := sha512.Sum384(b)
			return enc[:]
		}
	case 512:
		encoding = func(b []byte) []byte {
			enc := sha512.Sum512(b)
			return enc[:]
		}
	default:
		log.Fatal("Bad encoding format (256, 384, 512)")
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", encoding(b))
}
