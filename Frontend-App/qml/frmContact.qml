import QtQuick 2.0
import QtQuick.Controls 2.0

Page {
    width: 600
    height: 400
    visible: true
    id: root
    title: qsTr("Contacts")

    ListView {
        id: listView
        anchors.fill: parent
        model: PersonModel

        delegate: Item {
            x: 5
            width: 80
            height: 40
            Row {
                id: row1
                spacing: 10
                Rectangle {
                    width: 40
                    height: 40
                    color: colorCode
                }

                Text {
                        text: firstName + "  " + lastName
                    }
            }


        }


    }

    Button {
        id: button
        x: 250
        y: 180
        text: qsTr("Button")

        onClicked: root.StackView.view.push("qrc:/qml/frmChat.qml")
    }
}
