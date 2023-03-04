package controllers

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var handleMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
