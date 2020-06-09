package main

import (
	"log"
	"os"
	"time"
)

var t = time.Now()
var d = t.Format("2006-01-02")

func main() {

	/*file, e := os.OpenFile("log_go.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if e != nil {
		log.Fatalln("failed")
	}

	defer file.Close()

	log.SetOutput(file)

	log.Println("This is the log file.")*/
	fnmqtt_log()

	fnfbd_log()
}

func fnmqtt_log() {

	file, e := os.OpenFile(d+"_mqtt_log_go.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if e != nil {
		log.Fatalln("failed")
	}

	defer file.Close()

	log.SetOutput(file)

	log.Println("This is the mqtt_log file.")
}

func fnfbd_log() {

	file, e := os.OpenFile(d+"_fbd_log_go.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if e != nil {
		log.Fatalln("failed")
	}

	defer file.Close()

	log.SetOutput(file)

	log.Println("This is the fbd_log file.")
}
