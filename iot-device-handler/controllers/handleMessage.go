package controllers

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var HandleMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var parsedPayload map[string]interface{}

	if error := json.Unmarshal(msg.Payload(), &parsedPayload); error != nil {
		fmt.Println("unable to parse data")
		return
	}

	projectToken, projectOk := parsedPayload["projectToken"]
	dataModelName, modelOk := parsedPayload["dataModel"]

	if !projectOk {
		fmt.Println("Project id is required")
		return
	}

	if !modelOk {
		fmt.Println("Model name is required")
		return
	}

	fmt.Println(projectToken, dataModelName)
}
