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

	projectToken, projectOk := parsedPayload["projectToken"]
	modelName, modelOk := parsedPayload["dataModel"]

	if !projectOk {
		fmt.Println("Project id is required")
		return
	}

	if !modelOk {
		fmt.Println("Model name is required")
		return
	}

	userDefinedModel := []utils.DefiniedModel{
		utils.DefiniedModel{
			Name:     "temperature",
			Type:     "float64",
			Required: true,
		},
	}

	fmt.Println("datapoint=", len(userDefinedModel))
	resData, err := utils.ModelValidator(userDefinedModel, parsedPayload)
	tags := map[string]string{
		"projectToken": projectToken.(string),
		"modelName":    modelName.(string),
	}
	// err := utils.ModelValidator(userDefinedModel, parsedPayload)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All required fields present and have correct types")
		// fmt.Println(parsedPayload)
		utils.PushToInflux(resData, tags)
		// for key, value := range userDefinedModel {
		// 	fmt.Println(key, value)
		// }
	}

}
