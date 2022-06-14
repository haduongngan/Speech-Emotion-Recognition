package infrastructure

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/go-chi/jwtauth"
)

// Alogirthm algorithm deinfe
const Alogirthm = "RS256"

func loadAuthToken() error {
	// Load private key
	privateReader, err := ioutil.ReadFile(privatePath)
	if err != nil {
		InfoLog.Println("NO RSA private pem file")
		return err
	}
	privatePem, _ := pem.Decode(privateReader)

	if privatePem.Type != "RSA PRIVATE KEY" {
		InfoLog.Println("RSA private key is of the wrong type")
	}

	// privatePemBytes, err := x509.DecryptPEMBlock(privatePem, []byte(privatePassword))
	// if err != nil {
	// 	log.Println(err)
	// }

	privateKey, err := x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		log.Println(err)
	}
	// Read public key
	publicReader, err := ioutil.ReadFile(publicPath)
	if err != nil {
		InfoLog.Println("No RSA public pem file")
		return err
	}
	publicPem, _ := pem.Decode(publicReader)
	publicKey, _ := x509.ParsePKIXPublicKey(publicPem.Bytes)

	encodeAuth = jwtauth.New(Alogirthm, privateKey, publicKey)
	decodeAuth = jwtauth.New(Alogirthm, nil, publicKey)

	return nil
}
