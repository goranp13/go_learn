/*
 * Yokai, a simple MQTT "I'm alive" publisher for Dodomeki"
 * (c) jme@opium.io for crashdump.net
 * BSD LICENCE
 */

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const mqtt_protocol string = "tls://"                            // mqtt protocol to use (tls:// is prefered)
const mqtt_host string = "vernemq-vernemq.apps.thingstalk.eu"    // mqtt server fqdn or ip
const mqtt_port string = "8883"                                  // mqtt server port
const mqtt_login string = "9365daa4-acd6-46b3-b5f9-8abb37128b78" // mqtt server login
const mqtt_passwd string = ""                                    // mqtt server password
const mqtt_topic string = "telemetry/v1/APPLICATION/156"         //topic to publish
const mqtt_id string = "9365daa4-acd6-46b3-b5f9-8abb37128b78"    //mqtt iD and topic trailing id.
const hmac_secret string = ""                                    //hmac secret to sign msg

const broker = mqtt_protocol + mqtt_host + ":" + mqtt_port
const mqtt_subscribe = mqtt_topic + "/" + mqtt_id

var message_string string = ""

//yes that is a self signed certificate but nothing wrong with them
//yes, I am aware that the fqdn is in the Subject…
const mqtt_tls_ca = `-----BEGIN CERTIFICATE-----
MIIC7DCCAdQCCQCgpxNx6WGWADANBgkqhkiG9w0BAQsFADA4MQswCQYDVQQGEwJI
UjEXMBUGA1UECgwOT21lZ2EgU29mdHdhcmUxEDAOBgNVBAMMB1Jvb3QgQ0EwHhcN
MTkxMDIxMDY1MzAzWhcNNDQxMDE0MDY1MzAzWjA4MQswCQYDVQQGEwJIUjEXMBUG
A1UECgwOT21lZ2EgU29mdHdhcmUxEDAOBgNVBAMMB1Jvb3QgQ0EwggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDJSE5YuolF1XoCzFbl3w+0ai2MKE2CNJxv
DYXGgMTl85tucu+vyjvC4qC3e7g0VArya5i5uqVEZldUDaWtNICJw5YJ2lrRg6cJ
Una3O/XUrFkOXbmcbsipGzvUc/f0rJzqs0Halc962EeTk8RYl/FIoYKTTJkwHVnr
OPhtmdvud+r7sbjVWvdB8uDfIR7lxK99HL1kGKoPsdnFPpODljY2ti9xBWWb4ERQ
rR3AQpMjg4PpN4AVw4cAdXrjJf68Wbyc0ysm46ziL7IjP126UMC3ZBxEzhto5xYT
KyB3KEX0E4pGM8c/A6XUlQze5K9f8LWwCbORAuIuMkxCk6CFI1CPAgMBAAEwDQYJ
KoZIhvcNAQELBQADggEBAMMnQXRPdp3xz0MJ0U+58MN1Wq4P684SKaRVGVCOClvY
GcZfA8aP+LLO/kvSaa0dTK9j2HlqVMBZDMLMjBDxA31Zcu1sS+cfFVUixKBRTk5F
Ba/6MhklsR/c7IYcNMpyTXBq/AG6YI3eyowC3f19DzrkLrBbUIy+Hk+Jx810eH9u
RK45ANvl6Y8iSblfNX3ySxIJk205Aq6yAioZtFMMchkhc/ZxEhOJssL65lyWdyBn
86lJYFEcxfnnfJaNES4qvUabHVokzvA4bSEt1bbPpi7UIa/iCKREKWBZoJGAJG1k
IbzB4qTT2l06vZ+Q9fASsKnzhn7BUoK5vosHF3xiiNY=
-----END CERTIFICATE-----`

/* func build_message() string {
	// return the I'm alive message in the form of [current_epoch_time]:[hmac_verification]
	// the [hmac_verification] is a sha256 hmac of [mqtt_id]:[the alive message itself]
	// the hmac is to prevent a rogue device to send alive messages in place of another one
	// the current time in epoch format provide the nonce against replay attacks
	now := time.Now()
	var my_msg string = fmt.Sprintf("%d", now.Unix())
	var my_verification string = mqtt_id + ":" + my_msg
	message_hmac := hmac.New(sha256.New, []byte(hmac_secret))
	message_hmac.Write([]byte(my_verification))
	myHmac := fmt.Sprintf("%s", base64.StdEncoding.EncodeToString(message_hmac.Sum(nil)))

	//var my_msg string = "12"
	//fmt.Println("ok")
	my_msg += ":" + myHmac
	return my_msg
} //message */

func main() {
	root_ca := x509.NewCertPool()
	load_ca := root_ca.AppendCertsFromPEM([]byte(mqtt_tls_ca))
	if !load_ca {
		panic("failed to parse root certificate")
	}

	tlsConfig := &tls.Config{RootCAs: root_ca}

	opts := MQTT.NewClientOptions()
	opts.SetTLSConfig(tlsConfig)  //we set the tls configuration
	opts.AddBroker(broker)        //we add the broker
	opts.SetClientID(mqtt_id)     //we set the mqtt id
	opts.SetCleanSession(true)    // Sets true to client and server should remember state across restarts and reconnect
	opts.SetUsername(mqtt_login)  // Set the mqtt server login
	opts.SetPassword(mqtt_passwd) // Set the mqtt server password
	c := MQTT.NewClient(opts)     // Launch the client using the set options
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	var message string = "34"
	text := fmt.Sprintf("%s", message)
	token := c.Publish(mqtt_subscribe, 0, false, text)
	token.Wait()
	c.Disconnect(250)
} //main
