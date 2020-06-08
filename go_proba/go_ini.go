package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type gogi int

func main() {

	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println("MQTT_BROKER_HOST", cfg.Section("COMMON").Key("MQTT_BROKER_HOST").String())
	fmt.Println("MQTT_BROKER_PORT", cfg.Section("COMMON").Key("MQTT_BROKER_PORT").String())
	fmt.Println("MQTT_USE_TLS", cfg.Section("COMMON").Key("MQTT_USE_TLS").String())
	fmt.Println("MQTT_USER_NAME", cfg.Section("COMMON").Key("MQTT_USER_NAME").String())
	fmt.Println("MQTT_USER_PWD", cfg.Section("COMMON").Key("MQTT_USER_PWD").String())
	fmt.Println("MQTT_CLIENT_ID", cfg.Section("COMMON").Key("MQTT_CLIENT_ID").String())
	fmt.Println("MQTT_CERT_FILE", cfg.Section("COMMON").Key("MQTT_CERT_FILE").String())
	fmt.Println("MQTT_TOPIC", cfg.Section("COMMON").Key("MQTT_TOPIC").String())
	fmt.Println("MQTT_DEBUG_LEVEL", cfg.Section("COMMON").Key("MQTT_DEBUG_LEVEL").String())
	fmt.Println("\nDB_SERVER", cfg.Section("COMMON").Key("DB_SERVER").String())
	fmt.Println("DB_NAME", cfg.Section("COMMON").Key("DB_NAME").String())
	fmt.Println("DB_PWD", cfg.Section("COMMON").Key("DB_PWD").String())

	//kak mijenjati vrijednosti u ini file-u
	//	cfg.Section("COMMON").Key("DB_NAME").SetValue("AURADB")
	//    cfg.SaveTo("my.ini.local")
}
