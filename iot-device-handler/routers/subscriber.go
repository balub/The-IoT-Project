package routers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/balub/The-IoT-Project/controllers"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe() {
	// mqtt client setup
	opts := mqtt.NewClientOptions()
	opts.AddBroker("mqtt://localhost:1883")

	opts.OnConnect = controllers.HandleConnect
	opts.OnConnectionLost = controllers.HandleDisconnect

	client := mqtt.NewClient(opts)

	// check if connection successful or throw error
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe to a topic
	subscribeToTopic(client)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
}

// subscribe to a topic and listen to events
func subscribeToTopic(client mqtt.Client) {
	topic := "sonar"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)

	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s\n", msg.Payload())
	}

	client.AddRoute(topic, messageHandler)
}
