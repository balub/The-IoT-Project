package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/balub/The-IoT-Project/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var HandleMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var parsedPayload map[string]interface{}

	if error := json.Unmarshal(msg.Payload(), &parsedPayload); error != nil {
		fmt.Println("unable to parse data")
		return
	}

	_, projectOk := parsedPayload["projectToken"]
	_, modelOk := parsedPayload["dataModel"]

	if !projectOk {
		fmt.Println("Project id is required")
		return
	}

	if !modelOk {
		fmt.Println("Model name is required")
		return
	}

	userDefinedModel := map[string]map[string]interface{}{
		"val": {
			"type":     "float64",
			"required": true,
		},
		"val2": {
			"type":     "string",
			"required": true,
		},
	}

	err := utils.ModelValidator(userDefinedModel, parsedPayload)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All required fields present and have correct types")
		// fmt.Println(parsedPayload)
		utils.PushToInflux(parsedPayload)
	}

}
