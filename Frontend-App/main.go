package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

type User struct {
	Id        string
	Username  string
	Password  string
	FirstName string
	LastName  string
}

type QmlBridge struct {
	core.QObject

	_ func(source, action, data string) `signal:"sendToQml"`
	_ func(source, action, data string) `slot:"sendToGo"`

	_ func(object *core.QObject) `slot:"registerToGo"`
	_ func(objectName string)    `slot:"deregisterToGo"`
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	qmlBridge.SendToQml("receiveMes", friendID, string(msg.Payload()))
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")

}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client, data, topic string) {

	token := client.Publish(topic, 0, false, data)
	token.Wait()
	time.Sleep(time.Second)

}

func sub(client mqtt.Client) {
	topic := myID
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
	qmlBridge.SendToQml("Sub", myID, topic+"Sub ok")
}

var qmlBridge = NewQmlBridge(nil)
var myID string
var friendID string

func main() {

	users := []User{
		{
			Id:        "1",
			Username:  "hoang1",
			Password:  "1",
			FirstName: "Hoang",
			LastName:  "Do",
		},
		{
			Id:        "2",
			Username:  "hoang2",
			Password:  "2",
			FirstName: "Hoang",
			LastName:  "Do Quoc",
		},
		{
			Id:        "3",
			Username:  "yen",
			Password:  "3",
			FirstName: "Yen",
			LastName:  "Nguyen Thi",
		},
	}

	/* -------- connect MQTT Broker --------*/

	if myID == "none" {
		fmt.Println(myID, friendID)
	}
	var broker = "localhost"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(myID)
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	/*--------------------------------------*/
	/* -------- Init QML GUI --------*/

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var app = qml.NewQQmlApplicationEngine(nil)
	app.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	app.RootContext().SetContextProperty("qmlBridge", qmlBridge)

	var personModel = NewPersonModel(nil)

	for _, v := range users {
		var p = NewPerson(nil)
		p.SetId(v.Id)
		p.SetFirstName(v.FirstName)
		p.SetLastName(v.LastName)
		personModel.AddPerson(p)
		//personModel.SetPeople([]*Person{p})
	}

	app.RootContext().SetContextProperty("PersonModel", personModel)
	//qmlBridge.SendToQml("FriendID", myID, friendID)

	//mesCount := 0
	qmlBridge.ConnectSendToGo(func(source, action, data string) {
		// if source == "username" && action == "click" {
		// 	//username = data
		// }
		if source == "login" {
			check := false

			for _, v := range users {
				if action == v.Username {
					check = true
					myID = v.Id
				}
			}

			if check {
				qmlBridge.SendToQml("loginResponse", "", "true")
				// for _, v := range users {
				// 	if myID != v.Id {
				// 		p.SetId(v.Id)
				// 		p.SetFirstName(v.FirstName)
				// 		p.SetLastName(v.LastName)
				// 		personModel.AddPerson(p)
				// 	}
				// }
			} else {
				qmlBridge.SendToQml("loginResponse", "", "false")
			}
		}
		if source == "usrID" {
			myID = action
			friendID = data
			qmlBridge.SendToQml("FriendID", myID, friendID)
			go sub(client)
		}
		if source == "sendBtn" {
			//mesCount++
			qmlBridge.SendToQml("mesSent", myID, data)
			publish(client, data, action)

		}
	})

	gui.QGuiApplication_Exec()

}
