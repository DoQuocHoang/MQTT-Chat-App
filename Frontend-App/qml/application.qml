import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.3
import QtQuick.Window 2.1

ApplicationWindow {
    id: chatView
    visible: true
    title: chatWith
    property int margin: 11
    minimumWidth: 600
    minimumHeight: 450

    property string chatWith
    property string chatWith

    Connections {
        target: qmlBridge
                        onSendToQml:
                        {
                            if (source == "FriendID" && action == "txt" && data !== "")
                            {
                                chatWith = data
                                whoSend = data
                            }
                            if (source == "receiveMes" && action == "txt" && data !== "")
                            {
                                whoSend = "friend"
//                                var row = {"mesData": data }
                                listView.model.insert(0, {"mesData": data })
                            }
                            if (source == "mesContent" && action == "txt" && data !== "")
                            {
                                //mesIDCount++
                                //mesField.text = data
                                //mesField.width = mesField.text.contentWidth

                                listView.model.insert(0, {"mesData": data })
                                // mesListView.model.append({"mesData": data})

                                typingMes.text = ""
                            }
                                    
                        }
                    }

    ColumnLayout {
        id: mainLayout
        anchors.fill: parent
        anchors.margins: margin
        GroupBox {
            id: messagesBox
            Layout.fillHeight: true
            Layout.fillWidth: true
            anchors.top: parent.top
            anchors.bottom: typingBox.top
            anchors.bottomMargin: margin

            ScrollView {
                width: parent.width+12
                height: parent.height 
                clip: true

                ListView {
                    id: listView
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    Layout.margins: typingBox.leftPadding + typingMes.leftPadding
                    anchors.right: parent.right
                    anchors.rightMargin: 17
                    verticalLayoutDirection: ListView.BottomToTop
                    spacing: 12
                    model: ListModel {}
                    delegate: Column {
                        anchors.right: sentByMe ? parent.right : undefined
                        spacing: 6

                        readonly property bool sentByMe: whoSend !== "me"

                        Row {
                            id: messageRow
                            spacing: 6
                            anchors.right: sentByMe ? parent.right : undefined

                            Image {
                                id: avatar
                                //source: !sentByMe ? "qrc:/" + model.author.replace(" ", "_") + ".png" : ""
                            }

                            Rectangle {
                                width: Math.min(messageText.implicitWidth + 24,
                                    listView.width - (!sentByMe ? avatar.width + messageRow.spacing : 0))
                                height: messageText.implicitHeight + 24
                                color: sentByMe ? "lightgrey" : "steelblue"

                                Label {
                                    id: messageText
                                    text: mesData
                                    color: sentByMe ? "black" : "white"
                                    anchors.fill: parent
                                    anchors.margins: 12
                                    wrapMode: Label.Wrap
                                }
                                
                            }
                        }

                        Label {
                            id: timestampText
                            text: Qt.formatDateTime(model.timestamp, "d MMM hh:mm")
                            color: "lightgrey"
                            anchors.right: sentByMe ? parent.right : undefined
                        }
                    }

                    //ScrollBar.vertical: ScrollBar {}
        

                    

                    // Connections {
                    //     target: qmlBridge
                    //     onSendToQml:
                    //     {
                    //         if (source == "mesContent" && action == "txt" && data !== "")
                    //         {
                    //             //mesIDCount++
                    //             //mesField.text = data
                    //             //mesField.width = mesField.text.contentWidth
                    //             var row = {"mesData": data }
                    //             listView.model.insert(0, row)
                    //             // mesListView.model.append({"mesData": data})

                    //             typingMes.text = ""
                    //         }
                                    
                    //     }
                    // }
                }

            

        }
        }
        GroupBox {
            id: typingBox
            Layout.fillWidth: true
            anchors.bottom: parent.bottom

            RowLayout {
                id: typingLayout
                anchors.fill: parent
                TextField {
                    id: typingMes
                    placeholderText: "Type something..."
                    Layout.fillWidth: true
                    selectByMouse: true
                }
                Button {
                    id: sendBtn
                    text: "Send"
                    onClicked: {
                        qmlBridge.sendToGo("sendBtn", "click", typingMes.text)
                        whoSend = "me"
                    }
                }
            }
        }

    }
}

/*##^##
Designer {
    D{i:0;autoSize:true;height:480;width:640}
}
##^##*/
