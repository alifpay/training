package hashes

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/scrypt"
)

//Argon is the winner of the password hashing competition and should be considered as your first choice for new applications
func Argon(plainText, salt string) string {
	hashByte := argon2.IDKey([]byte(plainText), []byte(salt), 1, 64*1024, 4, 32)
	return hex.EncodeToString(hashByte)
}

//HmacSha256 Generate returns hash based on provided type.
func HmacSha256(key, val string) string {
	hashhmac := hmac.New(sha256.New, []byte(key))
	hashhmac.Write([]byte(val))
	return hex.EncodeToString(hashhmac.Sum(nil))
}

//Scrypt crypto hash algorithm
func Scrypt(plainText, salt string) string {
	//salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

	dk, err := scrypt.Key([]byte(plainText), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}
