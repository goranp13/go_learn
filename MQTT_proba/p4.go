package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	tlsConfig := new(tls.Config)

	tlsConfig.ClientCAs = x509.NewCertPool()
	tlsConfig.RootCAs = x509.NewCertPool()

	if cert, err := tls.LoadX509KeyPair("/vernemq_ca.pem", "/vernemq_ca.pem"); err == nil {
		tlsConfig.Certificates = []tls.Certificate{cert}

		if pemData, err := ioutil.ReadFile("/vernemq_ca.pem"); err == nil {
			if !tlsConfig.ClientCAs.AppendCertsFromPEM(pemData) {
				log.Fatal("Failed")
			}
		}

		if pemData, err := ioutil.ReadFile("/vernemq_ca.pem"); err == nil {
			if !tlsConfig.RootCAs.AppendCertsFromPEM(pemData) {
				log.Fatal("Failed")
			}
		} else {
			log.Fatal("Failed")
		}
	} else {
		log.Fatal("Failed")
		return
	}

	opts := mqtt.NewClientOptions()
	opts.SetMaxReconnectInterval(time.Minute).SetAutoReconnect(true).SetClientID("9365daa4-acd6-46b3-b5f9-8abb37128b78").
		SetTLSConfig(tlsConfig).AddBroker("tls://vernemq-vernemq.apps.thingstalk.eu:8883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.WaitTimeout(time.Minute*10) && token.Error() == nil {
		log.Println("Connected to MQTT")
	} else {
		log.Println("Connection to MQTT failed", token.Error())
	}
}
