import QtQuick 2.12
import QtQuick.Controls 2.12
import QtGraphicalEffects 1.0
import QtQuick.Layouts 1.3
import QtQuick.Window 2.1
Page{
    width: 600
    height: 400

    title: qsTr("Infomation")
    Component {
        id: mainView

        Row {

            Column {
                spacing: 10
                Row {

                    Text {
                        text: qsTr("Your Name: ")
                        width: 120
                        anchors.topMargin: 25
                    }
                    TextField {
                        id: txtMyName
                        placeholderText: "Type something..."
                        Layout.fillWidth: true
                        selectByMouse: true
                    }

                }



                Row {
                    Text {
                        width: 120
                        text: qsTr("Friend's Name: ")
                        anchors.topMargin: 25
                    }
                    TextField {
                        id: txtFriendName
                        placeholderText: "Type something..."
                        Layout.fillWidth: true
                        selectByMouse: true
                    }

                }
                Button {
                    text: "Push"
                    onClicked: {
                        stack.push(["qrc:/application.qml", {chatWith: txtFriendName.text}] )
                    }
                }

            }

        }
    }


}
