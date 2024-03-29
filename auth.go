package wisego

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func loadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	pemBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &rsa.PrivateKey{}, err
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return &rsa.PrivateKey{}, errors.New("no key found")
	}

	switch block.Type {
	case "PRIVATE KEY":
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return &rsa.PrivateKey{}, err
		}
		rsaKey, ok := privateKey.(*rsa.PrivateKey)
		if !ok {
			log.Fatalf("got unexpected key type: %T", privateKey)
		}
		return rsaKey, nil
	default:
		return &rsa.PrivateKey{}, fmt.Errorf("unsupported key type %q", block.Type)
	}
}

func (w *Wise) generateNewSignedToken(privateKeyPath string, resp *http.Response) (string, string) {
	oneTimeToken := resp.Header.Get("x-2fa-approval")
	if w.Debug {
		log.Printf("OTT was %v from the previous response (%v)", oneTimeToken, resp.Request.URL)
	}
	if oneTimeToken == "" {
		log.Println("Could not get twoFactor Header, stopping!")
		return "", ""
	}
	privateKey, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		fmt.Errorf("signer is damaged: %v", err)
		return "Error", oneTimeToken
	} else {
		if w.Debug {
			log.Printf("Private Key successfully laoded from: %v", privateKeyPath)
		}
	}
	hashed := sha256.Sum256([]byte(oneTimeToken))
	signedToken, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Errorf("could not sign token: %v", err)
		return "", ""
	} else {
		return base64.StdEncoding.EncodeToString(signedToken), oneTimeToken
	}
}

func (w *Wise) retryRequestWithToken(method, endPoint, signature, oneTimeToken string, data io.Reader, withUUID bool) []byte {
	var apiURL = "https://api.transferwise.com"
	if w.SandBox {
		apiURL = "https://api.sandbox.transferwise.tech"
	}
	req, err := http.NewRequest(method, apiURL+endPoint, data)
	if err != nil {
		log.Println(err.Error())
	}
	if w.Debug {
		log.Printf("Signing new request with %v and a OTT of %v ", signature, oneTimeToken)
	}
	req.Header.Set("Authorization", "Bearer "+w.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-2fa-approval", oneTimeToken)
	req.Header.Set("X-Signature", signature)
	if withUUID {
		newUUID := uuid.New().String()
		req.Header.Set("X-idempotence-uuid", newUUID)
	}
	brandNewClient := &http.Client{}
	resp, err := brandNewClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	return body
}
