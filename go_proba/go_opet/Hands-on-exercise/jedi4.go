package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type gogi int

var x gogi
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println("MQTT_BROKER_HOST", cfg.Section("").Key("MQTT_BROKER_HOST").String())

}
