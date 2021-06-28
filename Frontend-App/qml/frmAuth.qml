import QtQuick 2.6
import QtQuick.Controls 2.0

Page {
    width: 600
    height: 400
    visible: true
    id: root
    title: qsTr("Authentication")

    Connections {
        target: qmlBridge

        onSendToQml: {
            if(source == "loginResponse" && data == "true")
            {
                root.StackView.view.push("qrc:/qml/frmContact.qml", {chatWith: friendID.text, myID: myID.text, friendID: friendID.text, whoSend: myID})
            }
            else {
                label2.visible = true
            }
        }
    }

    Button {
        id: button
        x: 362
        y: 245
        text: qsTr("Login")

        onClicked: {
            qmlBridge.sendToGo("login", myID.text, friendID.text)

            //qmlBridge.sendToGo("usrID", myID.text, friendID.text)

        }
    }

    TextField {
        id: friendID
        x: 262
        y: 191
        selectByMouse: true
    }

    TextField {
        id: myID
        x: 262
        y: 134
        selectByMouse: true
    }

    Label {
        id: label
        x: 151
        y: 147
        text: qsTr("Username:")
    }

    Label {
        id: label1
        x: 151
        y: 204
        text: qsTr("Password:")
    }

    Label {
        id: label2
        x: 251
        y: 99
        visible: false
        color: "#ff0404"
        text: qsTr("Username or Password is invalid!")
    }

}

/*##^##
Designer {
    D{i:0;formeditorZoom:0.75}
}
##^##*/
