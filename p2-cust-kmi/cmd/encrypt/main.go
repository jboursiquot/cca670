package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(createHash(passphrase)))
	if err != nil {
		return nil, fmt.Errorf("failed to get new cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to get new GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cyphertext := gcm.Seal(nonce, nonce, data, nil)
	return cyphertext, nil
}

var passphrase = flag.String("passphrase", "", "passphrase to use for encrypt")
var file = flag.String("file", "", "file to encrypt")

const usage = "Usage: encrypt --passphrase <passphrase> --file <filepath>"

func main() {
	flag.Parse()

	if *passphrase == "" || *file == "" {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	bs, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %s", err)
	}

	data, err := encrypt(bs, *passphrase)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encrypt contents: %s", err)
	}

	outFile := fmt.Sprintf("%s.encrypted", *file)
	if err := ioutil.WriteFile(outFile, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file to disk: %s", err)
	}
}
