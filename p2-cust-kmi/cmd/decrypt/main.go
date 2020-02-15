package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get new cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to get nre GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}

	return plaintext, nil
}

var passphrase = flag.String("passphrase", "", "passphrase to use for encrypt")
var file = flag.String("file", "", "file to encrypt")
var outFile = flag.String("outfile", "decrypt.out", "name of the output file")

const usage = "Usage: decrypt --passphrase <passphrase> --file <filepath> [--outfile <filename>]"

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

	data, err := decrypt(bs, *passphrase)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decrypt contents: %s", err)
	}

	fn := func() string {
		if *outFile == "" {
			return fmt.Sprintf("%s.decrypted", *file)
		}
		return *outFile
	}()

	if err := ioutil.WriteFile(fn, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write file to disk: %s", err)
	}
}
