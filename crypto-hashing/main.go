package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	/*
		fmt.Println(hashes.Scrypt("Salom chi heli", "ParoliBajoi"))

		text := []byte("«Алиф Сармоя» была основана в 2014 году как микрокредитная организация.")
		key := []byte("the-key-has-to-be-32-bytes-long!")

		ciphertext, err := cryptos.Encrypt(text, key)
		if err != nil {
			// TODO: Properly handle error
			log.Fatal(err)
		}
		fmt.Printf("%s => %x\n", text, ciphertext)

		plaintext, err := cryptos.Decrypt(ciphertext, key)
		if err != nil {
			// TODO: Properly handle error
			log.Fatal(err)
		}
		fmt.Printf("%x => %s\n", ciphertext, plaintext)

	*/

	caCert, err := ioutil.ReadFile("client.crt")
	if err != nil {
		log.Println(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	srv := &http.Server{
		Addr:         "localhost:10443",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		TLSConfig:    cfg,
		IdleTimeout:  30 * time.Minute,
	}

	http.HandleFunc("/", index)
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

//get and delete card
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("response from :10443"))
}
