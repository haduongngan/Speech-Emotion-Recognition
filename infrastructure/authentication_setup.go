package infrastructure

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"

	"github.com/go-chi/jwtauth"
)

const Algorithm = "RS256"

func loadAuthToken() error {
	// Load privateKey
	privateReader, err := ioutil.ReadFile(rsaPrivatePath)
	if err != nil {
		log.Println("No RSA private pem file: ", err)
		return err
	}

	privatePem, _ := pem.Decode(privateReader)
	privateKey, err = x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		log.Println("Problem creating authentication:", err)
		return err
	}

	// Load publicKey
	publicReader, err := ioutil.ReadFile(rsaPublicPath)
	if err != nil {
		log.Println("No RSA public pem file: ", err)
		return err
	}

	publicPem, _ := pem.Decode(publicReader)
	publicKey, err = x509.ParsePKIXPublicKey(publicPem.Bytes)

	encodeAuth = jwtauth.New(Algorithm, privateKey, publicKey)
	decodeAuth = jwtauth.New(Algorithm, nil, publicKey)

	return nil
}
