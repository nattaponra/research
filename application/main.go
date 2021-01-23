package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/shirou/gopsutil/mem"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {

	time.Sleep(3 * time.Second)
	fmt.Println(time.Now())

	var broker = "167.172.5.53"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("543b763b-c819-496e-823b-fce6ef664880")
	opts.SetUsername("543b763b-c819-496e-823b-fce6ef664880")
	opts.SetPassword("ac58863d-cd20-46cf-9e63-1f6061cd2ee9")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer client.Disconnect(250)

	for {
		v, _ := mem.VirtualMemory()

		// // almost every return value is a struct
		// fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n")

		publish(client, v.Total, v.Free, v.UsedPercent)
	}

}
func publish(client mqtt.Client, total, free uint64, usedPercent float64) {
	text := fmt.Sprintf(`[{"n":"total","v":%d},{"n":"free","v":%d},{"n":"usedPercent","v":%f}]`, total, free, usedPercent)
	fmt.Println(text)
	token := client.Publish("organization/d8d74ab3-2689-4343-bf79-78f2b0083fd9/messages", 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}
